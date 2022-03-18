# Build & Use
```
cp .env.example .env
. .env
go build
./minimon
```

# Embedded Postgres Alone
```
go build epg.go
./epg
Crtl-C after started to keep runnig in background
```

# Database migrations
- https://github.com/pressly/goose
```
cd migrations
goose create addtable sql
<Add CREATE/ALTER to up/down>
goose up
```

# REST interfacing with PostgreSQL database
- https://postgrest.org/en/stable/


# HTTP Server if needed
- https://github.com/labstack/echo

Some times it is nice to have a simple http server to add little times.


# Considerations
## Performance
https://levelup.gitconnected.com/fastest-postgresql-client-library-for-go-579fa97909fb

## To ORM or not to ORM(Object Relation Mapper)
Tried GORM which is one of the better Go ORMs but ulitmately it created more headaches than it solved.<br>
It's best to let PostgreSQL be the authority for data/typing.<br>
Using sqlx/pgx, goose(migrations) & postgrest(http rest sql) is better than using GORM.<br>
ORMs just abstract and cause more issues then they solve when working with more complex data.<br>


# Libraries
## Using

## Look at
Libraries to look at and consider, use or more often rewrite but they can give ideas where to start.

- https://github.com/davecheney/httpstat
- https://github.com/davecheney/httpstat/blob/master/main.go
