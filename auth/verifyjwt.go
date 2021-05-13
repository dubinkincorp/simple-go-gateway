package auth

import "github.com/dgrijalva/jwt-go"

func VerifyWithSecret(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//TODO move to configuration file
		return []byte("123"), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
