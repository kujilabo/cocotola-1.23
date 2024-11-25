package service

import (
	"context"
	"time"
)

type AudioFile interface {
	Duration(ctx context.Context, audioContent []byte) (time.Duration, error)
}
