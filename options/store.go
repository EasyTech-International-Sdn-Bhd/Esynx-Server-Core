package options

type DatabaseStore int

const (
	SqlDb DatabaseStore = iota
	Firestore
	Mock
)
