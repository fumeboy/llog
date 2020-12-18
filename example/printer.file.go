package example

import (
	"fmt"
	"os"
	"time"
)

type PrinterFile struct {
	file *os.File
	accessFile *os.File
	lastSplitTime int
}
func (p *PrinterFile) Open() error {
	file, err := os.OpenFile(p.filename(time.Now()), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	p.file = file
	return nil
}
func (p *PrinterFile) Write(msg []byte) {
	p.ifSplit()
	p.file.Write(msg)
}

func (p *PrinterFile) Close() {
	p.file.Close()
}


func (p *PrinterFile) filename(t time.Time) string {
	return fmt.Sprintf("%s.%04d%02d%02d%02d", "./logs/log",
		t.Year(), t.Month(), t.Day(), t.Hour())
}

func (p *PrinterFile) ifSplit(){
	t := time.Now()
	if t.Hour() == p.lastSplitTime {
		return
	}
	file, err := os.OpenFile(p.filename(t), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		// TODO
		return
	}
	old := p.file
	p.file = file
	old.Close()
}