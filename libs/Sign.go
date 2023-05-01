package libs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
)

type VKParams struct {
	VKUserId                  uint32 `json:"vk_user_id"`
	VKAppId                   uint32 `json:"vk_app_id"`
	VKChatId                  string `json:"vk_chat_id"`
	VKIsAppUser               uint32 `json:"vk_is_app_user"`
	VKAreNotificationsEnabled uint32 `json:"vk_are_notifications_enabled"`
	VKLanguage                string `json:"vk_language"`
	VKRef                     string `json:"vk_ref"`
	VKAccessTokenSettings     string `json:"vk_access_token_settings"`
	VKGroupId                 uint32 `json:"vk_group_id"`
	VKViewerGroupRole         string `json:"vk_viewer_group_role"`
	VKPlatform                string `json:"vk_platform"`
	VKIsFavorite              uint32 `json:"vk_is_favorite"`
	VKTs                      uint32 `json:"vk_ts"`
	VKIsRecommended           uint32 `json:"vk_is_recommended"`
	VKProfileId               uint32 `json:"vk_profile_id"`
	VKHasProfileButton        uint32 `json:"vk_has_profile_button"`
	VKTestingGroupId          uint32 `json:"vk_testing_group_id"`
	ODREnabled                uint32 `json:"odr_enabled"`
	Sign                      string `json:"sign"`
}

func Sign(value string, secret string) (VKParams, error) {

	var params = VKParams{}
	err := json.Unmarshal([]byte(value), &params)
	if err != nil {
		return params, err
	}

	var query string
	var regex = *regexp.MustCompile(`"(vk_[\_A-Za-z]+)":"?([0-9a-zA-Z_]*)`)
	for _, match := range regex.FindAllStringSubmatch(value, -1) {
		query += (match[1] + "=" + match[2] + "&")
	}

	digest := hmac.New(sha256.New, []byte(secret))
	digest.Write([]byte(query[:len(query)-1]))
	checkSign := base64.URLEncoding.EncodeToString(digest.Sum(nil))

	if params.Sign != checkSign[:len(params.Sign)] {
		return params, fmt.Errorf("Invalid params")
	}

	return params, nil

}
