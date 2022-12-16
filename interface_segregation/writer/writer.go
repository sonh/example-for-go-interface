package writer

type Writer interface {
	Write()
}

func ApplyWriter(writer Writer) {
	// user writer to do something
	writer.Write()
}
