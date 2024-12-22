package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

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
	var audioModel *domain.AudioModel

	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		// try to find the audio content from the DB
		repo := rf.NewAudioRepository(ctx)
		tmpAudioModel, err := repo.FindByLangAndText(ctx, lang5, text)
		if err != nil {
			if errors.Is(err, service.ErrAudioNotFound) {
				return service.ErrAudioNotFound
			}

			return rsliberrors.Errorf("FindByLangAndText. err: %w", err)
		}

		audioModel = tmpAudioModel
		return nil
	}); err != nil {
		if !errors.Is(err, service.ErrAudioNotFound) {
			return nil, err
		}
	}

	if audioModel != nil {
		return audioModel, nil
	}

	u.logger.InfoContext(ctx, fmt.Sprintf("audio not found. lang: %s, text: %s", lang5.String(), text))

	audioContent, err := u.synthesizerClient.Synthesize(ctx, lang5, voice, text)
	if err != nil {
		return nil, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}
	audioContentBytes, err := base64.StdEncoding.DecodeString(audioContent)
	if err != nil {
		return nil, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}
	audioLength, err := u.audioFile.Duration(ctx, audioContentBytes)
	if err != nil {
		return nil, rsliberrors.Errorf("to u.synthesizerClient.Synthesize. err: %w", err)
	}

	var audioID *domain.AudioID
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		// synthesize text via the Web API
		repo := rf.NewAudioRepository(ctx)
		tmpAudioID, err := repo.AddAudio(ctx, lang5, text, audioContent, audioLength)
		if err != nil {
			return rsliberrors.Errorf("repo.AddAudio. err: %w", err)
		}
		audioID = tmpAudioID
		return nil
	}); err != nil {
		return nil, err
	}

	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		// try to find the audio content from the DB
		repo := rf.NewAudioRepository(ctx)
		tmpAudioModel, err := repo.FindAudioByAudioID(ctx, audioID)
		if err != nil {
			return rsliberrors.Errorf("repo.FindAudioByAudioID. err: %w", err)
		}
		audioModel = tmpAudioModel
		return nil
	}); err != nil {
		return nil, err
	}

	return audioModel, nil
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
