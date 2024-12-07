//go:build small

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
)

func Test_NewAudioModel_shouldReturnError_whenAudioIDIsNil(t *testing.T) {
	t.Parallel()
	lang5, err := libdomain.NewLang5("en-US")
	require.NoError(t, err)

	_, err = domain.NewAudioModel(nil, lang5, "text", "content", time.Second)

	assert.ErrorIs(t, err, rslibdomain.ErrInvalidField)
	assert.Equal(t, err.Error(), "Key: 'AudioModel.AudioID' Error:Field validation for 'AudioID' failed on the 'required' tag. err: invalid field")
}

func Test_NewAudioModel_shouldReturnError_whenLang5IsNil(t *testing.T) {
	t.Parallel()
	audioID, err := domain.NewAudioID(100)
	require.NoError(t, err)

	_, err = domain.NewAudioModel(audioID, nil, "text", "content", time.Second)

	assert.ErrorIs(t, err, rslibdomain.ErrInvalidField)
	assert.Equal(t, err.Error(), "Key: 'AudioModel.Lang5' Error:Field validation for 'Lang5' failed on the 'required' tag. err: invalid field")
}

func Test_NewAudioModel_shouldNotReturnError_whenAllArugumentAreValid(t *testing.T) {
	t.Parallel()
	audioID, err := domain.NewAudioID(100)
	require.NoError(t, err)
	lang5, err := libdomain.NewLang5("en-US")
	require.NoError(t, err)

	model, err := domain.NewAudioModel(audioID, lang5, "text", "content", time.Second)
	assert.NoError(t, err)
	assert.Equal(t, model.AudioID.Int(), 100)
	assert.True(t, model.AudioID.IsAudioID())
}
