// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/JodyGaggia/matterbridge_fork/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
