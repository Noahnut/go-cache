package gocache

type tinyLFU struct {
}

func (t *tinyLFU) Get(key interface{}) (interface{}, bool) {
	return nil, false
}

func (t *tinyLFU) Set(interface{}, interface{}) {

}

func (t *tinyLFU) Delete(key interface{}) bool {
	return true
}

func (t *tinyLFU) Contains(key interface{}) bool {
	return true
}

func (t *tinyLFU) Resize(size int) {

}

func (t *tinyLFU) Clean() {

}
