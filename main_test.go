package main

import (
    "example/hello/pkg/sign/auth"
	"testing"
)


const params = "vk_access_token_settings=&vk_app_id=8228892&vk_are_notifications_enabled=0&vk_is_app_user=1&vk_is_favorite=0&vk_language=ru&vk_platform=mobile_web&vk_ref=other&vk_ts=1673003349&vk_user_id=541919523&sign=b257T1Va8wysq74FAxJqmCwDYvZ7gK4PXGuaXUlcm6o"
const secret = "1nd5Goq7fr4f0XDleABv"

func BenchmarkCheckSign(b *testing.B) {
    for i := 0; i < b.N; i++ {
         auth.Verify(params, "1nd5Goq7fr4f0XDleABv")
         auth.Param("vk_access_token_settings", params)
    }
}
