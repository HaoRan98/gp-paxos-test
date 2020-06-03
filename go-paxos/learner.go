package paxos

import "fmt"

type learner struct {
	id       int
	receives chan message	//接收通道
}

// 学习者初始化
func NewLearner(id int, receives chan message) *learner {
	l := new(learner)
	l.id = id
	l.receives = receives
	return l
}

// 监听msg.t状态,为chosen
func (l *learner) Run() int {
	v := -1
	for v == -1 {
		msg := <-l.receives
		switch msg.t {
		case Chosen:
			v = msg.pv
			msg.round++
		default:
		}
	}
	fmt.Printf("Learner %v: Chosen %v\n", l.id, v)
	return v
}
