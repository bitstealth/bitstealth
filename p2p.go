package main

func initP2P() {

}

type Node struct {
	Trusted bool   `json:"trusted"`
	IP      string `json:"ip"`
	Port    string `json:"port"`
}

type P2PManager struct {
	nodes        []*Node
	trustedNodes []*Node
}

func (p *P2PManager) discoverNodes() {

}
