package transpiler

type GraphBuilder struct {
}

type Node struct {
	NodeName string
	Edges    []*Node
}

func (gBuilder *GraphBuilder) topologicalSort() {

}

func (gBuilder *GraphBuilder) IsDag() bool {

	return true
}

func MkNode(nodeName string) {

}
