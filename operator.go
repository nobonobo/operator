package operator

import (
	"context"
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

func (op *Operator) Pub(ctx context.Context, address string, data []byte) error {
	params := url.Values{}
	params.Set("address", address)
	params.Set("data", string(data))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, op.Endpoint+"/pub", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("operator returned status code %d", resp.StatusCode)
	}
	return nil
}

func (op *Operator) Sub(ctx context.Context, address string) ([]byte, error) {
	params := url.Values{}
	params.Set("address", address)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, op.Endpoint+"/sub", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
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
