package main
import (
	"fmt"
//	"encoding/json"
)
// //Studen 学生
// type Studen struct {
// 	ID int `json:"id"`
// 	Gender string
// 	Name string
// }
// //Class  班级
// type Class struct {
// 	Title string
// 	Studens []*Studen
// }

// func main() {
// 	c:=Class{
// 		Title:"110",
// 		Studens:make([]*Studen,0,200),
// 	}
// 	for i:=0;i<10;i++{
// 		stu := &Studen{
// 			Name: fmt.Sprintf("stu%02d",i),
// 			Gender: "男",
// 			ID : i,
// 		}
// 		c.Studens = append(c.Studens,stu)
// 	}
// 	fmt.Printf("%v\n",c)
// 	data,err :=json.Marshal(c)
// 	if err != nil {
// 		fmt.Println("系列化失败")
// 		return
// 	}
// 	fmt.Printf("%s\n",data)
// }
//Person s
type Person struct {
	name   string
	age    int8
	dreams []string
}
// SetDreams s
// func (p *Person) SetDreams(dreams []string) {
// 	p.dreams = dreams
// }
//SetDreams s
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	p1.dreams[1] = "不睡觉"
	fmt.Println(p1.dreams)  // ?
}
