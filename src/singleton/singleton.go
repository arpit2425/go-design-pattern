package singleton

import "fmt"
type Singleton struct{}
var instance *Singleton
func GetInstance() *Singleton{
	if(instance==nil){
		instance=&Singleton{}
	}
	return instance
}
func Run(){
	instance1:=GetInstance()
	instance2:=GetInstance()
	if instance1 == instance2 {
		fmt.Println("Same reference")
	} else{
		fmt.Println("Same reference")
	}
}