package main

import (
	"math"
)

type node struct {
	price    int
	children []*node
}

func minPathSumBruteforce(grid [][]int) int {
	nodeGrid := make([][]node, 0, len(grid))
	for i := 0; i < len(grid); i++ {
		nodeGrid = append(nodeGrid, make([]node, 0, len(grid[i])))
		for j := 0; j < len(grid[i]); j++ {
			nodeGrid[i] = append(nodeGrid[i], node{
				price:    grid[i][j],
				children: make([]*node, 0),
			})

			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				nodeGrid[i][j-1].children = append(nodeGrid[i][j-1].children, &nodeGrid[i][j])
				continue
			}

			if j == 0 {
				nodeGrid[i-1][j].children = append(nodeGrid[i-1][j].children, &nodeGrid[i][j])
				continue
			}

			nodeGrid[i][j-1].children = append(nodeGrid[i][j-1].children, &nodeGrid[i][j])
			nodeGrid[i-1][j].children = append(nodeGrid[i-1][j].children, &nodeGrid[i][j])
		}
	}

	price := minPathSumTraverse(nodeGrid[0][0])

	return price
}

func minPathSumTraverse(root node) int {
	if len(root.children) == 0 {
		return root.price
	}
	minPathPrice := math.MaxInt
	for _, child := range root.children {
		childPathPrice := minPathSumTraverse(*child)
		if childPathPrice < minPathPrice {
			minPathPrice = childPathPrice
		}
	}

	return root.price + minPathPrice
}
