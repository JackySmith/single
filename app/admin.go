package app

import (
	"runtime/debug"

	"github.com/JackySmith/single/utils"
)

func ReloadConfig() {
	debug.FreeOSMemory()
	utils.LoadConfig(utils.CfgFileName)
}
