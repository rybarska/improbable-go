package auth

import (
	"time"

	"improbable-module/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
