package tatoeba

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/config"
)

var timeoutImportMin = 30

func ImportTatoebaEngSentences() error {
	ctx := context.Background()
	logger := slog.Default()
	cfg, err := config.LoadConfig("local")
	if err != nil {
		return err
	}

	endpoint, err := url.JoinPath(cfg.TatoebaAPI.Endpoint, "v1", "admin", "sentence", "import")
	if err != nil {
		return err
	}

	file, err := os.Open(path.Join(cfg.DataSource.TatoebaDataSource.Dir, cfg.DataSource.TatoebaDataSource.EngSentencesFile))
	if err != nil {
		return err
	}

	body := bytes.Buffer{}
	mw := multipart.NewWriter(&body)

	fw, err := mw.CreateFormFile("file", cfg.DataSource.TatoebaDataSource.EngSentencesFile)
	if err != nil {
		return err
	}

	if _, err := io.Copy(fw, file); err != nil {
		return err
	}

	if err := mw.Close(); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, &body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(cfg.TatoebaAPI.Username, cfg.TatoebaAPI.Password)
	req.Header.Set("Content-Type", mw.FormDataContentType())

	client := http.Client{
		Timeout: time.Duration(timeoutImportMin) * time.Minute,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	logger.InfoContext(ctx, fmt.Sprintf("status: %d", resp.StatusCode))
	logger.InfoContext(ctx, fmt.Sprintf("body: %s", string(respBody)))

	return nil
}
