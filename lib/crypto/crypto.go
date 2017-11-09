package crypto

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/lucashenriqueteixeira/list_go/lib/util"
)

//create hash
func CryptoNew(password string) ([]byte) {

	Result, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	util.VerifyErr(err)

	return Result
}


//compare hash com do user
func Comparehash(userPassword string, dbPassword []byte) (bool) {

	if result := bcrypt.CompareHashAndPassword(dbPassword, []byte(userPassword)); result != nil {
		return false
	} else {

	return true
	}
}