package utils

import (
	"github.com/Portshift/escher-client/client"
	"github.com/Portshift/escher-client/escher_api/escherClient"
)

func GetServiceApi(httpClientWrapper *client.HttpClientWrapper) *escherClient.MgmtServiceApiCtx {
	return httpClientWrapper.EscherClient
}
