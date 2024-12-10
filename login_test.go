package mud

import (
	"context"
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/smxlong/mud/ent"
	"github.com/smxlong/mud/password"
	"github.com/stretchr/testify/require"
)

const TEST_EMAIL = "test@example.com"
const TEST_AUDIENCE = "audience"
const TEST_ISSUER = "issuer"

// newEntClient returns a new ent.Client for testing.
func newEntClient(t *testing.T) *ent.Client {
	t.Helper()
	cli, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	require.NoError(t, cli.Schema.Create(context.Background()))
	return cli
}

// Test_doLogin tests the doLogin function.
func Test_doLogin(t *testing.T) {
	cli := newEntClient(t)
	defer cli.Close()
	ctx := context.Background()

	// Create a player
	_, err := cli.Player.Create().
		SetName("test").
		SetEmail(TEST_EMAIL).
		SetPassword(password.Hash("password")).
		Save(ctx)
	require.NoError(t, err)

	// Test a successful login
	token, err := doLogin(ctx, cli, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), TEST_EMAIL, "password")
	require.NoError(t, err)
	require.NotEmpty(t, token)
	// Validate the token
	claims, err := validateJWT(token, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"))
	require.NoError(t, err)
	require.Equal(t, TEST_EMAIL, claims.Subject())

	// Test a failed login
	_, err = doLogin(ctx, cli, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), TEST_EMAIL, "wrong")
	require.Error(t, err)

	// Test a non-existent player
	_, err = doLogin(ctx, cli, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), "bad@example.com", "password")
	require.Error(t, err)
}

// Test_verifyPlayer tests the verifyPlayer function.
func Test_verifyPlayer(t *testing.T) {
	cli := newEntClient(t)
	defer cli.Close()
	ctx := context.Background()

	// Create a player
	_, err := cli.Player.Create().
		SetName("test").
		SetEmail(TEST_EMAIL).
		SetPassword(password.Hash("password")).
		Save(ctx)
	require.NoError(t, err)

	// Create a token
	token, err := newJWT(TEST_EMAIL, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), 24*time.Hour)
	require.NoError(t, err)

	// Test a successful verification
	player, err := verifyPlayer(ctx, cli, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), token)
	require.NoError(t, err)
	require.Equal(t, TEST_EMAIL, player.Email)

	// Test a failed verification
	_, err = verifyPlayer(ctx, cli, TEST_AUDIENCE, TEST_ISSUER, []byte("wrong"), token)
	require.Error(t, err)

	// Now create a token for a non-existent player, which should fail verification
	token, err = newJWT("bad@example.com", TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), 24*time.Hour)
	require.NoError(t, err)
	_, err = verifyPlayer(ctx, cli, TEST_AUDIENCE, TEST_ISSUER, []byte("secret"), token)
	require.Error(t, err)
}

// validateJWT validates a JWT token.
func validateJWT(token, expectedAudience, expectedIssuer string, secret []byte) (jwt.Token, error) {
	t, err := jwt.Parse([]byte(token),
		jwt.WithVerify(jwa.HS256, secret),
		jwt.WithAudience(expectedAudience),
		jwt.WithIssuer(expectedIssuer),
		jwt.WithValidate(true),
	)
	if err != nil {
		return nil, err
	}
	return t, nil
}
