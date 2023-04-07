package token

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tr0b/simplebank/generator"
)

func createRandomMaker(t *testing.T, NewMaker func(s string) (Maker, error)) Maker {
	secret, err := generator.GenerateSecret()
	if err != nil {
		t.Fatal(err)
	}

	maker, err := NewMaker(secret.Phrase)
	require.NoError(t, err)
	return maker
}
