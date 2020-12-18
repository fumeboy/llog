package llog


var l = &logManager{}


type logManager struct {
	printers []*Printer
}

type PrinterI interface{
	Open() error
	Write(msg []byte)
	Close()
}


func Close() {
	for _,p := range l.printers{
		close(p.waitCh)
		p.handler.Close()
	}
	for _,p := range l.printers{
		<- p.overCh
	}
}

func Use(p PrinterI) (*Printer, error){
	if err := p.Open(); err != nil{
		return nil, err
	}
	p_ := &Printer{
		handler: p,
		waitCh: make(chan []byte, 100),
		overCh: make(chan bool),
	}
	l.printers = append(l.printers, p_)
	go p_.run()
	return p_, nil
}

type Printer struct {
	overCh chan bool
	waitCh chan []byte
	handler PrinterI
}

func (p *Printer) Write(msg []byte) {
	p.waitCh <- msg
}

func (p *Printer) run() {
	for msg := range p.waitCh {
		p.handler.Write(msg)
	}
	p.overCh <- true
}

