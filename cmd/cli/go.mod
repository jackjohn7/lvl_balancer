module github.com/jackjohn7/lvl_balancer/cmd/cli

go 1.21.3

require github.com/jackjohn7/lvl_balancer/lib v0.0.0

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/jackjohn7/lvl_balancer/lib v0.0.0 => ../../lib
