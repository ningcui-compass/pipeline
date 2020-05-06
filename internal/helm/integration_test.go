// Copyright © 2020 Banzai Cloud
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

package helm_test

import (
	"context"
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/banzaicloud/pipeline/internal/cmd"
	"github.com/banzaicloud/pipeline/internal/common"
	"github.com/banzaicloud/pipeline/internal/common/commonadapter"
	"github.com/banzaicloud/pipeline/internal/global"
	"github.com/banzaicloud/pipeline/internal/helm"
	"github.com/banzaicloud/pipeline/internal/helm/helmadapter"
	"github.com/banzaicloud/pipeline/pkg/k8sclient"
	"github.com/banzaicloud/pipeline/src/secret"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterProviderData struct {
	K8sConfig []byte
	ID        uint
}

func (c *ClusterProviderData) GetID() uint {
	return c.ID
}

func (c *ClusterProviderData) GetK8sConfig() ([]byte, error) {
	return c.K8sConfig, nil
}

type Values struct {
	Service struct {
		ExternalPort int `json:"externalPort,omitempty"`
	} `json:"service,omitempty"`
}

func TestIntegration(t *testing.T) {
	if m := flag.Lookup("test.run").Value.String(); m == "" || !regexp.MustCompile(m).MatchString(t.Name()) {
		t.Skip("skipping as execution was not requested explicitly using go test -run")
	}

	var err error
	global.Config.Helm.Home, err = ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	t.Run("helmV2", testIntegrationV2(global.Config.Helm.Home, "istiofeature-helm-v2"))
	t.Run("helmV3", testIntegrationV3(global.Config.Helm.Home, "istiofeature-helm-v3"))
	t.Run("helmInstallV2", testIntegrationInstall(false, global.Config.Helm.Home, "helm-v2-install"))
	t.Run("helmInstallV3", testIntegrationInstall(true, global.Config.Helm.Home, "helm-v3-install"))
}

func testIntegrationV2(home, testNamespace string) func(t *testing.T) {
	return func(t *testing.T) {
		db := setupDatabase(t)
		secretStore := setupSecretStore()
		kubeConfig, clusterService := clusterKubeConfig(t)

		config := helm.Config{
			Home: home,
			V3:   false,
			Repositories: map[string]string{
				"stable": "https://kubernetes-charts.storage.googleapis.com",
			},
		}

		helmService, _ := cmd.CreateUnifiedHelmReleaser(config, db, secretStore, clusterService, common.NoopLogger{})

		t.Run("testDeleteChartmuseumBeforeSuite", testDeleteChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testCreateChartmuseum", testCreateChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testUpgradeChartmuseum", testUpgradeChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testHandleFailedDeployment", testUpgradeFailedChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testDeleteChartmuseumAfterSuite", testDeleteChartmuseum(helmService, kubeConfig, testNamespace))
	}
}

func testIntegrationV3(home, testNamespace string) func(t *testing.T) {
	return func(t *testing.T) {
		db := setupDatabase(t)
		secretStore := setupSecretStore()
		kubeConfig, clusterService := clusterKubeConfig(t)

		config := helm.Config{
			Home: home,
			V3:   true,
			Repositories: map[string]string{
				"stable": "https://kubernetes-charts.storage.googleapis.com",
			},
		}

		helmService, _ := cmd.CreateUnifiedHelmReleaser(config, db, secretStore, clusterService, common.NoopLogger{})

		t.Run("testDeleteChartmuseumBeforeSuite", testDeleteChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testCreateChartmuseum", testCreateChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testUpgradeChartmuseum", testUpgradeChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testHandleFailedDeployment", testUpgradeFailedChartmuseum(helmService, kubeConfig, testNamespace))
		t.Run("testDeleteChartmuseumAfterSuite", testDeleteChartmuseum(helmService, kubeConfig, testNamespace))
	}
}

func testIntegrationInstall(v3 bool, home, testNamespace string) func(t *testing.T) {
	return func(t *testing.T) {
		db := setupDatabase(t)
		secretStore := setupSecretStore()
		_, clusterService := clusterKubeConfig(t)

		config := helm.Config{
			Home: home,
			V3:   v3,
			Repositories: map[string]string{
				"stable":             "https://kubernetes-charts.storage.googleapis.com",
				"banzaicloud-stable": "https://kubernetes-charts.banzaicloud.com",
			},
		}

		t.Run("helmv3install", func(t *testing.T) {
			releaser, _ := cmd.CreateUnifiedHelmReleaser(config, db, secretStore, clusterService, common.NoopLogger{})

			err := releaser.InstallDeployment(
				context.Background(),
				1,
				testNamespace,
				"banzaicloud-stable/banzaicloud-docs",
				"helm-service-test-v3",
				[]byte{},
				"0.1.2",
				true,
			)
			require.NoError(t, err)

			err = releaser.DeleteDeployment(
				context.Background(),
				1,
				"helm-service-test-v3",
				testNamespace,
			)
			require.NoError(t, err)
		})
	}
}

func testDeleteChartmuseum(helmService helm.UnifiedReleaser, kubeConfig []byte, testNamespace string) func(*testing.T) {
	return func(t *testing.T) {
		err := helmService.Delete(
			&ClusterProviderData{K8sConfig: kubeConfig, ID: 1},
			"chartmuseum",
			testNamespace,
		)
		if err != nil {
			t.Fatalf("%+v", err)
		}

		assertChartmuseumRemoved(t, kubeConfig, testNamespace)
	}
}

func testCreateChartmuseum(helmService helm.UnifiedReleaser, kubeConfig []byte, testNamespace string) func(*testing.T) {
	return func(t *testing.T) {
		err := helmService.InstallOrUpgrade(
			&ClusterProviderData{K8sConfig: kubeConfig, ID: 1},
			helm.Release{
				ReleaseName: "chartmuseum",
				ChartName:   "stable/chartmuseum",
				Namespace:   testNamespace,
				Values:      nil,
				Version:     "2.12.0",
			},
			helm.Options{
				Namespace: testNamespace,
				Wait:      true,
				Install:   true,
			},
		)
		if err != nil {
			t.Fatalf("%+v", err)
		}

		assertChartmuseum(t, kubeConfig, testNamespace, 8080)
	}
}

func testUpgradeChartmuseum(helmService helm.UnifiedReleaser, kubeConfig []byte, testNamespace string) func(*testing.T) {
	return func(t *testing.T) {
		var expectPort int32 = 19191

		values := Values{}
		values.Service.ExternalPort = int(expectPort)

		serializedValues, err := helm.ConvertStructure(values)
		if err != nil {
			t.Fatalf("%+v", serializedValues)
		}

		err = helmService.InstallOrUpgrade(
			&ClusterProviderData{K8sConfig: kubeConfig, ID: 1},
			helm.Release{
				ReleaseName: "chartmuseum",
				ChartName:   "stable/chartmuseum",
				Namespace:   testNamespace,
				Values:      serializedValues,
				Version:     "2.12.0",
			},
			helm.Options{
				Namespace: testNamespace,
				Wait:      true,
				Install:   true,
			},
		)
		if err != nil {
			t.Fatalf("%+v", err)
		}

		assertChartmuseum(t, kubeConfig, testNamespace, expectPort)
	}
}

func testUpgradeFailedChartmuseum(helmService helm.UnifiedReleaser, kubeConfig []byte, testNamespace string) func(*testing.T) {
	return func(t *testing.T) {
		// invalid port will fail the release
		var expectPort int32 = 1111111

		values := Values{}
		values.Service.ExternalPort = int(expectPort)

		serializedValues, err := helm.ConvertStructure(values)
		if err != nil {
			t.Fatalf("%+v", serializedValues)
		}

		err = helmService.InstallOrUpgrade(
			&ClusterProviderData{K8sConfig: kubeConfig, ID: 1},
			helm.Release{
				ReleaseName: "chartmuseum",
				ChartName:   "stable/chartmuseum",
				Namespace:   testNamespace,
				Values:      serializedValues,
				Version:     "2.12.0",
			},
			helm.Options{
				Namespace: testNamespace,
				Wait:      true,
				Install:   true,
			},
		)
		if err == nil {
			t.Fatalf("this upgrade should fail because of the invalid port")
		}

		// restore with original values
		err = helmService.InstallOrUpgrade(
			&ClusterProviderData{K8sConfig: kubeConfig, ID: 1},
			helm.Release{
				ReleaseName: "chartmuseum",
				ChartName:   "stable/chartmuseum",
				Namespace:   testNamespace,
				Values:      nil,
				Version:     "2.12.0",
			},
			helm.Options{
				Namespace: testNamespace,
				Wait:      true,
				Install:   true,
			},
		)
		if err != nil {
			t.Fatalf("%+v", err)
		}

		assertChartmuseum(t, kubeConfig, testNamespace, 8080)
	}
}

func assertChartmuseum(t *testing.T, kubeConfig []byte, testNamespace string, expectedPort int32) {
	restConfig, err := k8sclient.NewClientConfig(kubeConfig)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	ds, err := clientSet.AppsV1().Deployments(testNamespace).Get("chartmuseum-chartmuseum", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if ds.Status.ReadyReplicas != ds.Status.Replicas || ds.Status.ReadyReplicas < 1 {
		t.Fatalf("chartmuseum is not running")
	}

	svc, err := clientSet.CoreV1().Services(testNamespace).Get("chartmuseum-chartmuseum", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(svc.Spec.Ports) < 1 {
		t.Fatalf("Missing chartmuseum service ports")
	}

	if svc.Spec.Ports[0].Port != expectedPort {
		t.Fatalf("chartmuseum service port mismatch, expected %d vs %d", expectedPort, svc.Spec.Ports[0].Port)
	}
}

func assertChartmuseumRemoved(t *testing.T, kubeConfig []byte, testNamespace string) {
	restConfig, err := k8sclient.NewClientConfig(kubeConfig)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	dsList, err := clientSet.AppsV1().Deployments(testNamespace).List(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(dsList.Items) > 0 {
		t.Fatalf("no deployments expected, chartmuseum should be removed")
	}
}

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open("sqlite3", "file::memory:")
	require.NoError(t, err)

	err = helmadapter.Migrate(db, common.NoopLogger{})
	require.NoError(t, err)

	return db
}

func setupSecretStore() common.SecretStore {
	return commonadapter.NewSecretStore(secret.Store, commonadapter.OrgIDContextExtractorFunc(func(ctx context.Context) (uint, bool) {
		return 0, false
	}))
}

func clusterKubeConfig(t *testing.T) ([]byte, helm.ClusterService) {
	kubeConfigFile := os.Getenv("KUBECONFIG")
	if kubeConfigFile == "" {
		t.Skip("skipping as Kubernetes config was not provided")
	}
	kubeConfigBytes, err := ioutil.ReadFile(kubeConfigFile)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	return kubeConfigBytes, helm.ClusterKubeConfigFunc(func(ctx context.Context, clusterID uint) ([]byte, error) {
		return kubeConfigBytes, nil
	})
}
