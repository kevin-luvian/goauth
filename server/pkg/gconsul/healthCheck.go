package gconsul

import (
	"log"
	"time"

	"github.com/kevin-luvian/goauth/server/pkg/logging"
)

func (c *Consul) UpdateTTL(check func() error) {
	ticker := time.NewTicker(c.HealthTTL / 2)
	for range ticker.C {
		c.update(check)
	}
}

func (c *Consul) update(check func() error) {
	err := check()
	if err != nil {
		log.Printf("err=\"Check failed\" msg=\"%s\"", err.Error())
		if agentErr := c.ConsulAgent.FailTTL("service:"+c.Name, err.Error()); agentErr != nil {
			logging.Errorln(agentErr)
		}
	} else {
		if agentErr := c.ConsulAgent.PassTTL("service:"+c.Name, ""); agentErr != nil {
			logging.Errorln(agentErr)
		}
	}
}
