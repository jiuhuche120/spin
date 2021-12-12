package spin

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	s := New("  \033[36m[name] Testing\033[m")
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
	time.Sleep(time.Second * 5)
}
