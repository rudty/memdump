// +build windows

package memdump

import (
	"os"
	"syscall"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modDbghelp  = syscall.NewLazyDLL("Dbghelp.dll")
)
var (
	procGetCurrentProcess   = modkernel32.NewProc("GetCurrentProcess")
	procGetCurrentProcessID = modkernel32.NewProc("GetCurrentProcessId")
	procGetCurrentThreadID  = modkernel32.NewProc("GetCurrentThreadId")
	procMiniDumpWriteDump   = modDbghelp.NewProc("MiniDumpWriteDump")
)

// getCurrentProcess WinApi HANDLE GetCurrentProcess()
func getCurrentProcess() syscall.Handle {
	r0, _, e1 := syscall.Syscall(procGetCurrentProcess.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		return syscall.Handle(0)
	}
	return syscall.Handle(r0)
}

// getCurrentProcessID WinApi DWORD GetCurrentProcessId()
func getCurrentProcessID() uint32 {
	r0, _, e1 := syscall.Syscall(procGetCurrentProcessID.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		return 0
	}
	return uint32(r0)
}

// miniDumpWriteDump WinApi BOOL MiniDumpWriteDump(...)
func miniDumpWriteDump(process syscall.Handle, pid uint32, fd uintptr, miniDumpType int32) bool {
	r0, _, e1 := syscall.Syscall9(procMiniDumpWriteDump.Addr(), 7, uintptr(process), uintptr(pid), fd, uintptr(miniDumpType), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return false
	}
	return uint(r0) != 0
}

// WriteFullDumpFd write memory dump file
// Task Manager => "Create dump file"
func WriteFullDumpFd(fd uintptr) bool {
	ps := getCurrentProcess()
	pid := getCurrentProcessID()
	return miniDumpWriteDump(ps, pid, fd, 2)
}

// WriteFullDump write memory dump file
// Task Manager => "Create dump file"
func WriteFullDump(f *os.File) bool {
	return WriteFullDumpFd(f.Fd())
}
