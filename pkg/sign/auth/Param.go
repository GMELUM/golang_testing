/*
	Param user to get value from query string.
	The function expects that the query string passed the
	validation and the entered data is correct.
	Input: param key and query string.
	Output: string or empty string.

	UserID := auth.Param("vk_user_id", params)
	if len(UserID) == 0 {
		fmt.Println("Param is empty")
	}
*/

package auth

import "strings"

type Params string

const (
	VKUserId                  Params = `vk_user_id`
	VKAppId                   Params = `vk_app_id`
	VKIsAppUser               Params = `vk_is_app_user`
	VKAreNotificationsEnabled Params = `vk_are_notifications_enabled`
	VKGroupId                 Params = `vk_group_id`
	VKIsFavorite              Params = `vk_is_favorite`
	VKTs                      Params = `vk_ts`
	VKIsRecommended           Params = `vk_is_recommended`
	VKProfileId               Params = `vk_profile_id`
	VKHasProfileButton        Params = `vk_has_profile_button`
	VKTestingGroupId          Params = `vk_testing_group_id`
	VKViewerGroupRole         Params = `vk_viewer_group_role`
	VKPlatform                Params = `vk_platform`
	VKLanguage                Params = `vk_language`
	VKRef                     Params = `vk_ref`
	VKAccessTokenSettings     Params = `vk_access_token_settings`
	VKChatId                  Params = `vk_chat_id`
)

func Param(key Params, query string) string {

	firstIndex := strings.Index(query, string(key))
	if firstIndex == -1 {
		return ""
	}

	lastIndex := strings.Index(query[firstIndex:], "&")
	if lastIndex == -1 {
		lastIndex = len(key)
	}

	return query[firstIndex+len(key)+1 : lastIndex+firstIndex]

}
