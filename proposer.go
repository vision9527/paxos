package paxos

type Proposer struct {
	proposerID    int64
	currentValue  interface{}
	acceptedID    int64
	acceptedValue interface{}
	quorumSize    int64
}

func (p *Proposer) Listen() {}
func (p *Proposer) Stop()   {}

func (p *Proposer) Propose() interface{} {
	// start propose
	// runTwoPhase
	return nil
}

func (p *Proposer) runTwoPhase() {
	for {
		// start prepare
		// --------- RecivePromise
		// start accept
		// --------- ReciveAccepted
	}
}

func (p *Proposer) RecivePromise(m messager) messager {
	return messager{}
}

func (p *Proposer) ReciveAccepted(m messager) messager {
	return messager{}
}
