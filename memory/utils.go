package memory

import (
	"github.com/Peakchen/xcommon/akLog"
	"runtime"
)

func GetMemoryUsage() {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc / MiB
	akLog.FmtPrintln("memory alloc: ", curMem)
}
