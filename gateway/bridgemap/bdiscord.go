// +build !nodiscord

package bridgemap

import (
	bdiscord "github.com/JodyGaggia/matterbridge_fork/bridge/discord"
)

func init() {
	FullMap["discord"] = bdiscord.New
	UserTypingSupport["discord"] = struct{}{}
}
