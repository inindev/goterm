package main

// taken from: https://github.com/tinygo-org/tinygo/blob/release/monitor.go

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mattn/go-tty"
	"go.bug.st/serial"
)

func Monitor(port string, baud int) error {
	if baud < 1 {
		baud = 115200
	}

	fmt.Printf("connecting %s %d\n", port, baud)

	wait := 300
	var err error
	var p serial.Port
	for i := 0; i <= wait; i++ {
		p, err = serial.Open(port, &serial.Mode{BaudRate: baud})
		if err != nil {
			if i < wait {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			return err
		}
		p.ResetInputBuffer()
		p.ResetOutputBuffer()
		break
	}
	defer p.Close()

	tty, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGQUIT)
	defer signal.Stop(sig)

	go func() {
		for {
			sig := <-sig
			switch sig {
			case os.Interrupt:
				p.Write([]byte{0x1b, 0x03}) // send ctrl+c to tty
			case syscall.SIGQUIT:
				p.Write([]byte{0x1b, 0x1c}) // send ctrl+\ to tty
			}
		}
	}()

	fmt.Printf("%s connected, use ctrl+] to exit\n", port)

	errCh := make(chan error, 1)

	go func() {
		buf := make([]byte, 100*1024)
		for {
			n, err := p.Read(buf)
			if err != nil {
				errCh <- fmt.Errorf("read error: %w", err)
				return
			}

			if n == 0 {
				continue
			}

			fmt.Printf("%v", string(buf[:n]))
		}
	}()

	go func() {
		for {
			r, err := tty.ReadRune()
			if err != nil {
				errCh <- err
				return
			}

			if r == 0 {
				continue
			}

			if r == 29 { // ctrl+]
				fmt.Println("ctrl+] received, exiting...")
				tty.Close()
				os.Exit(0)
			}

			p.Write([]byte(string(r)))
		}
	}()

	return <-errCh
}
