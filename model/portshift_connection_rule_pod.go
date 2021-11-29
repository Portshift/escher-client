package model

import (
	"github.com/escher-client/escher_api/model"
)

type ConnectionRulePod struct {
	ID        string
	PodConfig *model.SecureCNConnectionRulePod
}
