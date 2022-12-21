package gconsul

import (
	"github.com/kevin-luvian/goauth/server/pkg/logging"
)

func NotifyHealth(check func() error) {
	err := check()
	if err != nil {
		logging.Infof("err=\"Check failed\" msg=\"%s\"", err.Error())
		if agentErr := instance.ConsulAgent.FailTTL("service:"+instance.Name, err.Error()); agentErr != nil {
			logging.Errorln(agentErr)
		}
	} else {
		if agentErr := instance.ConsulAgent.PassTTL("service:"+instance.Name, ""); agentErr != nil {
			logging.Errorln(agentErr)
		}
	}
}
