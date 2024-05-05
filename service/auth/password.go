package auth


import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	cost := bcrypt.DefaultCost // higher cost -> more secure, compute slower

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, cost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}