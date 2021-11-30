package model

import (
	"github.com/Portshift/escher-client/escher_api/model"
)

type ConnectionRuleSource struct {
	ID           string
	SourceConfig *model.SecureCNConnectionRuleSource
}
