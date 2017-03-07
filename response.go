package main

type Response struct {
	Success bool `json:"success"`
}

type ResponseP2PListNodes struct {
	Success bool    `json:"success"`
	Nodes   []*Node `json:"nodes"`
}
