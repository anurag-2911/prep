package concepts

import
(
	"testing"
	"fmt"
)

func TestMap(t *testing.T){
	q04()
}

func q04(){
	_, exists := map[int]int{1: 10}[2]
    fmt.Println(exists)

	mx:=map[string]func(int) (int){
		"func1":func(x int) int {return x},
		"func2": func(x int) int {return x},
	}
	fmt.Println(mx["func1"](2))

	myMapx := map[string]int{"one": 1, "two": 2}
    valx, okx := myMapx["three"]
    fmt.Println(valx, okx)

	xm:=map[int][]string{}
	xm[1]= append(xm[1], "one")
	xm[1]= append(xm[1], "two")
	fmt.Println(xm[1])

	xmap:=map[interface{}]int{}
	xmap[1]=1
	xmap["1"]=2
	fmt.Println(xmap[1],xmap["1"])

	hmap:=map[string]int{"apple":1,"banana":2}
	delete(hmap,"orange")
	fmt.Println(len(hmap))


	mp01:=map[int]int{}
	mp01[1]++
	fmt.Println(mp01[1])


	mpsingh:=map[int]int{1:1,2:2}
	for k, v := range mpsingh {
		fmt.Println(k,v," ")
	}

	mmp:=map[int]int{1:1,2:2}
	mmp2:=mmp
	mmp2[1]=101
	fmt.Println(mmp[1])
	
	mmap:=map[string]*int{}
	val:=42
	mmap["hello"]=&val
	fmt.Println(*mmap["hello"])
	
	myMap := map[int]bool{1: true, 2: false}
    valu := myMap[3]
    fmt.Println(valu)
}
func q03(){
	mmap:=map[int]string{1:"one",2:"two"}
	delete(mmap,1)
	fmt.Println(len(mmap))
}
func Q02(){
	mymap:=make(map[string]int)
	mymap["hello"]=1
	fmt.Println(mymap["world"])
	if val,found:=mymap["hello"];found{
		fmt.Println(val)
	}else{
		fmt.Println("not found")
	}
}
func Q01(){
	mymap:=map[int]string{}
	fmt.Println(mymap[1])
}
