package authorization

import (
	"7solutions/backend/config"
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	Name    string `json:"name,omitempty"`
	Channel string `json:"channel,omitempty"`
	jwt.StandardClaims
}

type AppAuthorizationClaim struct {
	UserId   string `json:"sub,omitempty"`
	Name     string `json:"name,omitempty"`
	Audience string `json:"aud,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

// NOTE Adapter -----------------------------
type jwtHS256 struct {
	Signature string        `json:"signature"`
	Duration  time.Duration `json:"duration"`
}

// JWT แบบ HS256
func NewJWT_HS256() AppAuthorization {
	return jwtHS256{Signature: config.Env.SignatureKey, Duration: config.Env.SignatureExp}
}

func (c jwtHS256) GenerateToken(payload AppAuthorizationClaim) (tokenString string, err error) {
	// FIX EDIT PAYLOAD HERE ----------------------------
	claim := &authCustomClaims{
		payload.Name,
		payload.Channel,
		jwt.StandardClaims{
			Audience:  payload.Audience,                  // aud Audience (who or what the token intended for)
			ExpiresAt: time.Now().Add(c.Duration).Unix(), // exp Expiration time (seconds since Unix epoch)
			Id:        "",                                // jti JWT ID (unique identifier for this token)
			IssuedAt:  time.Now().Unix(),                 // iat isused at (seconds since Unix epoch)
			Issuer:    payload.Issuer,                    // iss issuer (who created and signed this token)
			NotBefore: 0,                                 // nbf No valid before (seconds since Unix epoch)
			Subject:   payload.UserId,                    // sub Subject (whom the token reference to)
		},
	}
	// FIX EDIT PAYLOAD HERE ----------------------------

	// NOTE Create a new JWT token & Set the claims for the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// NOTE Sign the token with the key
	tokenString, err = token.SignedString([]byte(c.Signature))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c jwtHS256) ValidateToken(tokenString string, data interface{}) (err error) {
	// NOTE Parse the token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// NOTE Check the signing method of the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		// NOTE Return the key for verifying the signature
		return []byte(c.Signature), nil
	})
	if err != nil {
		return err
	}
	// NOTE Check if the token is valid
	if !token.Valid {
		return errors.New("invalid token")
	}

	// NOTE Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	// NOTE Check expired from the token
	expirationTime, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid expiration time")
	}

	if time.Now().Unix() > int64(expirationTime) {
		return errors.New("token has expired")
	}

	// NOTE Check issuer from the token
	_, ok = claims["iss"].(string)
	if !ok {
		return errors.New("invalid issuer in token")
	}

	// NOTE Convert the data to the specified type
	dataBytes, err := json.Marshal(claims)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return err
	}

	return nil
}
