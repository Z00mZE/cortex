package hash

import "golang.org/x/crypto/bcrypt"

type BCryptHash struct {
	cost int
}

func NewBCryptHash(cost int) *BCryptHash {
	return &BCryptHash{cost: cost}
}
func (hasher *BCryptHash) GenerateFromPassword(in []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(in, hasher.cost)
}
func (h *BCryptHash) CompareHashAndPassword(hashedPassword, password []byte) bool {
	return bcrypt.ErrMismatchedHashAndPassword == bcrypt.CompareHashAndPassword(hashedPassword, password)
}
