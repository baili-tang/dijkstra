package dijkstra

import (
	"errors"
	"fmt"
)

//GetMapped gets the key assosciated with the mapped int
func (g *Graph) GetMapped(a int) (string, error) {
	if !g.usingMap || g.mapping == nil {
		return "", errors.New("Map is not being used/initialised")
	}
	for k, v := range g.mapping {
		if v == a {
			return k, nil
		}
	}
	return "", errors.New(fmt.Sprint(a, " not found in mapping"))
}

//GetMapping gets the index associated with the specified key
func (g *Graph) GetMapping(a string) (int, error) {
	if !g.usingMap || g.mapping == nil {
		return -1, errors.New("Map is not being used/initialised")
	}
	if b, ok := g.mapping[a]; ok {
		return b, nil
	}
	return -1, errors.New(fmt.Sprint(a, " not found in mapping"))
}

//AddMappedVertex adds a new Vertex with a mapped ID (or returns the index if
// ID already exists).
func (g *Graph) AddMappedVertex(ID string) int {
	if !g.usingMap || g.mapping == nil {
		g.usingMap = true
		g.mapping = map[string]int{}
		g.highestMapIndex = 0
	}
	if i, ok := g.mapping[ID]; ok {
		return i
	}
	i := g.highestMapIndex
	g.highestMapIndex++
	g.mapping[ID] = i
	return g.AddVertex(i).ID
}

//AddMappedArc adds a new Arc from Source to Destination, for when verticies are
// referenced by strings.
func (g *Graph) AddMappedArc(Source, Destination string, Distance int64) error {
	return g.AddArc(g.AddMappedVertex(Source), g.AddMappedVertex(Destination), Distance)
}

//AddArc is the default method for adding an arc from a Source Vertex to a
// Destination Vertex
func (g *Graph) AddArc(Source, Destination int, Distance int64) error {
	if len(g.Verticies) <= Source || len(g.Verticies) <= Destination {
		return errors.New("Source/Destination not found")
	}
	g.Verticies[Source].AddArc(Destination, Distance)
	return nil
}

func (g *Graph) AddSId(Source, Destination int, SegmentId int) error {
	if len(g.Verticies) <= Source || len(g.Verticies) <= Destination {
		return errors.New("Source/Destination not found")
	}
	g.Verticies[Source].AddSId(Destination, SegmentId)
	return nil
}

// func (g *Graph) AddNode(Source, Destination, SegmentId int) error {
// 	if len(g.Verticies) <= Source || len(g.Verticies) <= Destination {
// 		return errors.New("Source/Destination not found")
// 	}

// 	once.Do(func() {
// 		g.Verticies[Source].Node = &Node{ID: Source, NextNode: &Node{NextNode: &Node{ID: Destination}, WithNextNodeSegment: SegmentId}, WithNextNodeSegment: -1}
// 	})

// 	g.Verticies[Source].Node = &Node{ID: Source, NextNode: &Node{NextNode: &Node{ID: Destination}, WithNextNodeSegment: SegmentId}, WithNextNodeSegment: SegmentId}
// 	return nil
// }

//RemoveArc removes and arc from the Source vertex to the Destination vertex
// fails if either vertex doesn't exist, but will succeed if destination is
// not an arc of Source (as a nop)
func (g *Graph) RemoveArc(Source, Destination int) error {
	if len(g.Verticies) <= Source || len(g.Verticies) <= Destination {
		return errors.New("Source/Destination not found")
	}
	g.Verticies[Source].RemoveArc(Destination)
	return nil
}

func (g *Graph) RemoveSId(Source, Destination int) error {
	if len(g.Verticies) <= Source || len(g.Verticies) <= Destination {
		return errors.New("Source/Destination not found")
	}
	g.Verticies[Source].RemoveSId(Destination)
	return nil
}
