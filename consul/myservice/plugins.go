package main

import (
	"net"

	consul "github.com/go-micro/plugins/v4/registry/consul"
	consul_api "github.com/hashicorp/consul/api"
	"go-micro.dev/v4/registry"
)

func getConsulRegistry(l net.Listener) (registry.Registry, error) {

	consul.AllowStale(false)

	cfg := consul_api.DefaultConfig()

	//TLS
	tls := consul_api.TLSConfig{
		CAPath:   "../datacenter-deploy-secure/certs/consul-agent-ca.pem",
		CertFile: "../datacenter-deploy-secure/certs/dc1-server-consul-0.pem",
		KeyFile:  "../datacenter-deploy-secure/certs/dc1-server-consul-0-key.pem",
		// InsecureSkipVerify if set to true will disable TLS host verification.
		// InsecureSkipVerify: true,
	}
	cfg.TLSConfig = tls
	cfg.Address = l.Addr().String()
	return consul.NewRegistry(consul.Config(cfg)), nil
}
