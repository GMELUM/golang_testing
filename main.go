package main

import (
	"example/hello/pkg/sign/auth"
	"fmt"
)

// "fmt"
// "strings"
// "github.com/SevereCloud/vksdk/vkapps"

func main() {

	const params = "vk_access_token_settings=&vk_app_id=8228892&vk_are_notifications_enabled=0&vk_is_app_user=1&vk_is_favorite=0&vk_language=ru&vk_platform=mobile_web&vk_ref=other&vk_ts=1673003349&vk_user_id=541919523&sign=b257T1Va8wysq74FAxJqmCwDYvZ7gK4PXGuaXUlcm6o"
	const secret = "1nd5Goq7fr4f0XDleABv"

	valid := auth.Verify(params, "1nd5Goq7fr4f0XDleABv")
	if !valid {
		fmt.Println("Unauthorized")
	}

	value := auth.Param(auth.VKAppId, params)

	fmt.Println(value)

	// 	fmt.Println(result.VKUserId)
	

	// fmt.Println(params.VKAppId)

	// link := "https://example.com/?vk_user_id=494075&vk_app_id=6736218&vk_is_app_user=1&vk_are_notifications_enabled=1&vk_language=ru&vk_access_token_settings=&vk_platform=android&sign=htQFduJpLxz7ribXRZpDFUH-XEUhC9rBPTJkjUFEkRA"

  // v, _ := vkapps.ParamsVerify(link, "wvl68m4dR1UpLrVRli")

  // fmt.Println(v)

}

// CLOSE - Закрыт, не может быть востановлен.
// READ - Закрыт, ожидает синхронизации.
// OPEN - Открыт, можно выполнить.