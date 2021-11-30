package model

import (
	"github.com/Portshift/escher-client/escher_api/model"
)

type ConnectionRulePod struct {
	ID        string
	PodConfig *model.SecureCNConnectionRulePod
}
