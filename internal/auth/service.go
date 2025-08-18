package auth

import "time"

type TokenSrv struct {
	AccessSecret             string
	RefreshSecret            string
	AccessExpiresIn          time.Duration
	RefreshExpiresIn         time.Duration
	RefreshExpiresInRemember time.Duration
	Issuer                   string
}

func NewTokenSrv(accessSecret string, refreshSecret string, accessExpiresIn time.Duration, refreshExpiresIn time.Duration, refreshExpiresInRemember time.Duration, issuer string) *TokenSrv {
	return &TokenSrv{
		AccessSecret:             accessSecret,
		RefreshSecret:            refreshSecret,
		AccessExpiresIn:          accessExpiresIn,
		RefreshExpiresIn:         refreshExpiresIn,
		Issuer:                   issuer,
		RefreshExpiresInRemember: refreshExpiresInRemember,
	}
}

// CreateAccessToken generates a signed access token
func (ts *TokenSrv) CreateAccessToken(userID string, email string, roles []string) (string, int64, error) {
	claims := NewAccessTokenClaim(userID, email, ts.Issuer, roles, ts.AccessExpiresIn)
	token, err := claims.GetToken(ts.AccessSecret)
	if err != nil {
		return "", 0, err
	}
	return token, int64(ts.AccessExpiresIn.Seconds()), nil
}

// ParseAccessToken validates and extracts access token claims
func (ts *TokenSrv) ParseAccessToken(tokenString string) (*AccessTokenClaims, error) {
	return GetAccessTokenFromJWT(tokenString, ts.AccessSecret)
}

// CreateRefreshToken generates a signed refresh token
func (ts *TokenSrv) CreateRefreshToken(userID string, rememberMe bool) (string, string, int64, error) {
	ttl := ts.RefreshExpiresIn
	if rememberMe {
		ttl = ts.RefreshExpiresInRemember
	}

	claims := NewRefreshTokenClaim(userID, ttl)

	token, err := claims.GetToken(ts.RefreshSecret)
	if err != nil {
		return "", "", 0, err
	}
	return claims.ID, token, int64(ttl.Seconds()), nil
}

// ParseRefreshToken validates and extracts refresh token claims
func (ts *TokenSrv) ParseRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	return GetRefreshTokenFromJWT(tokenString, ts.RefreshSecret)
}
