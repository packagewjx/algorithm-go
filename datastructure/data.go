package datastructure

type Dummy struct {
	K int
}

func (d *Dummy) Key() int {
	return d.K
}

type Data interface {
	Key() int
}
