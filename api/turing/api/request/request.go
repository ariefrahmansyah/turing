package request

import (
	"errors"
	"fmt"

	"github.com/gojek/turing/api/turing/config"
	"github.com/gojek/turing/api/turing/models"
	"github.com/gojek/turing/api/turing/service"
	"github.com/gojek/turing/engines/experiment/manager"
	routercfg "github.com/gojek/turing/engines/router/missionctl/config"
)

// CreateOrUpdateRouterRequest structure defines the format of the request payload
// when creating or updating routers
type CreateOrUpdateRouterRequest struct {
	Environment string        `json:"environment_name" validate:"required"`
	Name        string        `json:"name" validate:"required"`
	Config      *RouterConfig `json:"config" validate:"required,dive"`
}

// RouterConfig defines the properties of the specific router version
type RouterConfig struct {
	Routes           models.Routes           `json:"routes" validate:"required"`
	DefaultRouteID   string                  `json:"default_route_id" validate:"required"`
	TrafficRules     models.TrafficRules     `json:"rules" validate:"dive"`
	ExperimentEngine *ExperimentEngineConfig `json:"experiment_engine" validate:"required,dive"`
	ResourceRequest  *models.ResourceRequest `json:"resource_request"`
	Timeout          string                  `json:"timeout" validate:"required"`

	LogConfig *LogConfig `json:"log_config" validate:"required"`

	Enricher  *EnricherEnsemblerConfig `json:"enricher,omitempty" validate:"omitempty,dive"`
	Ensembler *models.Ensembler        `json:"ensembler,omitempty" validate:"omitempty,dive"`
}

// ExperimentEngineConfig defines the experiment engine config
type ExperimentEngineConfig struct {
	Type   string      `json:"type" validate:"required,oneof=litmus nop xp"`
	Config interface{} `json:"config,omitempty" validate:"-"` // Skip validate to invoke custom validation
}

// LogConfig defines the logging configs
type LogConfig struct {
	ResultLoggerType models.ResultLogger `json:"result_logger_type"`
	BigQueryConfig   *BigQueryConfig     `json:"bigquery_config,omitempty"`
	KafkaConfig      *KafkaConfig        `json:"kafka_config,omitempty"`
}

// BigQueryConfig defines the configs for logging to BQ
type BigQueryConfig struct {
	Table                string `json:"table"`
	ServiceAccountSecret string `json:"service_account_secret"`
}

// KafkaConfig defines the configs for logging to Kafka
type KafkaConfig struct {
	Brokers             string                     `json:"brokers"`
	Topic               string                     `json:"topic"`
	SerializationFormat models.SerializationFormat `json:"serialization_format"`
}

// EnricherEnsemblerConfig defines the configs for the enricher / ensembler,
// used by the specific router config
type EnricherEnsemblerConfig struct {
	// Fully qualified docker image string used by the enricher, in the
	// format registry/repository:version.
	Image string `json:"image" validate:"required"`
	// Resource requests  for the deployment of the enricher.
	ResourceRequest *models.ResourceRequest `json:"resource_request" validate:"required"`
	// Endpoint to query.
	Endpoint string `json:"endpoint" validate:"required"`
	// Request timeout as a valid quantity string.
	Timeout string `json:"timeout" validate:"required"`
	// Port to query.
	Port int `json:"port" validate:"required"`
	// Environment variables to inject into the pod.
	Env models.EnvVars `json:"env" validate:"required"`
	// ServiceAccount specifies the name of the secret registered in the MLP project containing the service account.
	// The service account will be mounted into the user-container and the environment variable
	// GOOGLE_APPLICATION_CREDENTIALS will reference the service account file.
	ServiceAccount string `json:"service_account"`
}

// BuildEnricher builds the enricher model from the enricher config
func (cfg EnricherEnsemblerConfig) BuildEnricher() *models.Enricher {
	return &models.Enricher{
		Image:           cfg.Image,
		ResourceRequest: cfg.ResourceRequest,
		Endpoint:        cfg.Endpoint,
		Timeout:         cfg.Timeout,
		Port:            cfg.Port,
		Env:             cfg.Env,
		ServiceAccount:  cfg.ServiceAccount,
	}
}

// BuildRouter builds the router model from the entire request payload
func (r CreateOrUpdateRouterRequest) BuildRouter(projectID models.ID) *models.Router {
	return &models.Router{
		ProjectID:       projectID,
		EnvironmentName: r.Environment,
		Name:            r.Name,
		Status:          models.RouterStatusPending,
	}
}

