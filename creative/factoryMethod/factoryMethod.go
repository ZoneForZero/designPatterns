package factoryMethod

import "fmt"

/*
> - 简单工厂：唯一工厂类，一个产品抽象类，工厂类的创建方法依据入参判断并创建具体产品对象。
> - 工厂方法：多个工厂类，一个产品抽象类，利用多态创建不同的产品对象，避免了大量的if-else判断。
> - 抽象工厂：多个工厂类，多个产品抽象类，产品子类分组，同一个工厂实现类创建同组中的不同产品，减少了工厂子类的数量。
*/

// Operator 被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type SubOperator struct {
	*OperatorBase
}

func (p *SubOperator) Result() int {
	return p.a - p.b
}

type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperator) Result() int {
	return p.a + p.b
}

type SubOperatorFactory struct{}

func (p SubOperatorFactory) Create() Operator {
	return &SubOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperatorFactory  加法运算的工厂类
type PlusOperatorFactory struct{}

func (p PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

func init() {
	//加法
	plusFactory := PlusOperatorFactory{}
	plusOperator := plusFactory.Create()
	plusOperator.SetA(10)
	plusOperator.SetB(20)
	result := plusOperator.Result()
	fmt.Println("plusOperator=", result)
}
