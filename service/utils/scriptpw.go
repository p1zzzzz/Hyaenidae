package utils

import (
	"Hyaenidae/global"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

//@function: ScriptPW
//@description: 加密密码
//@param: path string
//@return: base64(scrypt(pw))
func ScriptPW(passwd string) string {
	dk, err := scrypt.Key([]byte(passwd), []byte(global.Hyaenidae_CONFIG.Scrypt.Salt), 32768, 8, 1, 32)
	if err != nil {
		global.Hyaenidae_LOG.Error("encrypt error")
	}
	return base64.StdEncoding.EncodeToString(dk)
}
