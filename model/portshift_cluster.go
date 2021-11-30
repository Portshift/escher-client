package model

import (
	"github.com/Portshift/escher-client/escher_api/model"
)

type Cluster struct {
	ID            string
	ClusterConfig *model.KubernetesCluster
}
