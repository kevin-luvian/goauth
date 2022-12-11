package gconsul

import (
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

type Consul struct {
	Name        string
	HealthTTL   time.Duration
	WatchTTL    time.Duration
	ConsulAgent *consul.Agent
	ConsulKV    *consul.KV
}

type Dependencies struct {
	Check func() error
}

var instance = &Consul{}

func Setup(check func() error) error {
	config := consul.DefaultConfig()
	config.Address = setting.Consul.Address

	client, err := consul.NewClient(config)
	if err != nil {
		return err
	}

	instance = &Consul{
		Name:        setting.Consul.ServiceName,
		HealthTTL:   setting.Consul.HealthTTL,
		WatchTTL:    setting.Consul.WatchTTL,
		ConsulAgent: client.Agent(),
		ConsulKV:    client.KV(),
	}

	serviceDef := &consul.AgentServiceRegistration{
		ID:   instance.Name,
		Name: instance.Name,
		Check: &consul.AgentServiceCheck{
			TTL: instance.HealthTTL.String(),
		},
	}

	instance.ConsulAgent.ServiceDeregister(serviceDef.ID)
	err = instance.ConsulAgent.ServiceRegister(serviceDef)
	if err != nil {
		return err
	}

	go instance.UpdateTTL(check)

	return nil
}
