package dsa

import
(
	"testing"
	"fmt"
)

func TestMap(t *testing.T){
	fmt.Print()
	xm:=&xmap{}
	xm.Put("one",1)
	xm.Put("two",2)
	val,exists:=xm.Get("one")
	if(exists){
		fmt.Println(val)
	}

}

type xmap struct{
	keys []interface{}
	Values []interface{}
}

func (this *xmap)Get(key interface{})(interface{},bool){
	for i, k := range this.keys {
		if(k==key){
			return this.Values[i],true
		}
	}
	return nil,false
}

func (this *xmap)Put(key interface{},value interface{}){
	for i, k := range this.keys {
		if(k==key){
			this.Values[i]=value
			return
		}
	}
	this.keys=append(this.keys, key)
	this.Values = append(this.Values, value)
}