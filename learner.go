package paxos

import (
	"log"
	"net"
	"net/rpc"
	"sync"
)

type Learner struct {
	mu            sync.Mutex
	localAddr     string
	acceptorCount map[float32]int
	acceptedValue map[float32]interface{}
	decidedValue  interface{}
	quorumSize    int
	listener      net.Listener
}

func (le *Learner) getQuorumSize() int {
	le.mu.Lock()
	size := le.quorumSize
	le.mu.Unlock()
	return size
}

func (le *Learner) RecieveAccepted(arg *AcceptedMsg, reply *EmptyMsg) error {
	// PASS
	return nil
}

func (le *Learner) startRpc() {
	rpcx := rpc.NewServer()
	rpcx.Register(le)
	l, e := net.Listen("tcp", le.localAddr)
	le.listener = l
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

func (le *Learner) clean() {
	le.acceptedValue = nil
}

func (le *Learner) close() {
	le.listener.Close()
}
