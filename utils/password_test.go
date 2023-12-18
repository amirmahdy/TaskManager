package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyCreateHash(t *testing.T) {
	pass := "password"
	hash, err := CreateHashPassword(pass)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = VerifyHashPassword(pass, hash)
	require.NoError(t, err)
}
