package integers

import (
	"testing"
)

func add(a,b int)int{
	return a+b+1
}
func TestAddr(t *testing.T){
	sum:=add(2,2)
	expect:=4
	if sum !=expect{
		t.Errorf("expect is '%d' but got '%d' ",expect,sum)
	}
}