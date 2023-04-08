package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	db "github.com/tr0b/simplebank/db/sqlc"
	"github.com/tr0b/simplebank/generator"
	"github.com/tr0b/simplebank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	secret, err := generator.GenerateSecret()
	require.NoError(t, err)

	config := util.Config{
		TokenSymmetricKey:   secret.Phrase,
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server

}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
