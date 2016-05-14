func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	redisConn := redis.Connect()
	expiration := backend.getTokenRemainingValidity(token.Claims["exp"])
	return redisConn.SetValue(tokenString, tokenString, expiration)
}

func (backend *JWTAuthenticationBackend) IsInBlackList(token string) bool {
	redisConn := redis.Connect()
	redisToken, _ := redisConn.GetValue(token)

	if redisToken == nil {
		return false
	}
	return true
}
