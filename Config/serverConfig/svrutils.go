package serverConfig

import (
	"github.com/Peakchen/xgameCommon/utls"
	"path/filepath"
)

func getserverpath() (path string) {
	exepath := utls.GetExeFilePath()
	path = filepath.Join(exepath, "serverconfig")
	return
}
