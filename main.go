package llog

import "sync"

type KeyVal struct {
	key interface{}
	val interface{}
}

type CTX struct {
	Fields    []*KeyVal
	fieldLock sync.Mutex
}

func (this *CTX) AddField(key, val interface{}) {
	this.fieldLock.Lock()
	this.Fields = append(this.Fields, &KeyVal{key: key, val: val})
	this.fieldLock.Unlock()
}
