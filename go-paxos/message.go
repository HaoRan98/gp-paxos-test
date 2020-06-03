package paxos

type messageType int

//状态值
const (
	Prepare messageType = iota
	Promise
	Accept
	Accepted
	Chosen
)

//消息结构
type message struct {
	t    messageType
	from int
	pn   int
	pv   int
	round int //轮次
}

// propose 第一阶段发出prepare请求
func NewPrepareMessage(from, pn int) message {
	return message{t: Prepare, from: from, pn: pn}
}

// accepter 消息通过后进行promise的回复
func NewPromiseMessage(from, apn, apv int) message {
	return message{t: Promise, from: from, pn: apn, pv: apv}
}

// proposer 第二阶段,Proposer收到多数Acceptors承诺的Promise后，
// 向Acceptors发出Propose请求
func NewAcceptMessage(from, pn, pv int) message {
	return message{t: Accept, from: from, pn: pn, pv: pv}
}

// accepter 针对收到的Propose请求进行Accept处理。
func NewAcceptedMessage(from, pn int) message {
	return message{t: Accepted, from: from, pn: pn}
}

// Learn阶段。Proposer在收到多数Acceptors的Accept之后，
// 标志着本次Accept成功，决议形成，将形成的决议发送给所有Learners。
func NewChosenMessage(from, pv int) message {
	return message{t: Chosen, from: from, pv: pv}
}
