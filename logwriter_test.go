package logwriter

import (
	"bytes"
	"testing"
)

func TestBasicPrint(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Print("text")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 6 {
		t.Errorf("expected 6, got %d", n)
	}
	if lw.newLine {
		t.Error("newLine must be false")
	}
}

func TestBasicPrintln(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Println("text")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 7 {
		t.Errorf("expected 7, got %d", n)
	}
	if !lw.newLine {
		t.Error("newLine must be true")
	}
}

func TestBasicPrintf(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Printf("%s %d", "string", 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 10 {
		t.Errorf("expected 10, got %d", n)
	}
	if lw.newLine {
		t.Error("newLine must be false")
	}
}

func TestWriteWithEndLine(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Write([]byte("text\n"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 7 {
		t.Errorf("expected 7, got %d", n)
	}
	if !lw.newLine {
		t.Error("newLine must be true")
	}
}

func TestWriteWithoutEndLine(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Write([]byte("text"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 6 {
		t.Errorf("expected 6, got %d", n)
	}
	if lw.newLine {
		t.Error("newLine must be false")
	}
}

func TestClose(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	_, err := lw.Write([]byte("text"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = lw.Close()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !lw.newLine {
		t.Error("newLine must be true")
	}

	if buf.String() != "- text\n" {
		t.Errorf(`expected "- text\n", got "%s"`, buf.String())
	}
}

func TestWriteMultipleLines(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Write([]byte("line1\nline2\nline3"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 23 {
		t.Errorf("expected 23, got %d", n)
	}
	if lw.newLine {
		t.Error("newLine must be false")
	}
}

func TestPrintMultipleLines(t *testing.T) {
	buf := new(bytes.Buffer)

	lw := New(buf)
	lw.TimeFormat = "-"

	n, err := lw.Print("line1\nline2\nline3")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != 23 {
		t.Errorf("expected 23, got %d", n)
	}
	if lw.newLine {
		t.Error("newLine must be false")
	}
	if buf.String() != "- line1\n- line2\n- line3" {
		t.Errorf("wrong output:\n%s", buf.String())
	}
}
