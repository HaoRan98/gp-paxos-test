package paxos

//网络结构
type network struct {
	proposers []*proposer
	acceptors []*acceptor
	learners  []*learner
}

// 初始化paxos网络,客户端数量,服务端数量,学习者数量,提案数量
func NewNetwork(nProposers, nAcceptors, nLearners int, vs []int) *network {
	cProposers := makeChannels(nProposers)
	cAcceptors := makeChannels(nAcceptors)
	cLearners := makeChannels(nLearners)

	n := new(network)
	n.proposers = make([]*proposer, nProposers)
	n.acceptors = make([]*acceptor, nAcceptors)
	n.learners = make([]*learner, nLearners)

	for i := range n.proposers {
		n.proposers[i] = NewProposer(i, vs[i], cProposers[i], cAcceptors, cLearners)
	}

	for i := range n.acceptors {
		n.acceptors[i] = NewAcceptor(i, cAcceptors[i], cProposers)
	}

	for i := range n.learners {
		n.learners[i] = NewLearner(i, cLearners[i])
	}

	return n
}

//建立角色的消息通道
func makeChannels(n int) []chan message {
	chans := make([]chan message, n)
	for i := range chans {
		chans[i] = make(chan message, 1024)
	}
	return chans
}

// 启动组件中的goruting
func (n *network) Start() {
	for _, l := range n.learners {
		go l.Run()
	}

	for _, a := range n.acceptors {
		go a.Run()
	}

	for _, p := range n.proposers {
		go p.Run()
	}
}
