package command

import "fmt"

// Invoker 调用者
func NewInvoker() *Invoker {
	invoker := new(Invoker)
	return invoker
}

// 指令队列
type Invoker struct {
	commands []ICommand
}

func (i *Invoker) AddCommand(cmd ICommand) {
	i.commands = append(i.commands, cmd)
}
func (i *Invoker) Call() {
	if len(i.commands) == 0 {
		return
	}
	for _, command := range i.commands {
		command.Execute()
	}
}

// ICommand 命令接口
type ICommand interface {
	Execute()
}

// 指令1 ICommand的实现
type ShutdownCommand struct {
	tv *TV
}

func (s *ShutdownCommand) Execute() {
	s.tv.ShutDown()
}

// 指令2
type TurnOnCommand struct {
	tv *TV
}

func (t *TurnOnCommand) Execute() {
	t.tv.TurnOn()
}

type TV struct {
	Name string
}

func NewTV(name string) *TV {
	rt := TV{
		name,
	}
	return &rt
}
func (t *TV) ShutDown() {
	fmt.Printf("关闭%s电视\n", t.Name)
}
func (t *TV) TurnOn() {
	fmt.Printf("打开%s电视\n", t.Name)
}

// 命令模式，客户端通过调用者，传递不同的命令，然后不同的接受者对此进行处理
func main() {
	tv := NewTV("长虹")
	invoker := NewInvoker()
	shutdownCommand := &ShutdownCommand{tv: tv}
	turnOnCommand := &TurnOnCommand{tv: tv}
	invoker.AddCommand(shutdownCommand)
	invoker.AddCommand(turnOnCommand)
	invoker.AddCommand(shutdownCommand)
	invoker.Call()
}
