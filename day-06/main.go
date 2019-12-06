package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Node is the node of a tree
// with children inside
type Node struct {
	Name   string
	Orbits int
	Nodes  []*Node
	Parent *Node
}

func (n *Node) String() string {
	var res strings.Builder
	if n.Parent != nil {
		res.WriteString(n.Parent.Name + " ) ")
	}
	res.WriteString(fmt.Sprintf("%s (%d)", n.Name, n.Orbits))
	for _, no := range n.Nodes {
		res.WriteString(fmt.Sprintf(" %s", no.Name))
	}
	return res.String()
}

// Traverse traverse the stellar system calculating
// orbits
func (n *Node) Traverse() {
	fmt.Println(n)
	for _, no := range n.Nodes {
		no.Orbits = n.Orbits + 1
		no.Traverse()
	}
}

// Add adds a new Node into de tree
func (n *Node) Add(new *Node) {
	new.Parent = n
	n.Nodes = append(n.Nodes, new)
}

func readFile(filename string) ([][2]string, error) {
	var result [][2]string
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(string(raw), "\n")
	for _, e := range splitted {
		if e == "" {
			continue
		}
		nodes := strings.Split(e, ")")
		result = append(result, [2]string{nodes[0], nodes[1]})
	}
	return result, nil
}

func main() {
	nodes, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}
	nodeMap := map[string]*Node{}
	for _, n := range nodes {
		var node1, node2 *Node
		var ok bool
		first, second := n[0], n[1]
		if node1, ok = nodeMap[first]; !ok {
			node1 = &Node{
				Name: first,
			}
			nodeMap[first] = node1
		}
		if node2, ok = nodeMap[second]; !ok {
			node2 = &Node{
				Name: second,
			}
			nodeMap[second] = node2
		}
		node1.Add(node2)
	}
	n := nodeMap["COM"]
	n.Traverse()
	var result int
	for _, n := range nodeMap {
		result += n.Orbits
	}
	fmt.Println(result)
}
