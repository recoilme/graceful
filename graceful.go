package graceful

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Exit codes. Generally, you should NOT
// automatically restart the process if the
// exit code is ExitCodeFailedStartup (1).
const (
	ExitCodeSuccess = iota
	ExitCodeFailedStartup
	ExitCodeForceQuit
	ExitCodeFailedQuit
)

// Terminate signals array (ctrl+C, kill, close terminal, quit)
var Terminate = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT}

// Unignore will ignore all signals except passed
func Unignore(quit chan os.Signal, fallback func() error, sig ...os.Signal) {
	signal.Notify(quit)
	go func() {
		for {
			signal := <-quit
			//fmt.Println(time.Now(), "signal:", signal)
			for _, s := range sig {
				if s == signal {
					if fallback != nil {
						err := fallback()
						if err != nil {
							fmt.Println(time.Now(), "fallback error:", err)
							os.Exit(ExitCodeFailedQuit)
						}
					}
					os.Exit(ExitCodeFailedStartup)
				}
			}
		}
	}()
}
