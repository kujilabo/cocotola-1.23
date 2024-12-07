package domain

import (
	"time"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
)

type AudioID struct {
	Value int `validate:"required,gte=1"`
}

func NewAudioID(value int) (*AudioID, error) {
	return &AudioID{
		Value: value,
	}, nil
}

func (v *AudioID) Int() int {
	return v.Value
}
func (v *AudioID) IsAudioID() bool {
	return true
}

type AudioModel struct {
	AudioID *AudioID `validate:"required"`
	Lang5   *libdomain.Lang5
	Text    string `validate:"required"`
	Content string `validate:"required"`
	Length  time.Duration
}

func NewAudioModel(audioID *AudioID, lang5 *libdomain.Lang5, text, content string, length time.Duration) (*AudioModel, error) {
	m := &AudioModel{
		AudioID: audioID,
		Lang5:   lang5,
		Text:    text,
		Content: content,
		Length:  length,
	}
	if err := rslibdomain.Validator.Struct(m); err != nil {
		return nil, rsliberrors.Errorf("%s. err: %w", err.Error(), rslibdomain.ErrInvalidField)
	}

	return m, nil
}
