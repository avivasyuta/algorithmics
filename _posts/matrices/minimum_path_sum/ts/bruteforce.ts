type Node = {
	price: number
	children: Node[]
}

const minPathSumTraverse = (root: Node): number => {
    if (root.children.length == 0) {
        return root.price
    }

    let minPathPrice = Number.MAX_VALUE

    root.children.forEach(child => {
        const childPathPrice = minPathSumTraverse(child)

        if (childPathPrice < minPathPrice) {
            minPathPrice = childPathPrice
        }
    })

    return root.price + minPathPrice
}

export const minPathSumBruteforce = (grid: number[][]): number => {
	const nodeGrid: Node[][] = [];

	for (let i = 0; i < grid.length; i++) {
		nodeGrid[i] = [];

		for (let j = 0; j < grid[i].length; j++) {
			nodeGrid[i].push({
				price:    grid[i][j],
				children: [],
			})

			if (i == 0 && j == 0) {
				continue
			}

			if (i == 0) {
				nodeGrid[i][j-1].children.push(nodeGrid[i][j])
				continue
			}

			if (j == 0) {
				nodeGrid[i-1][j].children.push(nodeGrid[i][j])
				continue
			}

			nodeGrid[i][j-1].children.push(nodeGrid[i][j])
			nodeGrid[i-1][j].children.push(nodeGrid[i][j])
		}
	}

	return minPathSumTraverse(nodeGrid[0][0])
}


