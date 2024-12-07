//go:build small

package domain_test

import (
	"testing"
	"time"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/domain"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewAudioModel_shouldReturnError_whenAudioIDIsNil(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T) {
		t.Helper()
		lang5, err := libdomain.NewLang5("en-US")
		require.NoError(t, err)

		_, err = domain.NewAudioModel(nil, lang5, "text", "content", time.Second)

		assert.ErrorIs(t, err, rslibdomain.ErrInvalidField)
		assert.Equal(t, err.Error(), "Key: 'AudioModel.AudioID' Error:Field validation for 'AudioID' failed on the 'required' tag. err: invalid field")
	}
	fn(t)
}
