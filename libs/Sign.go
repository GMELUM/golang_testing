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
	VKIsAppUser               uint32 `json:"vk_is_app_user"`
	VKAreNotificationsEnabled uint32 `json:"vk_are_notifications_enabled"`
	VKGroupId                 uint32 `json:"vk_group_id"`
	VKIsFavorite              uint32 `json:"vk_is_favorite"`
	VKTs                      uint32 `json:"vk_ts"`
	VKIsRecommended           uint32 `json:"vk_is_recommended"`
	VKProfileId               uint32 `json:"vk_profile_id"`
	VKHasProfileButton        uint32 `json:"vk_has_profile_button"`
	VKTestingGroupId          uint32 `json:"vk_testing_group_id"`
	ODREnabled                uint32 `json:"odr_enabled"`
	VKViewerGroupRole         string `json:"vk_viewer_group_role"`
	VKPlatform                string `json:"vk_platform"`
	VKLanguage                string `json:"vk_language"`
	VKRef                     string `json:"vk_ref"`
	VKAccessTokenSettings     string `json:"vk_access_token_settings"`
	VKChatId                  string `json:"vk_chat_id"`
	Sign                      string `json:"sign"`
}

func Sign(value string, secret string) (VKParams, error) {

	var params = VKParams{}
	var err = json.Unmarshal([]byte(value), &params)
	if err != nil {
		return params, err
	}

	var query string
	var regex = *regexp.MustCompile(`"(vk_[\_A-Za-z]+)":"?([0-9a-zA-Z_]*)`)
	for _, match := range regex.FindAllStringSubmatch(value, -1) {
		query += (match[1] + "=" + match[2] + "&")
	}

	var digest = hmac.New(sha256.New, []byte(secret))
	digest.Write([]byte(query[:len(query)-1]))
	var checkSign = base64.URLEncoding.EncodeToString(digest.Sum(nil))

	if params.Sign != checkSign[:len(params.Sign)] {
		return params, fmt.Errorf("Invalid params")
	}

	return params, nil

}
