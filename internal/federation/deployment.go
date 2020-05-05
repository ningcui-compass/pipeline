// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package federation

import (
	"strings"

	"emperror.dev/errors"

	internalHelm "github.com/banzaicloud/pipeline/internal/helm"
	"github.com/banzaicloud/pipeline/src/cluster"
	"github.com/banzaicloud/pipeline/src/helm"
)

type HelmService interface {
	InstallOrUpgrade(
		c internalHelm.ClusterProvider,
		release internalHelm.Release,
		opts internalHelm.Options,
	) error

	Delete(c internalHelm.ClusterProvider, releaseName, namespace string) error
}

func DeleteDeployment(c cluster.CommonCluster, releaseName string) error {
	kubeConfig, err := c.GetK8sConfig()
	if err != nil {
		return errors.WrapIf(err, "could not get k8s config")
	}

	err = helm.DeleteDeployment(releaseName, kubeConfig)
	if err != nil {
		e := errors.Cause(err)
		if e != nil && strings.Contains(e.Error(), "not found") {
			return nil
		}
		return errors.WrapIf(err, "could not remove deployment")
	}

	return nil
}
