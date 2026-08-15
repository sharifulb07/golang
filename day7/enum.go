package main

import (
	"fmt"
)

// time : 6 minutes

type ServerState int 

const (
	StateIdle ServerState=iota
	StateConnected
	StateError
	StateRetrying
)

var StateName=map[ServerState]string{
	StateIdle:"idle",
	StateConnected:"connected",
	StateError:"error",
	StateRetrying:"retrying",
}

func (ss ServerState) String()string {
	return StateName[ss] 
}


func transition(s ServerState)ServerState{
	switch s{
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return  StateIdle
	case StateError:
		return  StateError
	default:
		panic(fmt.Errorf("Error: %s", s))
	}
}

func main(){
	ns:=transition(StateIdle)
	fmt.Println(ns)

	ns1:=transition(ns)
	fmt.Println(ns1)
}