package object

import (
	"fmt"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//实例化

// e := Employee{"0","Bob",20}
// e1 := Employee{Name:"Mike",Age:30}
// e2 := new(Employee) //返回的引用,相当与 e := &Employee{},区别 不需要使用 -> 访问成员变量
// e2.Id = "2"
// e2.Age = 22
// e2.Name = "Rose"

//定义行为
func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//为了防止内存拷贝
func (e *Employee) String1() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
