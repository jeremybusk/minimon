Build & Use
```
cp .env.example .env
. .env
go build
./minimon
```

Embedded Postgres Alone
```
go build epg.go
./epg
Crtl-C after started to keep runnig in background
```


Colors
```
const (
        InfoColor    = "\033[1;34m%s\033[0m"
        NoticeColor  = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)
```
