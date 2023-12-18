package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secret := "changethissecretkeyforproduction"
	username1 := "testuser1"
	jwtMaker, err := NewJWTMaker(secret)
	require.NoError(t, err)
	token, payload, err := jwtMaker.CreateToken(username1, time.Duration(1*time.Minute))
	require.NoError(t, err)

	payloadVerfied, err := jwtMaker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, payloadVerfied.Username, payload.Username)
	require.Equal(t, payloadVerfied.Username, username1)
}
