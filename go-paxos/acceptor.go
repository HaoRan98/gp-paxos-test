package paxos

import "fmt"

type acceptor struct {
	id        int				//编号
	apn       int				//票号
	apv       int				//消息
	maxpn     int				//最大票号
	receives  chan message		//消息通道
	proposers []chan message
}

// 初始化
func NewAcceptor(id int, receives chan message,
	proposers []chan message) *acceptor {
	a := new(acceptor)
	a.id = id
	a.apn = 0
	a.apv = 0
	a.maxpn = 0
	a.receives = receives
	a.proposers = proposers
	return a
}

var p *proposer

// 接收端算法
func (a *acceptor) Run() {
	fmt.Printf("Acceptor %v: started\n", a.id)
	for {
		msg := <-a.receives
		switch msg.t {
		// 阶段 1: Prepare-Promise
		case Prepare:
			//挑选票号最大的提案
			if msg.pn > a.maxpn {
				a.maxpn = msg.pn
				msg.round++
			}else {
				fmt.Println("已存在票号大于当前提案的相同提案")
				msg .round++
				p.prepare()
			}
			a.proposers[msg.from] <- NewPromiseMessage(a.id, a.apn, a.apv)
			fmt.Printf("Acceptor %v: sending Promise    OK\n", a.id)
		// 阶段 2: Accept-Accepted
		case Accept:
			//如果票号大于等于服务端最大票号,则回复 success
			if msg.pn >= a.maxpn {
				a.maxpn = msg.pn
				a.apn = msg.pn
				a.apv = msg.pv
				msg.round++
			}
			a.proposers[msg.from] <- NewAcceptedMessage(a.id, a.maxpn)
			fmt.Printf("Acceptor %v: sending Accepted   success\n", a.id)
		default:
		}
	}

}
