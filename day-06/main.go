package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Node is the node of a tree
// with children inside
type Node struct {
	Name  string
	Edges []*Edge
	Count int
}

type Edge struct {
	From    *Node
	To      *Node
	Visited bool
}

func (n *Node) Traverse() {
	for _, child := range n.Edges {
		if !child.Visited && child.From.Name == n.Name {
			child.Visited = true
			child.To.Count = n.Count + 1
			child.To.Traverse()
		}
	}
}

func (n *Node) GetShortestPaths() {
	for _, child := range n.Edges {
		if child.From.Name == n.Name && !child.Visited {
			child.Visited = true
			child.To.Count = n.Count + 1
			child.To.GetShortestPaths()
		} else if child.To.Name == n.Name && !child.Visited {
			child.Visited = true
			child.From.Count = n.Count + 1
			child.From.GetShortestPaths()
		}
	}
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
	edges := []*Edge{}
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
		edge := Edge{
			From: node1,
			To:   node2,
		}
		edges = append(edges, &edge)
		node1.Edges = append(node1.Edges, &edge)
		node2.Edges = append(node2.Edges, &edge)
	}

	com := nodeMap["COM"]
	com.Traverse()
	result := 0
	for _, v := range nodeMap {
		result += v.Count
	}
	fmt.Println(result)
	// part 2
	// reset all counters
	for k := range nodeMap {
		nodeMap[k].Count = 0
	}
	for i := range edges {
		edges[i].Visited = false
	}

	you := nodeMap["YOU"]
	you.GetShortestPaths()

	san := nodeMap["SAN"]
	// we decrease by two since
	// we don't take into account
	// the two node orbits
	fmt.Println(san.Count - 2)

}
