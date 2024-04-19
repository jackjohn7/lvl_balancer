package lib

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Override default configurations to the LoadBalancer
// using this configuration format. This can also be parsed
// from YAML
type LoadBalancerConfig struct {
	// Determines whether or not to log the proxied website
	// each time a request is made.
	Logging bool

	// TargetFunc is the function that is ran to determine which
	// proxy to hit.
	//
	// If you do not provide a TargetFunc, it will default to
	// round-robin balancing defined in StandardRRTargeting.
	//
	// Custom configuration of this value cannot be done through
	// YAML. You can only use the predefined TargetFuncs this way.
	TargetFunc func(bal *LoadBalancer) Server

	// Middlewares []Middleware // PLANNED
}

type UserConfig struct {
	// spawns a balancer for each configuration
	Targets []ServerConfig `yaml:"targets"`
}

type ServerConfig struct {
	// Port on which to serve proxy
	Port string `yaml:"port"`
	// Servers to create proxy for
	Servers []string `yaml:"servers"`
	// Determines whether or not the server should log activity
	Logging bool `yaml:"logging"`
}

func ParseConfig(file os.File) (LoadBalancerConfig, error) {
	content := make([]byte, 4096)
	bytesRead, err := file.Read(content)
	if err != nil {
		return LoadBalancerConfig{}, err
	}

	// strContent := string(content[0:bytesRead])

	cfg := UserConfig{}
	yaml.Unmarshal(content[0:bytesRead], &cfg)

	fmt.Printf("%+v\n", cfg)

	return LoadBalancerConfig{}, errors.New("Unimplemented")
}
