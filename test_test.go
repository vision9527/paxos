package paxos

import (
	"fmt"
	"testing"
	"time"
)

// 初始化集群
func makeCluster(proposerNum, acceptorNum, learnNum int) (p []*Proposer, a []*Acceptor, l []*Learner) {
	acceptorPort := 4100
	learnerPort := 4200

	learnPeers := make([]string, learnNum)
	for i := 0; i < learnNum; i++ {
		learnerPort++
		learner := &Learner{
			localAddr: formatAddr(learnerPort),
		}
		l = append(l, learner)
		learnPeers[i] = learner.localAddr
		l[i].startRpc()
	}

	acceptorPeers := make([]string, acceptorNum)
	for i := 0; i < acceptorNum; i++ {
		acceptorPort++
		acceptor := &Acceptor{
			localAddr:    formatAddr(acceptorPort),
			learnerPeers: learnPeers,
		}
		a = append(a, acceptor)
		acceptorPeers[i] = acceptor.localAddr
		a[i].startRpc()
	}

	for i := 0; i < proposerNum; i++ {
		proposer := &Proposer{
			acceptorPeers: acceptorPeers,
			me:            i + 1,
		}
		p = append(p, proposer)
	}
	return
}

func formatAddr(port int) string {
	return fmt.Sprintf("127.0.0.1:%d", port)
}

// 清除数据
func clean(p []*Proposer, a []*Acceptor, l []*Learner) {
	for _, i := range p {
		i.clean()
	}
	for _, i := range a {
		i.clean()
	}
	for _, i := range l {
		i.clean()
	}
}

// 关闭端口
func close(p []*Proposer, a []*Acceptor, l []*Learner) {
	for _, i := range p {
		i.close()
	}
	for _, i := range a {
		i.close()
	}
	for _, i := range l {
		i.close()
	}
}

// 检测单个proposer
func checkOne(t *testing.T, p []*Proposer, value interface{}) {
	for _, i := range p {
		if i.decidedValue != value {
			t.Fatalf("wrong decided value, want value:%v decided value:%v", value, i.decidedValue)
		}
	}
}

// 检测多个proposer
func checkMany(t *testing.T, p []*Proposer) {
	var value interface{}
	for _, i := range p {
		if i.decidedValue == nil {
			t.Fatalf("wrong decided value, decided value: nil")
		}
		if value != nil && value != i.decidedValue {
			t.Fatalf("wrong decided value, previous decided value:%v current decided value:%v", value, i.decidedValue)
		}
		if value == nil {
			value = i.decidedValue
		}
	}
	if value == nil {
		t.Fatalf("wrong decided value, should have one")
	}
}

func TestBasicPaxos(t *testing.T) {
	pNum := 1
	aNum := 3
	lNum := 2
	logPrint("TestBasicPaxos proposer num:%d, acceptor num:%d, learner num:%d begin", pNum, aNum, lNum)
	p, a, l := makeCluster(pNum, aNum, lNum)
	defer close(p, a, l)
	clean(p, a, l)

	p[0].propose(100)
	time.Sleep(1 * time.Second)
	checkOne(t, p, 100)
	clean(p, a, l)
	logPrint("TestBasicPaxos proposer num:%d, acceptor num:%d, learner num:%d end", pNum, aNum, lNum)
}

func TestSingleProposer(t *testing.T) {
	pNum := 1
	aNum := 3
	lNum := 2
	logPrint("TestSingleProposer proposer num:%d, acceptor num:%d, learner num:%d begin", pNum, aNum, lNum)
	p, a, l := makeCluster(pNum, aNum, lNum)
	defer close(p, a, l)
	clean(p, a, l)

	p[0].propose(100)
	time.Sleep(1 * time.Second)
	checkOne(t, p, 100)
	clean(p, a, l)

	p[0].propose(200)

	time.Sleep(1 * time.Second)
	checkOne(t, p, 200)
	clean(p, a, l)

	p[0].propose(300)
	time.Sleep(1 * time.Second)
	checkOne(t, p, 300)
	clean(p, a, l)
	logPrint("TestSingleProposer proposer num:%d, acceptor num:%d, learner num:%d end", pNum, aNum, lNum)
}

func TestManyProposer(t *testing.T) {
	pNum := 3
	aNum := 3
	lNum := 2
	logPrint("TestManyProposer proposer num:%d, acceptor num:%d, learner num:%d begin", pNum, aNum, lNum)
	p, a, l := makeCluster(pNum, aNum, lNum)
	defer close(p, a, l)

	for i, proposer := range p {
		proposer.propose(100 + i)
	}
	time.Sleep(1 * time.Second)
	checkMany(t, p)
	clean(p, a, l)

	for i, proposer := range p {
		proposer.propose(200 + i)
	}
	time.Sleep(1 * time.Second)
	checkMany(t, p)
	clean(p, a, l)

	for i, proposer := range p {
		proposer.propose(300 + i)
	}
	time.Sleep(1 * time.Second)
	checkMany(t, p)
	clean(p, a, l)
	logPrint("TestManyProposer proposer num:%d, acceptor num:%d, learner num:%d end", pNum, aNum, lNum)
}

func TestManyProposerUnreliable(t *testing.T) {
	pNum := 3
	aNum := 3
	lNum := 2
	logPrint("TestManyProposerUnreliable proposer num:%d, acceptor num:%d, learner num:%d begin", pNum, aNum, lNum)
	p, a, l := makeCluster(pNum, aNum, lNum)
	for _, acceptor := range a {
		acceptor.isunreliable = true
	}
	defer close(p, a, l)

	// instance 1 -> log index 1
	for i, proposer := range p {
		proposer.propose(100 + i)
	}
	time.Sleep(1 * time.Second)
	checkMany(t, p)
	clean(p, a, l)

	// instance 2 -> log index 2
	for i, proposer := range p {
		proposer.propose(200 + i)
	}
	time.Sleep(1 * time.Second)
	checkMany(t, p)
	clean(p, a, l)

	// instance 3 -> log index 3
	for i, proposer := range p {
		proposer.propose(300 + i)
	}
	time.Sleep(1 * time.Second)
	checkMany(t, p)
	clean(p, a, l)
	logPrint("TestManyProposerUnreliable proposer num:%d, acceptor num:%d, learner num:%d end", pNum, aNum, lNum)
}
