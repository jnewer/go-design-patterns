package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type CPU struct {
	Name string
}

type ROM struct {
	Name string
}

type Disk struct {
	Name string
}

type Computer struct {
	Cpu  CPU
	Rom  ROM
	Disk Disk
}

func (s *Computer) Clone() *Computer {
	resume := *s
	return &resume
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
func (s *Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, s); err != nil {
		panic(err.Error())
	}
	return pc
}

func main() {
	cpu := CPU{"Intel"}
	rom := ROM{"金士顿"}
	disk := Disk{"三星"}

	c := Computer{
		Cpu:  cpu,
		Rom:  rom,
		Disk: disk,
	}
	c1 := c.BackUp()
	fmt.Printf("c1: %v\n", *c1)
}
