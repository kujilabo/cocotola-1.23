package gateway

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"github.com/tcolgate/mp3"
)

type AudioFile struct {
}

func NewAudioFile() *AudioFile {
	return &AudioFile{}
}

func (a *AudioFile) Duration(ctx context.Context, audioContent []byte) (time.Duration, error) {
	reader := bytes.NewReader(audioContent)
	readCloser := io.NopCloser(reader)
	d := mp3.NewDecoder(readCloser)
	var f mp3.Frame
	skipped := 0
	var t float64
	for {
		if err := d.Decode(&f, &skipped); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, err
		}
		t = t + f.Duration().Seconds()
	}
	return time.Duration(t * float64(time.Second)), nil
}
