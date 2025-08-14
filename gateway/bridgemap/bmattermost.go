// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/JodyGaggia/matterbridge_fork/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
