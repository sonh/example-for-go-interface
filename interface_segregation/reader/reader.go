package reader

type Reader interface {
	Read()
}

func ApplyReader(reader Reader) {
	// user reader to do something
	reader.Read()
}
