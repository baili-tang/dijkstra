package dijkstra

//Vertex is a single node in the network, contains it's ID, best distance (to
// itself from the src) and the weight to go to each other connected node (Vertex)
type Vertex struct {
	//ID of the Vertex
	ID int
	//Best distance to the Vertex
	distance int64

	bestVerticies []int
	//A set of all weights to the nodes in the map
	arcs map[int]int64

	//通过id +终点key取出段Id
	sIds map[int]int
}

//NewVertex creates a new vertex
func NewVertex(ID int) *Vertex {
	return &Vertex{ID: ID, bestVerticies: []int{-1}, arcs: map[int]int64{}}
}

//AddVerticies adds the listed verticies to the graph, overwrites any existing
// Vertex with the same ID.
func (g *Graph) AddVerticies(verticies ...Vertex) {
	for _, v := range verticies {
		v.bestVerticies = []int{-1}
		if v.ID >= len(g.Verticies) {
			newV := make([]Vertex, v.ID+1-len(g.Verticies))
			g.Verticies = append(g.Verticies, newV...)
		}
		g.Verticies[v.ID] = v
	}
}

func (v *Vertex) containsBest(id int) bool {
	for _, bestVertex := range v.bestVerticies {
		if bestVertex == id {
			return true
		}
	}
	return false
}

func (v *Vertex) AddArc(Destination int, Distance int64) {
	if v.arcs == nil {
		v.arcs = map[int]int64{}
	}
	v.arcs[Destination] = Distance
}

func (v *Vertex) AddSId(Destination int, segmentId int) {
	if v.sIds == nil {
		v.sIds = map[int]int{}
	}
	v.sIds[Destination] = segmentId
}

//RemoveArc completely removes the arc to Destination (if it exists), this method will
// not error if Destination doesn't exist, only nop
func (v *Vertex) RemoveArc(Destination int) {
	delete(v.arcs, Destination)
}

//GetArc gets the specified arc to Destination, bool is false if no arc found
func (v *Vertex) GetArc(Destination int) (distance int64, ok bool) {
	if v.arcs == nil {
		return 0, false
	}
	//idk why but doesn't work on one line?
	distance, ok = v.arcs[Destination]
	return
}

func (v *Vertex) RemoveSId(Destination int) {
	delete(v.sIds, Destination)
}
func (v *Vertex) GetSId(Destination int) (sId int, ok bool) {
	if v.sIds == nil {
		return 0, false
	}
	//idk why but doesn't work on one line?
	sId, ok = v.sIds[Destination]
	return
}
