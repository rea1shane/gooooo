package strings

import (
	"fmt"
	"testing"
)

func TestC2M(t *testing.T) {
	fmt.Println(Pascal2Snake("NodeExporter"))
	fmt.Println(PascalWithSpace2Snake("Node Exporter"))
}

func TestS2C(t *testing.T) {
	fmt.Println(Snake2Pascal("node_exporter"))
	fmt.Println(Snake2PascalWithSpace("node_exporter"))
}
