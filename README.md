# clipsnoop

## Summary

Creates a DLL file that can be used to snoop on the clipboard of a remote process on Windows x64. DLL file should be injected into the remote process to snoop on using one of the many techniques that enable this (CreateRemoteThread, SetWindowsHookEx, QueueUserAPC) or if your feeling very stealthy [Reflective DLL injection](https://github.com/monoxgas/sRDI).

Powered By [Winhook](https://github.com/stavinski/winhook).

## Why

I know that you can monitor the clipboard using `SetClipboardViewer` or `AddClipboardFormatListener` however these techniques are extremely broad, they both are not a stealthy approach and also lead to use capturing everything.

Clipsnoop takes a different approach entirely and instead allows you to target only the specific applications that your interested in snooping on such as Password Managers or Excel spreadsheets containing sensitive data. Clipsnoop also uses an inline hooking approach wereby the actual machine code in memory is changed in the clipboard function when the real function is diverted clipsnoop takes the text content and then allows the real function to continue where it left off. 

## Building

For convenience a `compile.bat` file is provided that can be tweaked for the built DLL. There are several variables that can be set:

* **DLLNAME** - Filename of the built DLL, you may want to disguise this as a legit windows/application DLL (default: sc.dll)
* **DEBUG** - Enables debugging of the hooking code to help diagnose issues. Obviously should only be turned on while troubleshooting local issue. (default: false)
* **LOGPATH** - Full path of the log file to write the captured clipboard text to. You may want to disguise this as a legit looking log file. Make sure that the path is writeable for the executing context. (default:c:\\users\\public\\documents\\ADVAPI32.DAT)

## Testing

I have tried this against a number of x64 windows apps and most have worked fine, however there could be some which could cause issues so it is also ways advised to test locally against the program to which you want to inject against to ensure it doesn't cause instability issues leading to the application crashing!

## Exfil

Clipsnoop currently supports exfiltrating clipboard content using a log file, please note that the content is written in plaintext.

## Limitations

At present there is no way to stop snooping on the application, this may be an issue if the application is left running and the machine is not rebooted in a while. In future I may look to add a kill switch that would at least stop clipsnoop from capturing clipboard data but I don't think there is anyway to unload the DLL and kill the Go runtime.

## Disclaimer

**This code should not be used for any illegal activity and therfore only used to test against systems that you have the explicit permission of the owner to do so. I take no responsibility for any illegal actions performed using any of the source code provided in this repository!**