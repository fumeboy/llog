package llog

type Outputer interface {
	Write(data *LogData)
	Close()
}
