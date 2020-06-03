

2客 1服 两个相同的提案
//func TestTwoProposersSameValue(t *testing.T) {
//	n := NewNetwork(2, 1, 2, []int{2, 2})
//	go n.acceptors[0].Run()
//	go n.proposers[0].Run()
//	go n.proposers[1].Run()
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}

2客 1服 两个不相同的提案
//func TestTwoProposersDifferentValue(t *testing.T) {
//	n := NewNetwork(2, 1, 2, []int{1, 2})
//	go n.acceptors[0].Run()
//	go n.proposers[0].Run()
//	go n.proposers[1].Run()
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}
//

好多客 一个服 多个提案
//func TestManyProposersDifferentValues(t *testing.T) {
//	n := NewNetwork(5, 1, 2, []int{1, 2, 3, 4, 5})
//	go n.acceptors[0].Run()
//	for _, p := range n.proposers {
//		go p.Run()
//	}
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}
//

两个服务端
//func TestTwoAcceptors(t *testing.T) {
//	n := NewNetwork(1, 2, 2, []int{3})
//	go n.acceptors[0].Run()
//	go n.acceptors[1].Run()
//	go n.proposers[0].Run()
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}


多客 多服 相同提案
//func TestManyProposersManyAcceptorsSameValue(t *testing.T) {
//	n := NewNetwork(5, 5, 2, []int{1, 1, 1, 1, 1})
//	for _, a := range n.acceptors {
//		go a.Run()
//	}
//	for _, p := range n.proposers {
//		go p.Run()
//	}
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}
//

多客 多服 多提案
//func TestManyProposersManyAcceptorsDifferentValues(t *testing.T) {
//	n := NewNetwork(5, 5, 2, []int{1, 2, 3, 4, 5})
//	for _, a := range n.acceptors {
//		go a.Run()
//	}
//	for _, p := range n.proposers {
//		go p.Run()
//	}
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}
//

多客 多服 多相同/不相同提案
//func TestManyProposersManyAcceptorsSemiSameValues(t *testing.T) {
//	n := NewNetwork(5, 5, 2, []int{1, 2, 1, 2, 1})
//	for _, a := range n.acceptors {
//		go a.Run()
//	}
//	for _, p := range n.proposers {
//		go p.Run()
//	}
//	if n.learners[0].Run() != n.learners[1].Run() {
//		t.Errorf("Did not receive the same value!")
//	}
//}
