/*
	Verify a query parameter from the launch url of vk mini apps.
	Input: query string and the secret key obtained in the application settings.
	Outputs: boolean value depending on the result of the verify

	params := "key=value&key=value&sign=value"
	secret := "wvl68m4dR1UpLrVRli"

	valid := Verify(params, secret)
	if !valid {
		fmt.Println("Unauthorized")
	}
*/

package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func Verify(params string, secret string) bool {

	index := strings.LastIndex(params, "&")
	query := params[:index]
	sign := params[index+6:]

	digest := hmac.New(sha256.New, []byte(secret))
	_, _ = digest.Write([]byte(query))
	checkSign := base64.URLEncoding.EncodeToString(digest.Sum(nil))

	return sign == checkSign[:len(sign)]

}
