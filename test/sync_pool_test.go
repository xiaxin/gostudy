package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	bufPool = sync.Pool{
		New: func() interface{} {
			// The Pool's New function should generally only return pointer
			// types, since a pointer can be put into the return interface
			// value without an allocation:
			println("new")
			return new(bytes.Buffer)
		},
	}
	wg sync.WaitGroup
)

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	defer wg.Done()

	wg.Add(1)
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteByte('\n')
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func TestPools(t *testing.T) {
	for i := 0; i < 100; i++ {
		go Log(os.Stdout, "i", fmt.Sprintf("%d", i))
	}

	wg.Wait()
}
