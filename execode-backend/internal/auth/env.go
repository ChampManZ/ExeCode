package auth

import "os"

const (
	COGNITO_JWK = "COGNITO_JWK_URL"
	REMOTE_JWK  = "SIGNING_JWK_URL"
	SIGNING_JWK = "SIGNING_JWK_FILE"

	AUTH_TYPE = "AUTH_TYPE"
)

type Env struct {
	LocalJWK      string
	LocalJWKURL   string
	CognitoJWKURL string

	AuthType string
}

func GetEnv() (Env, error) {
	return Env{
		LocalJWK:      os.Getenv(SIGNING_JWK),
		LocalJWKURL:   os.Getenv(REMOTE_JWK),
		CognitoJWKURL: os.Getenv(COGNITO_JWK),

		AuthType: os.Getenv(AUTH_TYPE),
	}, nil
}
