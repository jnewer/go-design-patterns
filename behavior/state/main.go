package main

import "fmt"

type State interface {
	DoAction(context *Context)
	ToString() string
}

type Context struct {
	state State
}

func NewContext() *Context {
	return &Context{state: nil}
}

func (c *Context) GetState() State {
	return c.state
}

func (c *Context) SetState(state State) {
	c.state = state
}

type StartState struct {
}

func NewStartState() *StartState {
	return &StartState{}
}

func (s *StartState) DoAction(context *Context) {
	fmt.Println("现在是开始状态")
	context.state = s
}

func (s *StartState) ToString() string {
	return "开始状态"

}

type StopState struct {
}

func NewStopState() *StopState {
	return &StopState{}
}

func (s *StopState) DoAction(context *Context) {
	fmt.Println("现在是停止状态")
	context.state = s
}

func (s *StopState) ToString() string {
	return "停止状态"
}

func main() {
	context := NewContext()
	startState := NewStartState()
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())
	fmt.Println("-------------")
	stopState := NewStopState()
	stopState.DoAction(context)
	fmt.Println(context.GetState().ToString())
}
