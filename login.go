package mud

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/smxlong/mud/ent"
	"github.com/smxlong/mud/ent/player"
	"github.com/smxlong/mud/password"
)

// loginError is an error that occurs during login.
type loginError string

func (e loginError) Error() string {
	return string(e)
}

const (
	// ErrInvalidPassword is returned when the password is invalid.
	ErrInvalidPassword loginError = "invalid password"
)

// login is the login handler
func (s *Service) login(c *gin.Context) {
	token, err := doLogin(c, s.entcli, s.JWTAudience, s.JWTIssuer, []byte(s.JWTSigningKey), c.PostForm("email"), c.PostForm("password"))
	if err != nil {
		s.l.Errorw("login failed", "email", c.PostForm("email"), "error", err)
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	// Set the token in a cookie, valid for 180 days
	c.SetCookie("token", token, 180*24*3600, "", "", false, true)
	returnUrl := c.Query("return")
	if returnUrl == "" {
		returnUrl = "/"
	}
	c.Redirect(302, returnUrl)
}

// logout is the logout handler
func (s *Service) logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(302, "/")
}

// doLogin performs the login.
func doLogin(ctx context.Context, entcli *ent.Client, audience, issuer string, jwtSecret []byte, email, pass string) (string, error) {
	player, err := entcli.Player.Query().Where(player.EmailEQ(email)).Only(ctx)
	if err != nil {
		return "", err
	}
	if !password.Compare(pass, player.Password) {
		return "", ErrInvalidPassword
	}
	return newJWT(email, audience, issuer, jwtSecret, 24*time.Hour)
}

// newJWT returns a complete JWT token.
func newJWT(email, audience, issuer string, jwtSecret []byte, validFor time.Duration) (string, error) {
	t := jwt.New()
	now := time.Now()
	expires := now.Add(validFor)
	_ = t.Set("aud", audience)
	_ = t.Set("exp", expires)
	_ = t.Set("iat", now)
	_ = t.Set("iss", issuer)
	_ = t.Set("nbf", now)
	_ = t.Set("sub", email)
	signed, err := jwt.Sign(t, jwa.HS256, jwtSecret)
	if err != nil {
		return "", err
	}
	return string(signed), nil
}

// verifyPlayer verifies the token and finds the player.
func verifyPlayer(ctx context.Context, entcli *ent.Client, audience, issuer string, jwtSecret []byte, token string) (*ent.Player, error) {
	t, err := jwt.Parse([]byte(token),
		jwt.WithVerify(jwa.HS256, jwtSecret),
		jwt.WithAudience(audience),
		jwt.WithIssuer(issuer),
		jwt.WithValidate(true),
	)
	if err != nil {
		return nil, err
	}
	email := t.Subject()
	player, err := entcli.Player.Query().Where(player.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return player, nil
}
