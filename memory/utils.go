package memory

import (
	"akLog"
	"runtime"
)

func GetMemoryUsage() {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc / MiB
	akLog.FmtPrintln("memory alloc: ", curMem)
}
