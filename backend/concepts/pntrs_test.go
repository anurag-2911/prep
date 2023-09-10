package concepts

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	fmt.Println("tests")
	t02()
}
func changestr(s string){
	s="world"
}
func t02(){
	var ax *int
	ax= new(int)
	*ax=401
	fmt.Println(*ax)
	
	fxc:=[3]int{1,2,3}
	bfx:=fxc[:]
	bfx[0]=908
	fmt.Println(fxc[0])

	cn:=make(chan int)
	bcn:=cn
	fmt.Println(cn==bcn)

	strx:="hello"
	changestr(strx)
	fmt.Println(strx)


	f:=101
	g:=&f
	fmt.Println(g)

	varr:=[3]int{1,2,3}
	xs:=&varr
	xs[0]=87
	fmt.Println(varr[0])

	var nl *int
	fmt.Println(nl==nil)

	za:=42
	zc:=&za
	za=999
	fmt.Println(*zc)

	cx:=make(map[string]int)
	bx:=cx
	bx["got"]=101
	fmt.Println(cx["go"])

	arr:=[]int{1,2,3}
	brr:=arr
	brr[0]=901
	fmt.Println(arr[0])


	a:=42
	b:=&a
	c:=&b
	fmt.Println(**c)
	fmt.Println(*c)
	fmt.Println(c)
}

func t01() {
	a := 42
	b := a
	b = 101
	fmt.Println(a)
	fmt.Println(b)

	s:="go"
	s1:=&s
	*s1="golang"
	fmt.Println(s)
	ar:=[]int{1,2,3}
	change(ar)
	fmt.Println(ar[0])
}
func change(arr []int){
	arr[0]=999
}
