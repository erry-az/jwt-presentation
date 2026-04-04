package models

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims represents standard and custom claims for the presentation
type CustomClaims struct {
	Name  string        `json:"name,omitempty"`
	Admin bool          `json:"admin,omitempty"`
	Confirmation *Confirmation `json:"cnf,omitempty"`
	jwt.RegisteredClaims
}

// Confirmation represents the DPoP key binding claim
type Confirmation struct {
	JKT string `json:"jkt"`
}

// Validate satisfies the validator.CustomClaims interface.
func (c *CustomClaims) Validate(ctx context.Context) error {
	return nil
}
