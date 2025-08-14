// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/JodyGaggia/matterbridge_fork/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
