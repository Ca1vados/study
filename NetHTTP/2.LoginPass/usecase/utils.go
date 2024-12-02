package usecase

import (
	"crypto/sha256"
	"encoding/hex"
)

func StrToSha256Str(pass string) string {
	// измени на проверку - u.PassHash == sha256(pass)
	passHash := sha256.Sum256([]byte(pass))           // pass поступающий в функцию перевел в хэш и на вывод поставил сравнение с PassHash
	passHashString := hex.EncodeToString(passHash[:]) // переводим byte в string
	return passHashString
}
