package interation

import "testing"

const repeatCount = 5

func Test_Repeat(t *testing.T) {
	resp := repeat("a")
	expect := "aaaa"
	if resp != expect {
		t.Errorf("expect is '%s' but got '%s'", expect, resp)
	}
}

func repeat(char string) string {
	var resp string
	for i := 0; i < repeatCount; i++ {
		resp += char
	}
	return resp
}

/**
基准测试
testing.B 可使你访问隐性命名（cryptically named）b.N。
基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。
*/
func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		repeat("a")
	}
}

/**
GOROOT=/usr/local/opt/go/libexec #gosetup
GOPATH=/Users/lgc/code/goLand/go_2020 #gosetup
/usr/local/opt/go/libexec/bin/go test -c -o /private/var/folders/zd/ppqvd37n0h34byl15vgmb50r0000gn/T/___BenchmarkRepeat_in_structure_interation structure/interation #gosetup
/private/var/folders/zd/ppqvd37n0h34byl15vgmb50r0000gn/T/___BenchmarkRepeat_in_structure_interation -test.v -test.bench ^BenchmarkRepeat$ -test.run ^$ #gosetup
goos: darwin
goarch: amd64
pkg: structure/interation
BenchmarkRepeat-12    	 9203545	       136 ns/op
PASS

测试了 9203545 次 运行一次 136纳秒
*/
