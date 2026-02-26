package humanhashgo

import (
	"slices"
	"strings"
	"testing"
)

func TestCompress(t *testing.T) {
	compressed, err := DefaultHasher.compress([]byte{1, 2, 3, 4})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []byte{1, 2, 3, 4}
	if !slices.Equal(compressed, []byte{1, 2, 3, 4}) {
		t.Fatalf("expected %v, got %v", expected, compressed)
	}

	compressed, err = DefaultHasher.compress([]byte{1, 2, 3, 4, 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected = []byte{1, 2, 3, 4 ^ 5}
	if !slices.Equal(compressed, expected) {
		t.Fatalf("expected %v, got %v", expected, compressed)
	}

	compressed, err = DefaultHasher.compress([]byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151, 32})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected = []byte{96 ^ 173 ^ 141, 13 ^ 135 ^ 27, 96 ^ 149 ^ 128, 130 ^ 151 ^ 32}
	if !slices.Equal(compressed, expected) {
		t.Fatalf("expected %v, got %v", expected, compressed)
	}
}

func TestHumanize(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	value, err := Humanize(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "alabama-alanine-alaska-alabama"
	if value != expected {
		t.Fatalf("expected %q, got %q", expected, value)
	}

	value, err = Humanize([]byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151, 32})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected = "equal-monkey-lake-double"
	if value != expected {
		t.Fatalf("expected %q, got %q", expected, value)
	}
}

func TestNewUuid(t *testing.T) {
	value, humanized, err := NewUuid()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	compressed := [4]byte{
		value[0] ^ value[1] ^ value[2] ^ value[3],
		value[4] ^ value[5] ^ value[6] ^ value[7],
		value[8] ^ value[9] ^ value[10] ^ value[11],
		value[12] ^ value[13] ^ value[14] ^ value[15],
	}

	humanizedManual := strings.Join(
		[]string{
			DefaultWordList[compressed[0]],
			DefaultWordList[compressed[1]],
			DefaultWordList[compressed[2]],
			DefaultWordList[compressed[3]],
		},
		"-",
	)

	if humanized != humanizedManual {
		t.Fatalf("expected %q, got %q", humanizedManual, humanized)
	}
}
