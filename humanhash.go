package humanhashgo

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type HumanHasher struct {
	Wordlist  []string
	Words     int
	Separator string
}

// Generate a human-readable representation of the data.
func (h *HumanHasher) Humanize(data []byte) (string, error) {
	compressed, err := h.compress(data)
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	for i, d := range compressed {
		if i > 0 {
			builder.WriteString(h.Separator)
		}
		builder.WriteString(h.Wordlist[d])
	}
	return builder.String(), nil
}

// Compress a list of byte values to a fixed HumanHasher.Words length.
func (h *HumanHasher) compress(data []byte) ([]byte, error) {
	length := len(data)
	if h.Words > length {
		return nil, errors.New("fewer input bytes than requested output")
	}
	seg_size := length / h.Words
	checksums := make([]byte, h.Words)
	for i, d := range data {
		i = i / seg_size
		if i >= h.Words {
			i = h.Words - 1
		}
		checksums[i] ^= d
	}
	return checksums, nil
}

// Generate a human-readable representation of UUID.
func (h *HumanHasher) Uuid(value uuid.UUID) (string, error) {
	humanized, err := h.Humanize(value[:])
	if err != nil {
		return "", err
	}
	return humanized, err
}

// Generate a UUID with a human-readable representation.
func (h *HumanHasher) NewUuid() (uuid.UUID, string, error) {
	value, err := uuid.NewRandom()
	if err != nil {
		return value, "", err
	}
	humanized, err := h.Humanize(value[:])
	if err != nil {
		return value, "", err
	}
	return value, humanized, err
}

// Default hasher
var DefaultHasher = HumanHasher{
	Wordlist:  DefaultWordList,
	Words:     4,
	Separator: "-",
}

var Humanize = DefaultHasher.Humanize
var NewUuid = DefaultHasher.NewUuid
