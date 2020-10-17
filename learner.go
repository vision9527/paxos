package paxos

import (
	"log"
	"net"
	"net/rpc"
)

type Learner struct {
	localAddr   string
	decideValue interface{}
	listener    net.Listener
}

func (learner *Learner) RecieveAccepted(arg *AcceptMsg, reply *EmptyMsg) error {
	// logPrint("[learner :%s RecieveAccepted msg :%v]", learner.localAddr, arg)
	learner.decideValue = arg.Value
	return nil
}

func (learner *Learner) startRpc() {
	rpcx := rpc.NewServer()
	rpcx.Register(learner)
	l, e := net.Listen("tcp", learner.localAddr)
	learner.listener = l
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}
			go rpcx.ServeConn(conn)
		}
	}()
}

func (learner *Learner) clean() {
	learner.decideValue = nil
}

func (learner *Learner) close() {
	learner.listener.Close()
}
