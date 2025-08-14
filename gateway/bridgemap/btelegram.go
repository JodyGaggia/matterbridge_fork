// +build !notelegram

package bridgemap

import (
	btelegram "github.com/JodyGaggia/matterbridge_fork/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
