package coreplugins_test

import (
	"gopher/src/coreplugins"
	"testing"
)

func TestHashPassword(t *testing.T) {

	t.Run("test is hash", func(t *testing.T) {
		text := "12314"
		hash, _ := coreplugins.HashPassword(text)

		if text == hash {
			t.Error("got text equal hash expected text not equal hash")
		}
	})

	t.Run("test correct hash", func(t *testing.T) {
		text := "12314"
		hash, _ := coreplugins.HashPassword(text)
		correct := coreplugins.CheckPasswordHash(text, hash)
		expected := true
		if !correct {
			t.Errorf("got %v expected %v", correct, expected)
		}
	})

	t.Run("test incorrect hash", func(t *testing.T) {
		text := "12314"
		textInCorrect := "dfhfdgdfg"
		hash, _ := coreplugins.HashPassword(text)
		correct := coreplugins.CheckPasswordHash(textInCorrect, hash)
		expected := false
		if correct {
			t.Errorf("got %v expected %v", correct, expected)
		}
	})

}

func BenchmarkHashPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coreplugins.HashPassword("asdasdasdasd")
	}
}
