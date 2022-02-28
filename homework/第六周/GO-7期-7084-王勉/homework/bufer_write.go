package main

import "os"

type BufferedFileWriter struct {
	buffer [1024]byte
	endPos int
	fileHandler *os.File
}

func NewBufferFileWrite(fd *os.File) *BufferedFileWriter {
	return &BufferedFileWriter{
		fileHandler: fd,
	}
}

func (writer *BufferedFileWriter) Flush() {
	if writer.endPos >0 {
		writer.fileHandler.Write(writer.buffer[:writer.endPos])
		writer.endPos = 0
	}
}

func (writer *BufferedFileWriter) Write(content []byte) {
	if len(content)>=1024{
		writer.fileHandler.Write(content)
	} else {
		if writer.endPos+len(content)>=1024{
			writer.Flush()
			writer.Write(content)
		} else {
			copy(writer.buffer[writer.endPos:],content)
			writer.endPos += len(content)
		}
	}
}

func (writer *BufferedFileWriter) writeString(content string) {
	writer.Write([]byte(content))
}

func testBufferWriter() {

}
