package gconsul

import (
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

type Consul struct {
	Name        string
	RootFolder  string
	HealthTTL   time.Duration
	ConsulAgent *consul.Agent
	ConsulKV    *consul.KV
}

var instance = &Consul{}

func Setup() error {
	config := consul.DefaultConfig()
	config.Address = setting.Consul.Address

	client, err := consul.NewClient(config)
	if err != nil {
		return err
	}

	instance = &Consul{
		Name:        setting.Consul.ServiceName,
		RootFolder:  setting.Consul.RootFolder,
		HealthTTL:   setting.App.TickerTTL,
		ConsulAgent: client.Agent(),
		ConsulKV:    client.KV(),
	}

	serviceDef := &consul.AgentServiceRegistration{
		ID:   instance.Name,
		Name: instance.Name,
		Check: &consul.AgentServiceCheck{
			TTL:                            instance.HealthTTL.String(),
			DeregisterCriticalServiceAfter: (instance.HealthTTL * 2).String(),
		},
	}

	err = instance.ConsulAgent.ServiceRegister(serviceDef)
	if err != nil {
		return err
	}

	return nil
}
