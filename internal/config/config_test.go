package config

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"gitlab.com/gitlab-org/gitlab-shell/client/testserver"
	"gitlab.com/gitlab-org/gitlab-shell/internal/testhelper"
)

func TestConfigApplyGlobalState(t *testing.T) {
	t.Cleanup(testhelper.TempEnv(map[string]string{"SSL_CERT_DIR": "unmodified"}))

	config := &Config{SslCertDir: ""}
	config.ApplyGlobalState()

	require.Equal(t, "unmodified", os.Getenv("SSL_CERT_DIR"))

	config.SslCertDir = "foo"
	config.ApplyGlobalState()

	require.Equal(t, "foo", os.Getenv("SSL_CERT_DIR"))
}

func TestHttpClient(t *testing.T) {
	url := testserver.StartHttpServer(t, []testserver.TestRequestHandler{})

	config := &Config{GitlabUrl: url}
	client, err := config.HttpClient()
	require.NoError(t, err)

	_, err = client.Get(url)
	require.NoError(t, err)

	ms, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	lastMetric := ms[0]
	require.Equal(t, lastMetric.GetName(), "gitlab_shell_http_request_seconds")

	labels := lastMetric.GetMetric()[0].Label

	require.Equal(t, "code", labels[0].GetName())
	require.Equal(t, "404", labels[0].GetValue())

	require.Equal(t, "method", labels[1].GetName())
	require.Equal(t, "get", labels[1].GetValue())
}
