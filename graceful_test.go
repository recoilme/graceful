package graceful

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestInterupt(t *testing.T) {
	quit := make(chan os.Signal, 1)
	Unignore(quit, fallback, Terminate...)
	time.Sleep(1 * time.Second)
	//must be ignored
	quit <- syscall.SIGPIPE
	time.Sleep(1 * time.Second)

	//must be catched
	quit <- syscall.SIGINT
	time.Sleep(1 * time.Second)
}

func TestNil(t *testing.T) {
	quit := make(chan os.Signal, 1)
	Unignore(quit, nil, syscall.SIGINT, syscall.SIGQUIT)
}

func fallback() error {
	time.Sleep(1 * time.Second)
	fmt.Println("fallback")
	return nil
}
