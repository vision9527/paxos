package paxos

import (
	"fmt"
	"net/rpc"
	"strconv"
)

type PrepareMsg struct {
	ProposeID float32
}

type PromiseMsg struct {
	AcceptorAddr string
	ProposeID    float32
	Success      bool
	AccepedID    float32
	AccepedValue interface{}
}
type AcceptMsg struct {
	ProposeID    float32
	AcceptorAddr string
	Value        interface{}
}
type AcceptedMsg struct {
	ProposeID    float32
	AcceptorAddr string
	Success      bool
}

type EmptyMsg struct{}

func callRpc(peerAddr, roleService, method string, arg interface{}, reply interface{}) error {
	c, err := rpc.Dial("tcp", peerAddr)
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Call(roleService+"."+method, arg, reply)
	if err != nil {
		return err
	}
	return nil
}

func generateNumber(me int, number float32) float32 {
	var strNum string
	if number == 0 {
		strNum = fmt.Sprintf("1.%d", me)
		n, err := strconv.ParseFloat(strNum, 32)
		if err != nil {
			panic("error parse num")
		}
		return float32(n)
	}
	i := int(number) + 1 // 暂时不考虑float的精度问题
	strNum = fmt.Sprintf("%d.%d", i, me)
	n, err := strconv.ParseFloat(strNum, 32)
	if err != nil {
		panic("error parse num")
	}
	return float32(n)
}

var logLevel = 1

func logPrint(format string, a ...interface{}) {
	if logLevel == 1 {
		fmt.Printf(format+"\n", a...)
	}

}
