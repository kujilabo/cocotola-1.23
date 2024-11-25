package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type googleTTSClient struct {
	httpClient HTTPClient
	apiKey     string
}

type googleTTSResponse struct {
	AudioContent string `json:"audioContent"`
}

func NewGoogleTTSClient(httpClient HTTPClient, apiKey string) service.SynthesizerClient {
	return &googleTTSClient{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

func (c *googleTTSClient) Synthesize(ctx context.Context, lang5 *libdomain.Lang5, voice, text string) (string, error) {
	ctx, span := tracer.Start(ctx, "googleTTSClient.Synthesize")
	defer span.End()

	type m map[string]interface{}

	values := m{
		"input": m{
			"text": text,
		},
		"voice": m{
			"languageCode": lang5.String(),
			"name":         voice,
		},
		"audioConfig": m{
			"audioEncoding": "MP3",
			"pitch":         0,
			"speakingRate":  1,
		},
	}

	b, err := json.Marshal(values)
	if err != nil {
		return "", rsliberrors.Errorf("json.Marshal. err: %w", err)
	}

	u, err := url.Parse("https://texttospeech.googleapis.com/v1/text:synthesize")
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("key", c.apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(b))
	if err != nil {
		return "", rsliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", rsliberrors.Errorf("%s", string(body))
	}

	googleTTSResponse := googleTTSResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&googleTTSResponse); err != nil {
		return "", rsliberrors.Errorf("json.NewDecoder. err: %w", err)
	}

	return googleTTSResponse.AudioContent, nil
}
