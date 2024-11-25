package service

import (
	"context"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
)

type SynthesizerClient interface {
	Synthesize(ctx context.Context, lang5 *libdomain.Lang5, voice, text string) (string, error)
}
