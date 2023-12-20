package api

import (
	"os"
	db "taskmanager/db/model"
	"taskmanager/utils"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

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
