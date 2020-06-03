package paxos

import "testing"

//一个客户端,一个服务端,一个提案
func TestSingleProposer(t *testing.T) {
	n := NewNetwork(1, 1, 2, []int{1})
	go n.acceptors[0].Run()
	go n.proposers[0].Run()
	if n.learners[0].Run() != n.learners[1].Run() {
		t.Errorf("Did not receive the same value!")
	}
}


