package utils

import (
	"github.com/escher-client/client"
	"github.com/escher-client/escher_api/escherClient"
)

func GetServiceApi(httpClientWrapper *client.HttpClientWrapper) *escherClient.MgmtServiceApiCtx {
	return httpClientWrapper.EscherClient
}
