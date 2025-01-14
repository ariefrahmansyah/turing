// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cluster "github.com/caraml-dev/turing/api/turing/cluster"

	corev1 "k8s.io/api/core/v1"

	io "io"

	mock "github.com/stretchr/testify/mock"

	rbacv1 "k8s.io/api/rbac/v1"

	v1 "k8s.io/api/batch/v1"

	v1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// ApplyConfigMap provides a mock function with given fields: ctx, namespace, configMap
func (_m *Controller) ApplyConfigMap(ctx context.Context, namespace string, configMap *cluster.ConfigMap) error {
	ret := _m.Called(ctx, namespace, configMap)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.ConfigMap) error); ok {
		r0 = rf(ctx, namespace, configMap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplyIstioVirtualService provides a mock function with given fields: ctx, routerEndpoint
func (_m *Controller) ApplyIstioVirtualService(ctx context.Context, routerEndpoint *cluster.VirtualService) error {
	ret := _m.Called(ctx, routerEndpoint)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.VirtualService) error); ok {
		r0 = rf(ctx, routerEndpoint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplyPersistentVolumeClaim provides a mock function with given fields: ctx, namespace, pvc
func (_m *Controller) ApplyPersistentVolumeClaim(ctx context.Context, namespace string, pvc *cluster.PersistentVolumeClaim) error {
	ret := _m.Called(ctx, namespace, pvc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.PersistentVolumeClaim) error); ok {
		r0 = rf(ctx, namespace, pvc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateJob provides a mock function with given fields: ctx, namespace, job
func (_m *Controller) CreateJob(ctx context.Context, namespace string, job cluster.Job) (*v1.Job, error) {
	ret := _m.Called(ctx, namespace, job)

	var r0 *v1.Job
	if rf, ok := ret.Get(0).(func(context.Context, string, cluster.Job) *v1.Job); ok {
		r0 = rf(ctx, namespace, job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, cluster.Job) error); ok {
		r1 = rf(ctx, namespace, job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNamespace provides a mock function with given fields: ctx, name
func (_m *Controller) CreateNamespace(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRole provides a mock function with given fields: ctx, namespace, role
func (_m *Controller) CreateRole(ctx context.Context, namespace string, role *cluster.Role) (*rbacv1.Role, error) {
	ret := _m.Called(ctx, namespace, role)

	var r0 *rbacv1.Role
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.Role) *rbacv1.Role); ok {
		r0 = rf(ctx, namespace, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *cluster.Role) error); ok {
		r1 = rf(ctx, namespace, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRoleBinding provides a mock function with given fields: ctx, namespace, roleBinding
func (_m *Controller) CreateRoleBinding(ctx context.Context, namespace string, roleBinding *cluster.RoleBinding) (*rbacv1.RoleBinding, error) {
	ret := _m.Called(ctx, namespace, roleBinding)

	var r0 *rbacv1.RoleBinding
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.RoleBinding) *rbacv1.RoleBinding); ok {
		r0 = rf(ctx, namespace, roleBinding)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.RoleBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *cluster.RoleBinding) error); ok {
		r1 = rf(ctx, namespace, roleBinding)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSecret provides a mock function with given fields: ctx, secret
func (_m *Controller) CreateSecret(ctx context.Context, secret *cluster.Secret) error {
	ret := _m.Called(ctx, secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.Secret) error); ok {
		r0 = rf(ctx, secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateServiceAccount provides a mock function with given fields: ctx, namespace, serviceAccount
func (_m *Controller) CreateServiceAccount(ctx context.Context, namespace string, serviceAccount *cluster.ServiceAccount) (*corev1.ServiceAccount, error) {
	ret := _m.Called(ctx, namespace, serviceAccount)

	var r0 *corev1.ServiceAccount
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.ServiceAccount) *corev1.ServiceAccount); ok {
		r0 = rf(ctx, namespace, serviceAccount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *cluster.ServiceAccount) error); ok {
		r1 = rf(ctx, namespace, serviceAccount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSparkApplication provides a mock function with given fields: ctx, namespace, request
func (_m *Controller) CreateSparkApplication(ctx context.Context, namespace string, request *cluster.CreateSparkRequest) (*v1beta2.SparkApplication, error) {
	ret := _m.Called(ctx, namespace, request)

	var r0 *v1beta2.SparkApplication
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.CreateSparkRequest) *v1beta2.SparkApplication); ok {
		r0 = rf(ctx, namespace, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.SparkApplication)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *cluster.CreateSparkRequest) error); ok {
		r1 = rf(ctx, namespace, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConfigMap provides a mock function with given fields: ctx, name, namespace, ignoreNotFound
func (_m *Controller) DeleteConfigMap(ctx context.Context, name string, namespace string, ignoreNotFound bool) error {
	ret := _m.Called(ctx, name, namespace, ignoreNotFound)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, name, namespace, ignoreNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIstioVirtualService provides a mock function with given fields: ctx, svcName, namespace
func (_m *Controller) DeleteIstioVirtualService(ctx context.Context, svcName string, namespace string) error {
	ret := _m.Called(ctx, svcName, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, svcName, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, namespace, jobName
func (_m *Controller) DeleteJob(ctx context.Context, namespace string, jobName string) error {
	ret := _m.Called(ctx, namespace, jobName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, namespace, jobName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteKnativeService provides a mock function with given fields: ctx, svcName, namespace, ignoreNotFound
func (_m *Controller) DeleteKnativeService(ctx context.Context, svcName string, namespace string, ignoreNotFound bool) error {
	ret := _m.Called(ctx, svcName, namespace, ignoreNotFound)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, svcName, namespace, ignoreNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteKubernetesDeployment provides a mock function with given fields: ctx, name, namespace, ignoreNotFound
func (_m *Controller) DeleteKubernetesDeployment(ctx context.Context, name string, namespace string, ignoreNotFound bool) error {
	ret := _m.Called(ctx, name, namespace, ignoreNotFound)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, name, namespace, ignoreNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteKubernetesService provides a mock function with given fields: ctx, svcName, namespace, ignoreNotFound
func (_m *Controller) DeleteKubernetesService(ctx context.Context, svcName string, namespace string, ignoreNotFound bool) error {
	ret := _m.Called(ctx, svcName, namespace, ignoreNotFound)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, svcName, namespace, ignoreNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePersistentVolumeClaim provides a mock function with given fields: ctx, pvcName, namespace, ignoreNotFound
func (_m *Controller) DeletePersistentVolumeClaim(ctx context.Context, pvcName string, namespace string, ignoreNotFound bool) error {
	ret := _m.Called(ctx, pvcName, namespace, ignoreNotFound)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, pvcName, namespace, ignoreNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSecret provides a mock function with given fields: ctx, secretName, namespace, ignoreNotFound
func (_m *Controller) DeleteSecret(ctx context.Context, secretName string, namespace string, ignoreNotFound bool) error {
	ret := _m.Called(ctx, secretName, namespace, ignoreNotFound)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, secretName, namespace, ignoreNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSparkApplication provides a mock function with given fields: ctx, namespace, appName
func (_m *Controller) DeleteSparkApplication(ctx context.Context, namespace string, appName string) error {
	ret := _m.Called(ctx, namespace, appName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, namespace, appName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployKnativeService provides a mock function with given fields: ctx, svc
func (_m *Controller) DeployKnativeService(ctx context.Context, svc *cluster.KnativeService) error {
	ret := _m.Called(ctx, svc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.KnativeService) error); ok {
		r0 = rf(ctx, svc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployKubernetesService provides a mock function with given fields: ctx, svc
func (_m *Controller) DeployKubernetesService(ctx context.Context, svc *cluster.KubernetesService) error {
	ret := _m.Called(ctx, svc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.KubernetesService) error); ok {
		r0 = rf(ctx, svc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetJob provides a mock function with given fields: ctx, namespace, jobName
func (_m *Controller) GetJob(ctx context.Context, namespace string, jobName string) (*v1.Job, error) {
	ret := _m.Called(ctx, namespace, jobName)

	var r0 *v1.Job
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1.Job); ok {
		r0 = rf(ctx, namespace, jobName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, jobName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKnativeServiceURL provides a mock function with given fields: ctx, svcName, namespace
func (_m *Controller) GetKnativeServiceURL(ctx context.Context, svcName string, namespace string) string {
	ret := _m.Called(ctx, svcName, namespace)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, svcName, namespace)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSparkApplication provides a mock function with given fields: ctx, namespace, appName
func (_m *Controller) GetSparkApplication(ctx context.Context, namespace string, appName string) (*v1beta2.SparkApplication, error) {
	ret := _m.Called(ctx, namespace, appName)

	var r0 *v1beta2.SparkApplication
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1beta2.SparkApplication); ok {
		r0 = rf(ctx, namespace, appName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta2.SparkApplication)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, appName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPodLogs provides a mock function with given fields: ctx, namespace, podName, opts
func (_m *Controller) ListPodLogs(ctx context.Context, namespace string, podName string, opts *corev1.PodLogOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, namespace, podName, opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *corev1.PodLogOptions) io.ReadCloser); ok {
		r0 = rf(ctx, namespace, podName, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *corev1.PodLogOptions) error); ok {
		r1 = rf(ctx, namespace, podName, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPods provides a mock function with given fields: ctx, namespace, labelSelector
func (_m *Controller) ListPods(ctx context.Context, namespace string, labelSelector string) (*corev1.PodList, error) {
	ret := _m.Called(ctx, namespace, labelSelector)

	var r0 *corev1.PodList
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *corev1.PodList); ok {
		r0 = rf(ctx, namespace, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PodList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewController interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t mockConstructorTestingTNewController) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
