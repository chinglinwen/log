// Example of turn debug on the fly
//      $ kill -s SIGUSR1 prog
package main

import (
        "os"
        "os/signal"
        "syscall"
        "time"

        "github.com/chinglinwen/log"
)

func main() {
        for {
                log.Println(time.Now())
                log.Debug.Println("debug:", time.Now())
                time.Sleep(1 * time.Second)
        }
}

func init() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, syscall.SIGUSR1)  //SIGUSR1 doesn't work on windows
        go func() {
                for _ = range c {
                        var level string
                        switch log.GetLevel() {
                        case "debug":
                                level = "info"
                        default:
                                level = "debug"
                        }
                        log.SetLevel(level)
                        log.Println("got signal, set log level to ", level)
                }
        }()
}
