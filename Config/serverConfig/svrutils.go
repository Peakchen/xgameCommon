package serverConfig

import (
	"path/filepath"

	"github.com/Peakchen/xgameCommon/utls"
)

func getserverpath() (path string) {
	exepath := utls.GetExeFilePath()
	path = filepath.Join(exepath, "serverconfig")
	return
}
