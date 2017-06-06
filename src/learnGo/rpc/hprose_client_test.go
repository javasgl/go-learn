package calc

import (
	"testing"
)

func BenchmarkTalk(benchmark *testing.B) {
	benchmark.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			talk()
		}
	})

}
