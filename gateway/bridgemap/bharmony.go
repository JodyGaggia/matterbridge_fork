//go:build !noharmony
// +build !noharmony

package bridgemap

import (
	bharmony "github.com/JodyGaggia/matterbridge_fork/bridge/harmony"
)

func init() {
	FullMap["harmony"] = bharmony.New
}
