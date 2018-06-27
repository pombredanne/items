package password

import (
	"strconv"
	"fmt"
	"strings"
	"encoding/hex"
	"math/rand"
	"crypto"
	"time"
)

func SetPassword(loginPwd string) string {
	algoList := []string{"MD5", "SHA1", "SHA256"}
	rand.Seed(time.Now().Unix())
	algo := algoList[rand.Intn(3)]
	salt := cryptoString(algo, strconv.Itoa(rand.Intn(999999)), strconv.Itoa(rand.Intn(999999)))[:5]
	hsh := cryptoString(algo, salt, loginPwd)
	return fmt.Sprintf("%s$%s$%s", algo, salt, hsh)

}
func CheckPassword(loginPwd, dbPwd string) bool {
	if !strings.Contains(dbPwd, "$") {
		return false
	}
	list := strings.Split(dbPwd, "$")
	if len(list) != 3 {
		return false
	}
	algo, salt, hsh := list[0], list[1], list[2]
	return hsh == cryptoString(algo, salt, loginPwd)
}
func cryptoString(algo, salt, loginPwd string) string {
	s := salt + loginPwd
	switch algo {
	case "MD5":
		h := crypto.MD5.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	case "SHA1":
		h := crypto.SHA1.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	case "SHA256":
		h := crypto.SHA256.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	}
	return ""
}
