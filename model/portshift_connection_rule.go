package model

import (
	"github.com/escher-client/escher_api/model"
)

type ConnectionRule struct {
	ID         string
	RuleConfig *model.SecureCNConnectionRule
}
