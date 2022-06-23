package go_trellis_db

import "github.com/wyattis/z/zset/zstringset"

func EdgesToAdjMatrix(edges []Edge, undirected bool) (nodes []string, mat [][]bool, degrees []int) {
	nodeSet := zstringset.New()
	for _, e := range edges {
		nodeSet.Add(e.SourceRespondentId, e.TargetRespondentId)
	}
	nodes = nodeSet.Items()
	nodeMap := map[string]int{}
	mat = make([][]bool, len(nodes))
	for i, n := range nodes {
		nodeMap[n] = i
		mat[i] = make([]bool, len(nodes))
	}
	for _, e := range edges {
		mat[nodeMap[e.SourceRespondentId]][nodeMap[e.TargetRespondentId]] = true
		if undirected {
			mat[nodeMap[e.TargetRespondentId]][nodeMap[e.SourceRespondentId]] = true
		}
	}
	for i := range nodes {
		d := 0
		for _, e := range mat[i] {
			if e {
				d++
			}
		}
		degrees = append(degrees, d)
	}
	return
}
