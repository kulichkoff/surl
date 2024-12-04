package random

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	t.Run("test length", func(t *testing.T) {
		length := 8
		randomStr := GetRandomString(length)
		if len(randomStr) != length {
			t.Errorf("Expected %d length but got %d", length, len(randomStr))
		}
	})

	t.Run("test randomness", func(t *testing.T) {
		str1 := GetRandomString(8)
		str2 := GetRandomString(8)
		if str1 == str2 {
			t.Errorf("Generated strings are the same: %s == %s", str1, str2)
		}
	})
}
