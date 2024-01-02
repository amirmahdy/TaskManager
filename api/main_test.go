package api

import (
	"fmt"
	"net/http"
	"os"
	db "taskmanager/db/model"
	"taskmanager/token"
	"taskmanager/utils"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func addAuthorization(t *testing.T, request *http.Request, tokenMaker token.Maker, tokenType, username string, duration time.Duration) {
	token, _, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", tokenType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.CreateRandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
