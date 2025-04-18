package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type googleAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type googleUserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GoogleAuthClient struct {
	HTTPClient   HTTPClient
	ClientID     string
	ClientSecret string
	RedirectURI  string
	GrantType    string
	logger       *slog.Logger
}

func NewGoogleAuthClient(httpClient HTTPClient, clientID, clientSecret, redirectURI string) *GoogleAuthClient {
	return &GoogleAuthClient{
		HTTPClient:   httpClient,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		GrantType:    "authorization_code",
		logger:       slog.Default().With(slog.String(rsliblog.LoggerNameKey, "GoogleAuthClient")),
	}
}

func (c *GoogleAuthClient) RetrieveAccessToken(ctx context.Context, code string) (*domain.AuthTokenSet, error) {
	ctx, span := tracer.Start(ctx, "googleAuthClient.RetrieveAccessToken")
	defer span.End()

	paramMap := map[string]string{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"redirect_uri":  c.RedirectURI,
		"grant_type":    c.GrantType,
		"code":          code,
	}

	paramBytes, err := json.Marshal(paramMap)
	if err != nil {
		return nil, rsliberrors.Errorf("json.Marshal. err: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://accounts.google.com/o/oauth2/token", bytes.NewBuffer(paramBytes))
	if err != nil {
		return nil, rsliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, rsliberrors.Errorf("retrieve access token.err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, rsliberrors.Errorf("io.ReadAll. err: %w", err)
		}

		if 400 <= resp.StatusCode && resp.StatusCode < 500 {
			c.logger.InfoContext(ctx, fmt.Sprintf("retrieve access token. status: %d, param: %s, body:%s", resp.StatusCode, string(paramBytes), string(respBytes)))
		}

		return nil, fmt.Errorf("retrieve access token. status: %d, body:%s", resp.StatusCode, string(respBytes))
	}

	googleAuthResponse := googleAuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&googleAuthResponse); err != nil {
		return nil, rsliberrors.Errorf("json.NewDecoder. err: %w", err)
	}

	return &domain.AuthTokenSet{
		AccessToken:  googleAuthResponse.AccessToken,
		RefreshToken: googleAuthResponse.RefreshToken,
	}, nil
}

func (c *GoogleAuthClient) RetrieveUserInfo(ctx context.Context, accessToken string) (*domain.UserInfo, error) {
	ctx, span := tracer.Start(ctx, "googleAuthClient.RetrieveUserInfo")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.googleapis.com/oauth2/v1/userinfo", http.NoBody)
	if err != nil {
		return nil, rsliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	q := req.URL.Query()
	q.Add("alt", "json")
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, rsliberrors.Errorf("c.client.Do. err: %w", err)
	}
	defer resp.Body.Close()

	// logger.Debugf("access_token:%s", googleAuthResponse.AccessToken)
	// logger.Debugf("status:%d", resp.StatusCode)

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, rsliberrors.Errorf("retrieve user info. err: %w", domain.ErrUnauthenticated)
	} else if resp.StatusCode != http.StatusOK {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, rsliberrors.Errorf("io.ReadAll. err: %w", err)
		}

		return nil, fmt.Errorf("retrieve user info. status: %d, body:%s", resp.StatusCode, string(respBytes))
	}

	googleUserInfo := googleUserInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&googleUserInfo); err != nil {
		return nil, rsliberrors.Errorf("json.NewDecoder. err: %w", err)
	}

	return &domain.UserInfo{
		Email: googleUserInfo.Email,
		Name:  googleUserInfo.Name,
	}, nil
}
