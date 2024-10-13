package templateMethod

import "fmt"

type Project interface {
	pre()
	reqinfo()
	end()
}

// 类似于一个抽象类
type req struct {
}

func (req) pre() {
	fmt.Println("请求前的行为..身份校验，参数转换")
}
func (req) end() {
	fmt.Println("数据封装，关闭请求")
}
func (req) reqinfo() {
	// 这里为空，不做任何事，fmt.Println("具体的业务逻辑，由继承类实现")
}

type somereq struct {
	req
}

func (*somereq) reqinfo() {
	fmt.Println("各自具体的业务内容，覆盖抽象类的同名方法")
}

// 一个具体子类的启动函数
func begin(r Project) {
	r.pre()
	r.reqinfo()
	r.end()
}

func main() {
	req1 := &somereq{}
	begin(req1)
	fmt.Println("============")
	req2 := &somereq{}
	begin(req2)
}
