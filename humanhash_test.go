package humanhashgo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	assert := require.New(t)

	compressed, err := DefaultHasher.compress([]byte{1, 2, 3, 4})
	assert.NoError(err)
	assert.Equal(compressed, []byte{1, 2, 3, 4})

	compressed, err = DefaultHasher.compress([]byte{1, 2, 3, 4, 5})
	assert.NoError(err)
	assert.Equal(compressed, []byte{1, 2, 3, 4 ^ 5})

	compressed, err = DefaultHasher.compress([]byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151, 32})
	assert.NoError(err)
	assert.Equal(compressed, []byte{96 ^ 173 ^ 141, 13 ^ 135 ^ 27, 96 ^ 149 ^ 128, 130 ^ 151 ^ 32})
}
func TestHumanize(t *testing.T) {
	assert := require.New(t)

	data := []byte{1, 2, 3, 4, 5}
	value, err := Humanize(data)
	assert.NoError(err)
	assert.Equal(value, "alabama-alanine-alaska-alabama")

	value, err = Humanize([]byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151, 32})
	assert.NoError(err)
	assert.Equal(value, "equal-monkey-lake-double")
}

func TestNewUuid(t *testing.T) {
	assert := require.New(t)

	value, humanized, err := NewUuid()
	assert.NoError(err)

	compressed := [4]byte{
		value[0] ^ value[1] ^ value[2] ^ value[3],
		value[4] ^ value[5] ^ value[6] ^ value[7],
		value[8] ^ value[9] ^ value[10] ^ value[11],
		value[12] ^ value[13] ^ value[14] ^ value[15],
	}
	humanizedManual := strings.Join(
		[]string{
			DefaultWordlist[compressed[0]],
			DefaultWordlist[compressed[1]],
			DefaultWordlist[compressed[2]],
			DefaultWordlist[compressed[3]],
		},
		"-",
	)

	assert.Equal(humanized, humanizedManual)
}
