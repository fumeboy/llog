package llog

import "sync"

var (
	defaultOutputer    = NewConsoleOutputer()
	lm                 *loggerMgr
)

type loggerMgr struct {
	outputers   []Outputer
	chanSize    int
	level       LogLevel
	logDataChan chan *LogData
	wg          sync.WaitGroup
	globalFields []*KeyVal

}

func InitLogger(level LogLevel, chanSize int, globalFields...*KeyVal) {
	if chanSize <= 0 {
		chanSize = DefaultLogChanSize
	}

	lm = &loggerMgr{
		chanSize:    chanSize,
		level:       level,
		logDataChan: make(chan *LogData, chanSize),
		globalFields:globalFields,
	}
	lm.wg.Add(1)
	go lm.run()
}

func SetLevel(level LogLevel) {
	lm.level = level
}

func (l *loggerMgr) run() {
	for data := range l.logDataChan {
		if len(l.outputers) == 0 {
			defaultOutputer.Write(data)
			continue
		}

		for _, outputer := range l.outputers {
			outputer.Write(data)
		}
	}

	l.wg.Done()
}

func AddOutputer(ouputer Outputer) {
	lm.outputers = append(lm.outputers, ouputer)
	return
}

func Stop() {
	close(lm.logDataChan)
	lm.wg.Wait()

	for _, outputer := range lm.outputers {
		outputer.Close()
	}
}

