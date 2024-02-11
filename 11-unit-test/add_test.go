package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// func TestAdd(t *testing.T) {
// 	re := add(1, 3)

// 	if re != 3 {
// 		t.Errorf("expect %d, acture %d", 3, re)
// 	}
// }

// // 跳过耗时测试用例
// // go test --short
// func TestAdd2(t *testing.T) {
// 	fmt.Println("skip before")

// 	if testing.Short() {
// 		t.Skip("short 模式下跳过")
// 	}

// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}

// 	re := add(1, 3)

// 	if re != 3 {
// 		t.Errorf("2: expect %d, acture %d", 3, re)
// 	}
// }

// 性能测试
// go test -bench=".*"

func BenchmarkAdd(bb *testing.B) {
	bb.ResetTimer()

	for i := 0; i < bb.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}

	bb.StopTimer()
}

const numbers = 10000

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string

		for j := 0; j < numbers; j++ {
			str += strconv.Itoa(j)
		}
	}

	b.StopTimer()
}

func BenchmarkBuilderAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder

		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}

		_ = builder.String()
	}

	b.StopTimer()
}
