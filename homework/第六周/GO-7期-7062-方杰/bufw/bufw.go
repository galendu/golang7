package bufw

import (
	"io"
)

const (
	defaultBufSize = 4096 // 默认缓冲区大小
)

type BufferedFileWriter struct {
	err error     // 存储错误信息，调用底层io.Writer写入数据的时候可能的报错
	buf []byte    // 缓冲区，初始化的时候会默认设置为 defaultBufSize 的大小，初始化之后长度保持不变
	n   int       // 记录已经写入了多少
	wr  io.Writer // 底层写入器，io.Writer类型
}

// 返回一个 BufferedFileWriter的指针，可以自定义设置默认缓冲区的大小 size
func NewWriterSize(w io.Writer, size int) *BufferedFileWriter {
	b, ok := w.(*BufferedFileWriter)
	if ok && len(b.buf) >= size {
		return b
	}
	if size <= 0 {
		size = defaultBufSize
	}
	return &BufferedFileWriter{
		buf: make([]byte, size),
		wr:  w,
	}
}

// 返回一个 BufferedFileWriter 的指针，缓冲区大小为 defaultBufSize
func NewWriter(w io.Writer) *BufferedFileWriter {
	//return NewWriterSize(w, defaultBufSize)
	return &BufferedFileWriter{
		buf: make([]byte, defaultBufSize),
		wr:  w,
	}
}

// 返回当前缓冲区的长度，用于判断是否需要将数据交给底层写入器进行写入操作
func (b *BufferedFileWriter) Size() int { return len(b.buf) }

/*
func (b *BufferedFileWriter) Reset(w io.Writer) {
	b.err = nil
	b.n = 0
	b.wr = w
}
*/

// 实现io.Writer接口的Write方法，也就是底层写入器真正写入数据的方法，正常达到缓冲区大小或者主动调用Flush都会调用该方法实现数据从缓冲区写入磁盘
func (b *BufferedFileWriter) Write(p []byte) (nn int, err error) {
	for len(p) > b.Available() && b.err == nil {
		var n int
		if b.Buffered() == 0 {
			n, b.err = b.wr.Write(p)
		} else {
			n = copy(b.buf[b.n:], p)
			b.n += n
			b.Flush() // 调用Flush方法
		}
		nn += n
		p = p[n:]
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}

// 如果缓冲区数据没有达到要自动写入的长度，可以调用Flush方法将数据交给底层写入器写入数据
func (b *BufferedFileWriter) Flush() error {
	if b.err != nil {
		return b.err
	}
	if b.n == 0 {
		return nil
	}
	n, err := b.wr.Write(b.buf[0:b.n])
	if n < b.n && err == nil {
		err = io.ErrShortWrite // 写入操作写入的数据比提供的少
	}
	if err != nil {
		if n > 0 && n < b.n {
			copy(b.buf[0:b.n-n], b.buf[n:b.n])
		}
		b.n -= n
		b.err = err
		return err
	}
	b.n = 0
	return nil
}

// 返回当前缓冲区还剩多少数据没有写入
func (b *BufferedFileWriter) Available() int { return len(b.buf) - b.n }

// 返回当前缓存区已经写入了多少数据
func (b *BufferedFileWriter) Buffered() int { return b.n }

func (b *BufferedFileWriter) WriteByte(c byte) error {
	if b.err != nil {
		return b.err
	}
	if b.Available() <= 0 && b.Flush() != nil {
		return b.err
	}
	b.buf[b.n] = c
	b.n++
	return nil
}

/*
func (b *BufferedFileWriter) WriteRune(r rune) (size int, err error) {
	// Compare as uint32 to correctly handle negative runes.
	if uint32(r) < utf8.RuneSelf {
		err = b.WriteByte(byte(r))
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
	if b.err != nil {
		return 0, b.err
	}
	n := b.Available()
	if n < utf8.UTFMax {
		if b.Flush(); b.err != nil {
			return 0, b.err
		}
		n = b.Available()
		if n < utf8.UTFMax {
			// Can only happen if buffer is silly small.
			return b.WriteString(string(r))
		}
	}
	size = utf8.EncodeRune(b.buf[b.n:], r)
	b.n += size
	return size, nil
}
*/

func (b *BufferedFileWriter) WriteString(s string) (int, error) {
	nn := 0
	for len(s) > b.Available() && b.err == nil {
		n := copy(b.buf[b.n:], s)
		b.n += n
		nn += n
		s = s[n:]
		b.Flush()
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], s)
	b.n += n
	nn += n
	return nn, nil
}
