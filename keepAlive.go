package main

import (
	"fmt"
	"time"
)

type KeepAlive struct {
	tasks int
}

func (itSelf *KeepAlive) Wait() {
	for itSelf.tasks > 0 {
		fmt.Printf("KeepAlive because of %d task/s\n", itSelf.tasks)
		time.Sleep(time.Millisecond * 200)
	}
}

func (itSelf *KeepAlive) Add(amount int) {
	itSelf.tasks += amount
}

func (itSelf *KeepAlive) Done() {
	itSelf.tasks -= 1
}
