package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"

	"ecmn/config"
	"ecmn/logger"

	"github.com/gin-gonic/gin"
)

const (
	SignatureHeader = "X-Ech0-Signature"
)

func SignatureVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.Get().Webhook.Secret
		if secret == "" {
			logger.Warn("Webhook secret is not configured, skipping signature verification")
			c.Next()
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error("Failed to read request body", logger.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(io.NopCloser(bytesBufferPool(body)))

		signature := c.GetHeader(SignatureHeader)

		if signature == "" {
			logger.Warn("Missing signature header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing signature"})
			c.Abort()
			return
		}

		if !verifySignature(secret, body, signature) {
			logger.Warn("Invalid signature")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			c.Abort()
			return
		}

		c.Set("rawBody", body)
		c.Next()
	}
}

func verifySignature(secret string, body []byte, signature string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expected := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(expected), []byte(signature))
}

type nopCloser struct {
	io.Reader
}

func (n nopCloser) Close() error {
	return nil
}

func bytesBufferPool(body []byte) io.Reader {
	return nopCloser{bytes.NewReader(body)}
}
