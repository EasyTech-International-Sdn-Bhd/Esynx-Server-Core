package options

type DatabaseStore int

const (
	MySQL DatabaseStore = iota
	Firestore
)
