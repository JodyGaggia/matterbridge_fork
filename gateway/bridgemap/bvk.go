// +build !novk

package bridgemap

import (
	bvk "github.com/JodyGaggia/matterbridge_fork/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
