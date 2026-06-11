package jfrogclient

import "fmt"

var agentName = "jfrog-client-go"
var agentVersion = "1.5.1"

func GetVersion() string {
	return agentVersion
}

func GetName()

func SetAgentName(name string) {
	agentName = name
}
