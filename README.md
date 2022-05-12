# logwriter

Simple Go library for SystemD's journalctl like log formatting.

Provides time prefix generation for each line of provided messages.

## Usage example

```go
	import "github.com/nxshock/logwriter"

	// Create new writer that writes result to stdout
	writer := logwriter.New(os.Stdout)

	// Set custom time format and timezone if needed
	writer.TimeFormat = "02.01.06 15:04:05"
	writer.TimeZone = time.UTC

	writer.Print("hello world")
	// result:
	// 02.01.06 15:04:05 hello world

	writer.Write([]byte("line 1\nline 2\nline 3"))
	// result:
	// 02.01.06 15:04:05 line 1
	// 02.01.06 15:04:05 line 2
	// 02.01.06 15:04:05 line 3

	writer.Print("hello ")
	writer.Print("world")
	// result:
	// 02.01.06 15:04:05 hello world

	writer.Close()
	// writes final \n if not written before
```
