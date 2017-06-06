package calc

import (
	"fmt"
	"testing"
)

func Print() {
	fmt.Println("hello")
}

func Benchmark_Print(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Print()
	}

}
