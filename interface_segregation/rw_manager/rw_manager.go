package rw_manager

import (
	"interace_homework/interface_segregation/reader"
	"interace_homework/interface_segregation/writer"
)

type RWManager interface {
	reader.Reader
	writer.Writer
}

func ApplyReadWriter(readWrite RWManager) {
	reader.ApplyReader(readWrite)
	writer.ApplyWriter(readWrite)
}
