const { ModuleFederationPlugin } = require("webpack").container;
// Remove source code dependency react-lazylog which causes problems with sharing
const { ["react-lazylog"]: undefined, ...sharedDeps } = require("./package.json").dependencies;

module.exports = ({ env }) => ({
    plugins: [
        {
            plugin: {
                overrideWebpackConfig: ({ webpackConfig }) => {
                    // Suppress source-map related warnings (currently, react-use-dimensions is problematic).
                    // This is setting was applied by CRA4, but CRA5 doesn't.
                    webpackConfig.ignoreWarnings = [
                        ...(webpackConfig.ignoreWarnings || []),
                        /Failed to parse source map/,
                    ];
                    return webpackConfig;
                },
            },
        }
    ],
    webpack: {
        plugins: {
            add: [
                new ModuleFederationPlugin({
                    name: "turing",
                    shared: {
                        ...sharedDeps,
                        react: {
                            shareScope: "default",
                            singleton: true,
                            requiredVersion: sharedDeps.react,
                        },
                        "react-dom": {
                            singleton: true,
                            requiredVersion: sharedDeps["react-dom"],
                        },
                    },
                }),
            ],
        }
    },
});