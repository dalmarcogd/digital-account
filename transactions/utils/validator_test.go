package utils

import (
	v1 "github.com/dalmarcogd/digital-account/transactions/handlers/v1"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestCustomValidator_Validate(t *testing.T) {
	a:= v1.transactionsCreateRequest{}
	assert.Error(t, NewCustomValidator(validator.New()).Validate(a))
	a.DocumentNumber = "12345678900"
	assert.NoError(t, NewCustomValidator(validator.New()).Validate(a))
}