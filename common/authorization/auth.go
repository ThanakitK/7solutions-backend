package authorization

type AppAuthorization interface {
	// สำหรับ Cenerate JWT Tokan
	GenerateToken(payload AppAuthorizationClaim) (token string, err error)

	// สำหรับ Validate JWT Tokan
	ValidateToken(tokenString string, paserTo interface{}) (err error)
}
