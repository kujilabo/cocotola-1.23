package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	libapi "github.com/kujilabo/cocotola-1.23/lib/api"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type cocotolaAuthClient struct {
	httpClient   HTTPClient
	authEndpoint *url.URL
	authUsername string
	authPassword string
}

func NewCocotolaAuthClient(httpClient HTTPClient, authEndpoint *url.URL, authUsername, authPassword string) service.CocotolaAuthClient {
	return &cocotolaAuthClient{
		httpClient:   httpClient,
		authEndpoint: authEndpoint,
		authUsername: authUsername,
		authPassword: authPassword,
	}
}

func (c *cocotolaAuthClient) RetrieveUserInfo(ctx context.Context, bearerToken string) (*libapi.AppUserInfoResponse, error) {
	ctx, span := tracer.Start(ctx, "cocotolaAuthClient.RetrieveUserInfo")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "v1", "auth", "userinfo")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, rsliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, rsliberrors.Errorf("auth request. err: %w", err)
	}
	defer resp.Body.Close()

	response := libapi.AppUserInfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, rsliberrors.Errorf("json.NewDecoder. err: %w", err)
	}

	return &response, nil
}
