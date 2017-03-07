package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func initServer() {
	server = &Server{ip: clip.IP, port: clip.Port}
}

type Server struct {
	ip   string
	port string
}

func (s *Server) listen() {
	http.HandleFunc("/", s.handleRequestDefault)
	http.HandleFunc("/p2p/list", s.handleRequestP2PListNodes)

	err := http.ListenAndServeTLS(s.ip+":"+s.port, "ssl/server.crt", "ssl/server.key", nil)
	log.Fatal(err)
}

func (s *Server) generateResponse(res interface{}, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func (s *Server) handleRequestDefault(w http.ResponseWriter, req *http.Request) {
	if req.RequestURI != "/" {
		http.NotFound(w, req)
		return
	}
	res := Response{Success: true}
	s.generateResponse(res, w, req)
}

func (s *Server) handleRequestP2PListNodes(w http.ResponseWriter, req *http.Request) {
	nd := &Node{IP: "127.0.0.1", Trusted: true}
	res := ResponseP2PListNodes{Success: true}
	res.Nodes = append(res.Nodes, nd)
	s.generateResponse(res, w, req)
}
