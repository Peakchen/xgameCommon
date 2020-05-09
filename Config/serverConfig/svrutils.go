package serverConfig

import (
	"utls"
	"path/filepath"
)

func getserverpath() (path string) {
	exepath := utls.GetExeFilePath()
	path = filepath.Join(exepath, "serverconfig")
	return
}
