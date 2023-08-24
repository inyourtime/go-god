package coreplugins_test

import (
	"gopher/src/coreplugins"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

type MockConfig struct {
    JwtSecret string
}

func (m *MockConfig) Load() {
    // Provide mock values for JwtSecret and other configuration options
    m.JwtSecret = "asfsgergergerdsfsd"
}

func TestJwtToken(t *testing.T) {
	t.Run("signed token success", func(t *testing.T) {
		config := MockConfig{}
		config.Load()
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   float64(time.Now().Add(time.Hour * 72).Unix()),
		}
		token, _ := coreplugins.Token(claims, config.JwtSecret)
		assert.NotEmpty(t, token)
	})

	t.Run("decoded token success", func(t *testing.T) {
		config := MockConfig{}
		config.Load()
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   float64(time.Now().Add(time.Hour * 72).Unix()),
		}
		token, _ := coreplugins.Token(claims, config.JwtSecret)
		c, _ := coreplugins.DecodedToken(token, config.JwtSecret)
		assert.Exactly(t, claims, c)
	})

	t.Run("test token expired", func(t *testing.T) {
		config := MockConfig{}
		config.Load()
		expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTY5MjgwNTczMCwibmFtZSI6IkdvZCJ9.TZpHVS2yBYjv_trrS_WdZUgVA7MwnFbxRbWfz930HEE"
		_, err := coreplugins.DecodedToken(expiredToken, config.JwtSecret)
		assert.ErrorIs(t, err, jwt.ErrTokenExpired)
	})
	
	t.Run("test invalid signature", func(t *testing.T) {
		config := MockConfig{}
		config.Load()
		invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjkyODAzMjA4LCJuYW1lIjoiSm9obiBEb2UifQ.SNrr_DxyjESSkMQNkI4qODo-csjBazIgj2PkZlGz90s"
		_, err := coreplugins.DecodedToken(invalidToken, config.JwtSecret)
		assert.ErrorIs(t, err, jwt.ErrTokenSignatureInvalid)
	})
}
