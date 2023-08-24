package coreplugins_test

import (
	"gopher/src/coreplugins"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {

	t.Run("test is hash", func(t *testing.T) {
		text := "12314"
		hash, _ := coreplugins.HashPassword(text)
		assert.NotEqual(t, text, hash)
	})

	t.Run("test correct hash", func(t *testing.T) {
		text := "12314"
		hash, _ := coreplugins.HashPassword(text)
		result := coreplugins.CheckPasswordHash(text, hash)
		expected := true
		assert.Equal(t, expected, result)
	})

	t.Run("test incorrect hash", func(t *testing.T) {
		text := "12314"
		textInCorrect := "dfhfdgdfg"
		hash, _ := coreplugins.HashPassword(text)
		result := coreplugins.CheckPasswordHash(textInCorrect, hash)
		expected := false
		assert.Equal(t, expected, result)
	})

}

func BenchmarkHashPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coreplugins.HashPassword("asdasdasdasd")
	}
}
