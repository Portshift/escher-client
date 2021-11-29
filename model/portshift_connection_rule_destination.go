package model

import (
	"github.com/escher-client/escher_api/model"
)

type ConnectionRuleDestination struct {
	ID                string
	DestinationConfig *model.SecureCNConnectionRuleDestination
}
