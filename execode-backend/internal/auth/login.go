package auth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
)

// TODO: Make sure we remove all the hardcoded secrets
type BearerToken struct {
	Token string `json:"token"`
}
type JwtClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var (
	signingPrivateKey *rsa.PrivateKey
	signingKeyJWK     jwk.Key
	cognitoPubJWK     []jwk.Key        = make([]jwk.Key, 0, 3)
	cognitoPubKeys    []*rsa.PublicKey = make([]*rsa.PublicKey, 0, 3)
)

func (e Env) InitKeys() error {
	var authCount int

	authCount++
	resp, err := http.Get(e.CognitoJWKURL)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	set, err := jwk.Parse(b)
	if err != nil {
		return err
	}

	for it := set.Iterate(context.Background()); it.Next(context.Background()); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)

		var rawkey interface{} // This is the raw key, like *rsa.PrivateKey or *ecdsa.PrivateKey
		if err := key.Raw(&rawkey); err != nil {
			return err
		}

		pub, ok := rawkey.(*rsa.PublicKey)
		if !ok {
			return errors.New("cannot parse key in JWK set as RSA public key")
		}

		cognitoPubJWK = append(cognitoPubJWK, key)
		cognitoPubKeys = append(cognitoPubKeys, pub)
		fmt.Println(key.KeyID())
	}

	return nil
}

// Returns JWT middleware config. Must be called ater InitKeys
func (e Env) JwtConfig() middleware.JWTConfig {
	signingKeys := make(map[string]interface{})

	for i, pubkey := range cognitoPubKeys {
		signingKeys[cognitoPubJWK[i].KeyID()] = pubkey
	}
	return middleware.JWTConfig{
		// TODO: Fix this
		// SigningKey: &SigningPrivateKey.PublicKey,
		SigningKeys:   signingKeys,
		SigningMethod: "RS256",
		Skipper: func(c echo.Context) bool {
			// Allow GET requests to all endpoints

			if c.Path() == "/login" {
				return true
			}
			if c.Request().Method == "GET" {
				return true
			}

			return false
		},
		Claims: &JwtClaims{},
	}
}

// LoginHandler godoc
// @Summary     Login endpoint
// @Description Authenticates with basic authentication return a JWT token
// @Tags        Authentication
// @Accept      application/json
// @Produce     json
// @Param       loginCredentials body     auth.LoginHandler.request  true "Description of the user to created"
// @Success     200             {object} auth.BearerToken "Describes the created user"
// @Router      /login [post]
// func LoginHandler(c echo.Context) error {
// 	type request struct {
// 		Username string `json:"user_name"`
// 		Password string `json:"password"`
// 	}
// 	body := request{}
// 	if err := c.Bind(&body); err != nil {
// 		return c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: "failed to bind message body"})
// 	}

// 	user, err := entities.GetUserByUsername(body.Username)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(body.Password)); err != nil {
// 		return c.JSON(http.StatusUnauthorized, api.ErrorResponse{Message: "invalid credentials"})
// 	}

// 	currentTime := time.Now()

// 	token := GetAccessToken(user.ID, currentTime)
// 	refreshToken := GetRefreshToken(user.ID, currentTime)

// 	t, err := token.SignedString(signingPrivateKey)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
// 	}
// 	rt, err := refreshToken.SignedString(signingPrivateKey)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
// 	}

// 	c.SetCookie(&http.Cookie{
// 		Name:     "refresh-token",
// 		Value:    rt,
// 		Path:     "/refresh",
// 		Expires:  currentTime.Add(time.Hour * (24 * 14)),
// 		HttpOnly: true,
// 	})

// 	return c.JSON(http.StatusOK, BearerToken{t})
// }

// // RefreshHandler godoc
// // @Summary     Refresh endpoint
// // @Description Provides a new access token given that there is a valid refresh-token cookie
// // @Tags        Authentication
// // @Accept      application/json
// // @Produce     json
// // @Success     200             {object} auth.BearerToken "Describes the created user"
// // @Router      /refresh [get]
// func RefreshHandler(c echo.Context) error {
// 	rtCookie, err := c.Cookie("refresh-token")
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, "failed to get refresh token")
// 	}

// 	rt := rtCookie.Value
// 	token, err := jwt.ParseWithClaims(rt, &JwtClaims{}, ParsingKeyFunc)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, api.ErrorResponse{Message: err.Error()})
// 	}

// 	currentTime := time.Now()
// 	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
// 		token := GetAccessToken(claims.UserID, currentTime)
// 		refreshToken := GetRefreshToken(claims.UserID, currentTime)

// 		t, err := token.SignedString(signingPrivateKey)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
// 		}
// 		rt, err := refreshToken.SignedString(signingPrivateKey)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
// 		}

// 		c.SetCookie(&http.Cookie{
// 			Name:     "refresh-token",
// 			Value:    rt,
// 			Path:     "/refresh",
// 			Expires:  currentTime.Add(time.Hour * (24 * 14)),
// 			HttpOnly: true,
// 		})

// 		return c.JSON(http.StatusOK, BearerToken{t})

// 	} else {
// 		return c.JSON(http.StatusUnauthorized, "invalid or expired refresh token")
// 	}
// }

// func GetRefreshToken(user uint, currentTime time.Time) *jwt.Token {
// 	claims := &JwtClaims{
// 		user,
// 		jwt.StandardClaims{
// 			ExpiresAt: currentTime.Add(time.Hour * 72).Unix(),
// 		},
// 	}

// 	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
// 	refreshToken.Header["kid"] = signingKeyJWK.KeyID()
// 	return refreshToken
// }

// func GetAccessToken(user uint, currentTime time.Time) *jwt.Token {
// 	claims := &JwtClaims{
// 		user,
// 		jwt.StandardClaims{
// 			ExpiresAt: currentTime.Add(time.Hour * 72).Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
// 	token.Header["kid"] = signingKeyJWK.KeyID()
// 	return token
// }

func ParsingKeyFunc(token *jwt.Token) (interface{}, error) {
	return &signingPrivateKey.PublicKey, nil
}
