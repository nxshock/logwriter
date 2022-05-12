package logwriter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"
)

type LogWriter struct {
	TimeFormat string
	TimeZone   *time.Location
	writer     io.Writer

	newLine bool
}

// New created new LogWriter.
func New(w io.Writer) *LogWriter {
	lw := &LogWriter{
		TimeFormat: defaultTimeFormat,
		TimeZone:   defaultTimeZone,
		writer:     w,
		newLine:    true}

	return lw
}

// Println formats using the default formats for its operands and writes to writer.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (lw *LogWriter) Println(a ...any) (n int, err error) {
	s := fmt.Sprintln(a...)

	n, err = lw.Write([]byte(s))
	if err != nil {
		return n, err
	}

	lw.newLine = true

	return n, nil
}

// Print formats using the default formats for its operands and writes to writer.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (lw *LogWriter) Print(a ...any) (n int, err error) {
	s := fmt.Sprint(a...)

	n, err = lw.Write([]byte(s))
	if err != nil {
		return n, err
	}

	if s[len(s)-1] == '\n' {
		lw.newLine = true
	} else {
		lw.newLine = false
	}

	return n, nil
}

// Printf formats according to a format specifier and writes to writer.
// It returns the number of bytes written and any write error encountered.
func (lw *LogWriter) Printf(format string, a ...any) (n int, err error) {
	s := fmt.Sprintf(format, a...)

	n, err = lw.Write([]byte(s))
	if err != nil {
		return n, err
	}

	if s[len(s)-1] == '\n' {
		lw.newLine = true
	} else {
		lw.newLine = false
	}

	return n, nil
}

// Write writes len(p) bytes from p to the writer.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
func (lw *LogWriter) Write(p []byte) (n int, err error) {
	r := bufio.NewReader(bytes.NewReader(p))

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF && len(line) > 0 {
			nn, err := io.WriteString(lw.writer, lw.prefix()+" "+line)
			n += nn
			if err != nil {
				return n, err
			}

			if line[len(line)-1] == '\n' {
				lw.newLine = true // TODO: uncovered or unused?
			} else {
				lw.newLine = false
			}
			break
		}
		if err == io.EOF {
			break
		}

		if lw.newLine {
			nn, err := io.WriteString(lw.writer, lw.prefix()+" "+line)
			n += nn
			if err != nil {
				return n, err
			}
		} else {
			nn, err := io.WriteString(lw.writer, line)
			n += nn
			if err != nil {
				return n, err
			}
			lw.newLine = true
		}
	}

	return n, nil
}

func (lw *LogWriter) Close() error {
	if lw.newLine {
		return nil
	}

	_, err := io.WriteString(lw, "\n")
	if err != nil {
		return err
	}

	lw.newLine = true

	return nil
}

func (lw *LogWriter) prefix() string {
	return fmt.Sprintf("%s", time.Now().Format(lw.TimeFormat))
}
