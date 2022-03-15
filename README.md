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

Considerations
```
Using sqlx, goose(migrations) & postgrest(http rest sql) may be a good option but going to try Gorm.
Rewrite in V2 should make it a lot better but we'll see.
```
https://levelup.gitconnected.com/fastest-postgresql-client-library-for-go-579fa97909fb


https://github.com/davecheney/httpstat may be better http stat tool and use https://github.com/davecheney/httpstat/blob/master/main.go
