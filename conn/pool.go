package conn

import (
	"fmt"
//	"google.golang.org/grpc"
)

// Pool defines a list of connected nodes

type NodeConnection struct {
	Name string
}

type Pool struct {
	Connections []NodeConnection
}

func (p *Pool) AddNode(node NodeConnection) {

	p.Connections = append(p.Connections, node)
	fmt.Println("Added node name:", node.Name)

//	fmt.Println(node.Name)
//	return Pool.connections
}

func (p *Pool) RemoveNode() {
	fmt.Println("Remove node")
}

func (p *Pool) GetAllNodes() {
	

	fmt.Println("List of connections:", p.Connections)
}
