package memory

import (
	"Log"
	"runtime"
)

func GetMemoryUsage() {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc / MiB
	Log.FmtPrintln("memory alloc: ", curMem)
}
