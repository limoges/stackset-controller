package main

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	zv1 "github.com/zalando-incubator/stackset-controller/pkg/apis/zalando.org/v1"
	"github.com/zalando-incubator/stackset-controller/pkg/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/tools/clientcmd"
)

func TestCRD(t *testing.T) {
	kubeconfig := flag.String(
		"kubeconfig", LookupEnvOrString("KUBECONFIG", "~/.kube/config"), "kubeconfig file",
	)

	flag.Parse()
	ctx := context.Background()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	require.NoError(t, err)

	client, err := clientset.NewForConfig(config)
	require.NoError(t, err)

	test, err := client.ZalandoV1().Tests("default").Create(
		ctx,
		&zv1.Test{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-1",
				Labels: map[string]string{
					"user": "msaleem",
				},
			},
			Spec: zv1.TestSpec{},
		},
		metav1.CreateOptions{},
	)
	require.NoError(t, err)

	require.Equal(t, "test-1", test.Name)

}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
