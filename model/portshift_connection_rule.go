package model

import (
	"github.com/Portshift/escher-client/escher_api/model"
)

type ConnectionRule struct {
	ID         string
	RuleConfig *model.SecureCNConnectionRule
}
