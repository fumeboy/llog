package example

import (
	"os"
)

type PrinterConsole struct {
}

func (p *PrinterConsole) Open() error {
	return nil
}

func (p *PrinterConsole) Write(msg []byte) {
	os.Stdout.Write(msg)
}

func (p *PrinterConsole) Close() {
}