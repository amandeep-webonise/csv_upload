package envprovider

import (
	"github.com/caarlos0/env"
)

// EnvProvider implements environment mapping methods
type EnvProvider interface {
	Parse(v interface{}) error
}

// EnvConfigProvider contains empty struct
type EnvConfigProvider struct{}

// Parse parses structure and checks for a value using tag "env"
func (v *EnvConfigProvider) Parse(s interface{}) error {
	if err := env.Parse(s); err != nil {
		return err
	}

	return nil
}
