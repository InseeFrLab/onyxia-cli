module onyxia-cli

go 1.13

replace github.com/docker/docker => github.com/docker/docker v0.0.0-20190731150326-928381b2215c

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kubeapps/kubeapps v1.8.2 // indirect
	github.com/urfave/cli/v2 v2.1.1
	gopkg.in/yaml.v2 v2.2.8
	helm.sh/helm v2.16.3+incompatible // indirect
)
