package pkg

import (
	"fmt"
	"testing"
	"time"
)

func TestMillTimeStampToTime(t *testing.T) {
	ti := MillTimeStampToTime(time.Now().UnixNano() / int64(time.Millisecond))
	s := ti.Format("2006 01-02")
	fmt.Printf("%#v\n", s)
}
