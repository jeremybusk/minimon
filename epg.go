package main

/*
https://pkg.go.dev/github.com/fergusstrange/embedded-postgres#section-readme
https://github.com/fergusstrange/embedded-postgres
https://github.com/zonkyio/embedded-postgres
https://mvnrepository.com/artifact/io.zonky.test.postgres/embedded-postgres-binaries-bom
*/

import (
        "bytes"
        "fmt"
        "os"
        "strconv"
        "time"

        "github.com/drael/GOnetstat"
        "github.com/fergusstrange/embedded-postgres"
        // "sync"
)

// var wg sync.WaitGroup

func main() {
        port := 9999
        // var port64 int64
        // port64 = int64(port)
        // killPort(port64)
        runseconds := 60
        var port32 uint32
        port32 = uint32(port)
        // Version("14.2.0").
        // Version("13.6.0").
        // Version("12.10.0").
        logger := &bytes.Buffer{}
        conf := embeddedpostgres.DefaultConfig().
                Username("minimon").
                Password("minimon").
                Database("minimon").
                // Version("14.2.0").
                Version("12.10.0").
                RuntimePath("./data/runtime").
                BinariesPath("./data/bin").
                DataPath("./data/pgdata").
                BinaryRepositoryURL("https://repo1.maven.org/maven2").
                Port(port32).
                StartTimeout(15 * time.Second).
                Logger(logger)

        postgres := embeddedpostgres.NewDatabase(conf)
        err := postgres.Start()
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        fmt.Printf("Running postgres server for %d seconds.\n", runseconds)
        time.Sleep(time.Duration(runseconds) * time.Second)

        err = postgres.Stop()
        if err != nil {
                fmt.Println(err)
        }
}

func killPort(port int64) {
        // defer wg.Done()
        d := GOnetstat.Tcp()
        for _, p := range d {
                if p.State == "LISTEN" && p.Port == port {
                        fmt.Printf("Killing port, pid, name %v %v %v\n", port, p.Pid, p.Name)
                        pid, err := strconv.Atoi(p.Pid)
                        if err != nil {
                                panic(err.Error())
                        }
                        process, err := os.FindProcess(pid)
                        if err != nil {
                                panic(err.Error())
                        }
                        process.Signal(os.Interrupt)
                        fmt.Printf("Waiting 10 seconds for port to close.\n")
                        time.Sleep(time.Duration(10) * time.Second)
                }
        }
}


// NOTES ============================

func killPID(pid int) {
        process, err := os.FindProcess(pid)
        if err != nil {
                panic(err.Error())
        }
        process.Signal(os.Interrupt)
}

func xxkillPort(port int64) {
        d := GOnetstat.Tcp()
        for _, p := range d {
                if p.State == "LISTEN" && p.Port == port {
                        fmt.Printf("Killing port, pid, name %v %v %v\n", port, p.Pid, p.Name)
                        pid, err := strconv.Atoi(p.Pid)
                        if err != nil {
                                panic(err.Error())
                        }
                        killPID(pid)
                }
        }
}
