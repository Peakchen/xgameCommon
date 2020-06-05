package serverConfig

import (
	"fmt"
	"path/filepath"

	"github.com/Peakchen/xgameCommon/Config"
	"github.com/Peakchen/xgameCommon/akLog"
)

/*
	export from centerConfig.json by tool.
*/
type TCenterconfigBase struct {
	Id         int32  `json:"id"`
	No         int32  `json:"No"`
	Listenaddr string `json:"ListenAddr"`
	Pprofaddr  string `json:"PProfAddr"`
	Zone       string `json:"Zone"`
}

type TCenterconfig struct {
	TCenterconfigBase
	Name string
}

type TCenterconfigConfig struct {
	data []*TCenterconfig
}

type tArrCenterconfig []*TCenterconfigBase

var (
	GCenterconfigConfig *TCenterconfigConfig = &TCenterconfigConfig{}
	cstCenterDef                             = "CenterGate"
)

func init() {
	akLog.FmtPrintln("load	centerConfig.json")
}

func loadCenterconfig() {
	var (
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	}
	path = filepath.Join(SvrPath, "centerConfig.json")
	Config.ParseJson2Cache(GCenterconfigConfig, &tArrCenterconfig{}, path)
}

func (this *TCenterconfigConfig) ComfireAct(data interface{}) (errlist []string) {
	errlist = []string{}
	cfg := data.(*tArrCenterconfig)
	for _, item := range *cfg {
		if len(item.Listenaddr) == 0 {
			errlist = append(errlist, fmt.Sprintf("centerConfig listeraddr invalid, id: %v.", item.Id))
		}

		if len(item.Zone) == 0 {
			errlist = append(errlist, fmt.Sprintf("centerConfig Zone invalid, id: %v.", item.Id))
		}

		if len(item.Pprofaddr) == 0 {
			errlist = append(errlist, fmt.Sprintf("centerConfig Pprofaddr invalid, id: %v.", item.Id))
		}
	}
	return
}

func (this *TCenterconfigConfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrCenterconfig)
	this.data = []*TCenterconfig{}
	for _, item := range *cfg {
		new := &TCenterconfig{}
		new.Id = item.Id
		new.Listenaddr = item.Listenaddr
		new.No = item.No
		new.Pprofaddr = item.Pprofaddr
		new.Zone = item.Zone
		new.Name = cstCenterDef
		this.data = append(this.data, new)
	}
	return
}

func (this *TCenterconfigConfig) Get(idx int) *TCenterconfig {
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}
