package gateway

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
)

type audioEntity struct {
	ID             int    `validate:"required"`
	Lang5          string `validate:"required"`
	Text           string `validate:"required"`
	AudioContent   string `validate:"required"`
	AudioLengthSec float64
}

func (e *audioEntity) TableName() string {
	return "audio"
}

func (e *audioEntity) toAudioModel() (*domain.AudioModel, error) {
	audioID, err := domain.NewAudioID(e.ID)
	if err != nil {
		return nil, rsliberrors.Errorf("domain.NewAppUserModel. err: %w", err)
	}

	lang5, err := libdomain.NewLang5(e.Lang5)
	if err != nil {
		return nil, err
	}

	return domain.NewAudioModel(audioID, lang5, e.Text, e.AudioContent, time.Duration(float64(time.Second)*e.AudioLengthSec))
}

type audioRepository struct {
	db *gorm.DB
}

func newAudioRepository(ctx context.Context, db *gorm.DB) service.AudioRepository {
	return &audioRepository{
		db: db,
	}
}

func (r *audioRepository) AddAudio(ctx context.Context, lang5 *libdomain.Lang5, text, audioContent string, audioLength time.Duration) (*domain.AudioID, error) {
	_, span := tracer.Start(ctx, "audioRepository.AddAudio")
	defer span.End()

	entity := audioEntity{
		Lang5:          lang5.String(),
		Text:           text,
		AudioContent:   audioContent,
		AudioLengthSec: audioLength.Seconds(),
	}
	if result := r.db.Create(&entity); result.Error != nil {
		return nil, result.Error
	}

	audioID, err := domain.NewAudioID(entity.ID)
	if err != nil {
		return nil, rsliberrors.Errorf("domain.NewAppUserModel. err: %w", err)
	}

	return audioID, nil
}

func (r *audioRepository) FindAudioByAudioID(ctx context.Context, audioID *domain.AudioID) (*domain.AudioModel, error) {
	_, span := tracer.Start(ctx, "audioRepository.FindAudioByAudioID")
	defer span.End()

	entity := audioEntity{}
	if result := r.db.Where("id = ?", audioID.Int()).First(&entity); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrAudioNotFound
		}
		return nil, result.Error
	}
	audioModel, err := entity.toAudioModel()
	if err != nil {
		return nil, err
	}
	return audioModel, nil
}

func (r *audioRepository) FindByLangAndText(ctx context.Context, lang5 *libdomain.Lang5, text string) (*domain.AudioModel, error) {
	_, span := tracer.Start(ctx, "audioRepository.FindByLangAndText")
	defer span.End()

	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	entity := audioEntity{}
	if result := r.db.Where("lang5 = ? and text = ?", lang5.String(), text).First(&entity); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.InfoContext(ctx, "audio not found. lang: %s, text: %s", lang5.String(), text)
			return nil, service.ErrAudioNotFound
		}

		logger.InfoContext(ctx, fmt.Sprintf("err: %v", result.Error))
		return nil, result.Error
	}

	logger.InfoContext(ctx, fmt.Sprintf("found: %s", text))
	audioModel, err := entity.toAudioModel()
	if err != nil {
		logger.InfoContext(ctx, fmt.Sprintf("err: %v", err))
		return nil, err
	}

	return audioModel, nil
}

func (r *audioRepository) FindAudioIDByText(ctx context.Context, lang5 *libdomain.Lang5, text string) (*domain.AudioID, error) {
	_, span := tracer.Start(ctx, "audioRepository.FindAudioIDByText")
	defer span.End()

	entity := audioEntity{}
	if result := r.db.Where("lang5 = ? and text = ?", lang5.String(), text).First(&entity); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrAudioNotFound
		}
		return nil, result.Error
	}

	audioID, err := domain.NewAudioID(entity.ID)
	if err != nil {
		return nil, rsliberrors.Errorf("domain.NewAppUserModel. err: %w", err)
	}

	return audioID, nil
}
