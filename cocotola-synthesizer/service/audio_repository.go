package service

import (
	"context"
	"errors"
	"time"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
)

var ErrAudioNotFound = errors.New("Audio not found")

type AudioRepository interface {
	AddAudio(ctx context.Context, lang5 *libdomain.Lang5, text, audioContent string, audioLength time.Duration) (*domain.AudioID, error)

	FindAudioByAudioID(ctx context.Context, audioID *domain.AudioID) (*domain.AudioModel, error)

	FindByLangAndText(ctx context.Context, lang5 *libdomain.Lang5, text string) (*domain.AudioModel, error)

	FindAudioIDByText(ctx context.Context, lang5 *libdomain.Lang5, text string) (*domain.AudioID, error)
}
