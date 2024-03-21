package adventofcode2018

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

// Node holds one node
type Node struct {
	name       string
	priority   int
	successors []*Node
}

// Edge holds one edge from Node to Node
type Edge struct {
	pred, succ *Node
}

// DAC holds a direct acyclic graph of edges and nodes.
type DAC struct {
	nodes map[string]*Node
	edges []Edge
	car   *Node
}

func NewDAC() *DAC {
	return &DAC{
		nodes: make(map[string]*Node),
	}
}

func (g *DAC) record(name string) *Node {
	node, ok := g.nodes[name]
	if ok {
		return node
	}
	node = &Node{
		name:     name,
		priority: int(name[0]),
	}
	g.nodes[name] = node
	return node
}

// Record inserts a new edge into graph.
func (g *DAC) Record(from, into string) {
	pred := g.record(from)
	succ := g.record(into)
	g.edges = append(g.edges, Edge{
		pred: pred,
		succ: succ,
	})

	if g.car == nil || g.car.name == into {
		g.car = pred
		log.Printf("car: %q\n", g.car.name)
	}

	pred.successors = append(pred.successors, succ)
	log.Printf("updated %+v\n", pred)
}

// Parse transforms any number of lines in the form 'Step C must be finished
// before step A can begin.' into a graph.
func Parse(r io.Reader) (*DAC, error) {
	g := NewDAC()
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) != 10 {
			return nil, fmt.Errorf("want 10 words but got %d", len(parts))
		}
		g.Record(parts[1], parts[7])
		lines++
	}
	if sc.Err() != nil {
		return nil, sc.Err()
	}
	return g, nil
}

// Flatten DAC into names.
func (g *DAC) Flatten() string {
	sb := strings.Builder{}
	var car *Node
	var prospects []*Node
	prospects = append(prospects, g.car)
	for len(prospects) > 0 {
		sort.SliceStable(prospects, func(i, j int) bool {
			return prospects[i].priority < prospects[j].priority
		})
		car, prospects = prospects[0], prospects[1:]
		sb.WriteString(car.name)
		for _, node := range car.successors {
			log.Printf("appending prospect %q\n", node)
			prospects = append(prospects, node)
		}
	}
	return sb.String()
}
