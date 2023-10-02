package concepts

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	fmt.Println("tests")
	testm02()
}
func testm01() {
	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is 10")
	} else {
		fmt.Println("x is less than 10")
	}
}
func testm02(){
	p:=5/2
	fmt.Println(p)
	


}
