package serverConfig

import (
	"fmt"
	"path/filepath"

	"github.com/Peakchen/xgameCommon/Config"
)

/*
	export from serverGlobalConfig.json by tool.
*/
type TServerglobalconfigBase struct {
	Id    int32  `json:"id"`
	Value string `json:"Value"`
}

type TServerglobalconfig struct {
	data []*TServerglobalconfigBase
}

type tArrServerglobalconfig []*TServerglobalconfigBase

var (
	GServerglobalconfig *TServerglobalconfig = &TServerglobalconfig{}
)

func init() {
	akLog.FmtPrintln("load	serverGlobalConfig.json")
}

func loadServerglobalconfig() {
	var (
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	}
	path = filepath.Join(SvrPath, "serverGlobalConfig.json")
	Config.ParseJson2Cache(GServerglobalconfig, &tArrServerglobalconfig{}, path)
}

func (this *TServerglobalconfig) ComfireAct(data interface{}) (errlist []string) {
	errlist = []string{}
	cfg := data.(*tArrServerglobalconfig)
	for _, item := range *cfg {
		if len(item.Value) == 0 {
			errlist = append(errlist, fmt.Sprintf("serverGlobalConfig value invalid, id: %v.", item.Id))
		}
	}
	return
}

func (this *TServerglobalconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrServerglobalconfig)
	this.data = []*TServerglobalconfigBase{}
	for _, item := range *cfg {
		this.data = append(this.data, item)
	}
	return
}

func (this *TServerglobalconfig) Get(idx int) *TServerglobalconfigBase {
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}
