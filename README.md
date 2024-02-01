# LVL Balancer

LVL (pronounced "Level") is a simple load balancer written in Go. You can configure
LVL with YAML or by using arguments to the CLI tool. This is a WIP, so some functionality
may be a bit rough around the edges or nonexistent alltogether.

Currently, LVL is focused on load balancing of multiple services, but it would be nice to 
have more general reverse proxy features eventually.

# Why

The idea of LVL is to have a more modern and intuitive experience than NGINX or another
tool that can be used for reverse proxying or load balancing.

# Planned Features

## Init Command

I would like to have a command in the CLI that allows a user to quickly scaffold a 
configuration file or dockerfile for their LVL load balancer. Knowing what configurations
do and don't exist can be learned through documentation, but it's also nice to just have 
them in front of you.

## Documentation

There currently exists no documentation for this package aside from comments I've added
to certain types, fields, and functions.

In a similar vein, I would like to have more examples in the examples folder for 
use cases that are meant to be supported by the application. As of now, the only example 
is the most basic example that shows how to use the most basic features of the API.

## Multiple load balancers on different paths

I would like to allow a user to define services in the YAML file and load balance them on
their endpoints. For example, you have a frontend and a backend. Your frontend has two
instances, and your backend has three. You should be able to load balance your frontend
on one route while load balancing your backend on another route. The API does allow for 
such functionality. It need only be implemented in the CLI, YAML configuration, etc.

## TOML Support

For those of us who prefer TOML for configuration, TOML should be an available syntax 
for configuring LVL.

## Docker Container

I would like to have a pullable docker container for running this reverse proxy in docker.
The docker container should be able to either be configured through environment variables
in docker compose, a YAML/TOML file provided, or command-line arguments through the 
docker command line.

The container should be small making use of multi-stage builds such that the Go toolchain
is not bundled in with the container.

## Support for Web Frameworks

At the moment, the load balancer only accepts an `http.ServeMux` from the standard
library, but I'd like to allow for more batteries included frameworks like Echo, Gin, etc.
to serve a load balancer as well.

## More Targeting Strategies

I would like to extend the targeting API to allow a user to write more sophisticated
targeting logic. Some of these targeting strategies may require the usage of middleware.

## Middleware

I would like to extend to the API to allow a user to define middleware functions 
in the standard way that they usually are for web services in Golang.
