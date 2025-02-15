package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"time"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
)

type SynthesizerUsecase struct {
	txManager         service.TransactionManager
	nonTxManager      service.TransactionManager
	synthesizerClient service.SynthesizerClient
	audioFile         service.AudioFile
	logger            *slog.Logger
}

func NewSynthesizerUsecase(txManager, nonTxManager service.TransactionManager, synthesizerClient service.SynthesizerClient, audioFile service.AudioFile) *SynthesizerUsecase {
	return &SynthesizerUsecase{

		txManager:         txManager,
		nonTxManager:      nonTxManager,
		synthesizerClient: synthesizerClient,
		audioFile:         audioFile,
		logger:            slog.Default().With(slog.String(rsliblog.LoggerNameKey, "SynthesizerUsecase")),
	}
}

func (u *SynthesizerUsecase) Synthesize(ctx context.Context, lang5 *libdomain.Lang5, voice, text string) (*domain.AudioModel, error) {
	savedAudioModel, err := rslibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.AudioModel, error) {
		// try to find the audio content from the DB
		repo := rf.NewAudioRepository(ctx)
		savedAudioModel, err := repo.FindByLangAndText(ctx, lang5, text)
		if err != nil {
			// if errors.Is(err, service.ErrAudioNotFound) {
			// 	return nil, service.ErrAudioNotFound
			// }

			return nil, rsliberrors.Errorf("FindByLangAndText. err: %w", err)
		}

		return savedAudioModel, nil
	})
	if err != nil && !errors.Is(err, service.ErrAudioNotFound) {
		return nil, err
	}

	if savedAudioModel != nil {
		return savedAudioModel, nil
	}

	u.logger.InfoContext(ctx, fmt.Sprintf("audio not found. lang: %s, text: %s", lang5.String(), text))

	audioContent, audioLength, err := u.synthesize(ctx, lang5, voice, text)
	if err != nil {
		return nil, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}

	audioID, err := rslibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.AudioID, error) {
		// synthesize text via the Web API
		repo := rf.NewAudioRepository(ctx)
		audioID, err := repo.AddAudio(ctx, lang5, text, audioContent, audioLength)
		if err != nil {
			return nil, rsliberrors.Errorf("repo.AddAudio. err: %w", err)
		}

		return audioID, nil
	})
	if err != nil {
		return nil, err
	}

	audioModel, err := rslibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.AudioModel, error) {
		// try to find the audio content from the DB
		repo := rf.NewAudioRepository(ctx)
		audioModel, err := repo.FindAudioByAudioID(ctx, audioID)
		if err != nil {
			return nil, rsliberrors.Errorf("repo.FindAudioByAudioID. err: %w", err)
		}
		return audioModel, nil
	})
	if err != nil {
		return nil, err
	}

	return audioModel, nil
}

func (u *SynthesizerUsecase) synthesize(ctx context.Context, lang5 *libdomain.Lang5, voice string, text string) (string, time.Duration, error) {
	audioContent, err := u.synthesizerClient.Synthesize(ctx, lang5, voice, text)
	if err != nil {
		return "", 0, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}
	audioContentBytes, err := base64.StdEncoding.DecodeString(audioContent)
	if err != nil {
		return "", 0, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}
	audioLength, err := u.audioFile.Duration(ctx, audioContentBytes)
	if err != nil {
		return "", 0, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}
	return audioContent, audioLength, nil
}

func (u *SynthesizerUsecase) FindAudioByID(ctx context.Context, audioID *domain.AudioID) (*domain.AudioModel, error) {
	var audio *domain.AudioModel
	if err := u.nonTxManager.Do(ctx, func(rf service.RepositoryFactory) error {
		repo := rf.NewAudioRepository(ctx)
		tmpAudio, err := repo.FindAudioByAudioID(ctx, audioID)
		if err != nil {
			return err
		}

		audio = tmpAudio
		return nil
	}); err != nil {
		return nil, err
	}

	return audio, nil
}
