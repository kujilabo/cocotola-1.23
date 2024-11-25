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
	AudioID *AudioID
	Lang5   *libdomain.Lang5
	Text    string `validate:"required"`
	Content string `validate:"required"`
	Length  time.Duration
}

func NewAudioModel(audioID *AudioID, Lang5 *libdomain.Lang5, text, content string, length time.Duration) (*AudioModel, error) {
	m := &AudioModel{
		AudioID: audioID,
		Text:    text,
		Content: content,
		Length:  length,
	}
	if err := rslibdomain.Validator.Struct(m); err != nil {
		return nil, rsliberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}