// BuildRouterVersion builds the router version model from the entire request payload
func (r CreateOrUpdateRouterRequest) BuildRouterVersion(
	router *models.Router,
	defaults *config.RouterDefaults,
	cryptoSvc service.CryptoService,
	expSvc service.ExperimentsService,
) (*models.RouterVersion, error) {
	if r.Config == nil {
		return nil, errors.New("router config is empty")
	}
	rv := &models.RouterVersion{
		RouterID:       router.ID,
		Router:         router,
		Image:          defaults.Image,
		Status:         models.RouterVersionStatusPending,
		Routes:         r.Config.Routes,
		DefaultRouteID: r.Config.DefaultRouteID,
		TrafficRules:   r.Config.TrafficRules,
		ExperimentEngine: &models.ExperimentEngine{
			Type: models.ExperimentEngineType(r.Config.ExperimentEngine.Type),
		},
		ResourceRequest: r.Config.ResourceRequest,
		Timeout:         r.Config.Timeout,
		LogConfig: &models.LogConfig{
			LogLevel:             routercfg.LogLevel(defaults.LogLevel),
			CustomMetricsEnabled: defaults.CustomMetricsEnabled,
			FiberDebugLogEnabled: defaults.FiberDebugLogEnabled,
			JaegerEnabled:        defaults.JaegerEnabled,
			ResultLoggerType:     models.ResultLogger(r.Config.LogConfig.ResultLoggerType),
		},
	}
	if r.Config.Enricher != nil {
		rv.Enricher = r.Config.Enricher.BuildEnricher()
	}
	if r.Config.Ensembler != nil {
		// Ensure ensembler config is set based on the ensembler type
		if r.Config.Ensembler.Type == models.EnsemblerDockerType && r.Config.Ensembler.DockerConfig == nil {
			return nil, errors.New("missing ensembler docker config")
		}
		if r.Config.Ensembler.Type == models.EnsemblerStandardType && r.Config.Ensembler.StandardConfig == nil {
			return nil, errors.New("missing ensembler standard config")
		}
		rv.Ensembler = r.Config.Ensembler
	}
	switch rv.LogConfig.ResultLoggerType {
	case models.BigQueryLogger:
		rv.LogConfig.BigQueryConfig = &models.BigQueryConfig{
			Table:                r.Config.LogConfig.BigQueryConfig.Table,
			ServiceAccountSecret: r.Config.LogConfig.BigQueryConfig.ServiceAccountSecret,
			BatchLoad:            true, // default for now
		}
	case models.KafkaLogger:
		rv.LogConfig.KafkaConfig = &models.KafkaConfig{
			Brokers:             r.Config.LogConfig.KafkaConfig.Brokers,
			Topic:               r.Config.LogConfig.KafkaConfig.Topic,
			SerializationFormat: r.Config.LogConfig.KafkaConfig.SerializationFormat,
		}
	}
	if rv.ExperimentEngine.Type != models.ExperimentEngineTypeNop {
		experimentConfig, err := r.BuildExperimentEngineConfig(router, cryptoSvc, expSvc)
		if err != nil {
			return nil, err
		}
		rv.ExperimentEngine.Config = experimentConfig
	}

	return rv, nil
}

// BuildExperimentEngineConfig creates the Experiment config from the given input properties
func (r CreateOrUpdateRouterRequest) BuildExperimentEngineConfig(
	router *models.Router,
	cryptoSvc service.CryptoService,
	expSvc service.ExperimentsService,
) (interface{}, error) {
	rawExpConfig := r.Config.ExperimentEngine.Config

	// Handle missing passkey / encrypt it in Standard experiment config
	if expSvc.IsStandardExperimentManager(r.Config.ExperimentEngine.Type) {
		// Convert the new config to the standard type
		expConfig, err := expSvc.GetStandardExperimentConfig(rawExpConfig)
		if err != nil {
			return nil, fmt.Errorf("Cannot parse standard experiment config: %v", err)
		}
		clientPasskey := expConfig.Client.Passkey

		if clientPasskey == "" {
			// Extract existing router version config
			if router.CurrRouterVersion != nil &&
				string(router.CurrRouterVersion.ExperimentEngine.Type) == r.Config.ExperimentEngine.Type {
				currVerExpConfig, err := expSvc.GetStandardExperimentConfig(router.CurrRouterVersion.ExperimentEngine.Config)
				if err != nil {
					return nil, fmt.Errorf("Error parsing existing experiment config: %v", err)
				}
				if expConfig.Client.Username == currVerExpConfig.Client.Username {
					// Copy the passkey
					clientPasskey = currVerExpConfig.Client.Passkey
				}
			}
			// If the passkey is still empty, we cannot proceed
			if clientPasskey == "" {
				return nil, errors.New("Passkey must be configured")
			}
		} else {
			// Passkey has been supplied, encrypt it
			var err error
			clientPasskey, err = cryptoSvc.Encrypt(clientPasskey)
			if err != nil {
				return nil, fmt.Errorf("Passkey could not be encrypted: %s", err.Error())
			}
		}

		// Build Experiment engine config
		return &manager.TuringExperimentConfig{
			Client: manager.Client{
				ID:       expConfig.Client.ID,
				Username: expConfig.Client.Username,
				Passkey:  clientPasskey,
			},
			Experiments: expConfig.Experiments,
			Variables:   expConfig.Variables,
		}, nil
	}

	// Custom experiment manager config, return as is.
	return rawExpConfig, nil
}
