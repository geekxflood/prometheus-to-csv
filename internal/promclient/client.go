// internal/promclient/client.go
package promclient

import (
	"crypto/tls"
	"net/http"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var InsecureSkipVerify bool

func CreatePrometheusClient(address string) (v1.API, error) {
	httpClientConfig := api.DefaultRoundTripper.(*http.Transport).Clone()
	if InsecureSkipVerify {
		httpClientConfig.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client, err := api.NewClient(api.Config{
		Address:      address,
		RoundTripper: httpClientConfig,
	})
	if err != nil {
		return nil, err
	}

	return v1.NewAPI(client), nil
}
