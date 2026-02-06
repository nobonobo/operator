package operator

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const endpoint = "https://rtctunnel-operator.fly.dev"

type Operator struct {
	Endpoint string
}

func New() *Operator {
	return &Operator{Endpoint: endpoint}
}

func (op *Operator) Pub(address string, data []byte) error {
	params := url.Values{}
	params.Set("address", address)
	params.Set("data", string(data))
	resp, err := http.Post(op.Endpoint+"/pub", "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("operator returned status code %d", resp.StatusCode)
	}
	return nil
}

func (op *Operator) Sub(address string) ([]byte, error) {
	params := url.Values{}
	params.Set("address", address)
	resp, err := http.Post(op.Endpoint+"/sub", "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("operator returned status code %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
