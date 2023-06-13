package signal

import (
	"fmt"
	"testing"
)

func TestWait(t *testing.T) {
	fmt.Printf("\nReceived signal: %s\n", Wait())
}
