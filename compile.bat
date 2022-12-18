set DEBUG=false
set LOGPATH=c:\\users\\public\\documents\\ADVAPI32.DAT
go build -buildmode=c-shared -ldflags="-w -s -H=windowsgui -X 'main.debug=%DEBUG%' -X 'main.logpath=%LOGPATH%'" -trimpath -o cs.dll