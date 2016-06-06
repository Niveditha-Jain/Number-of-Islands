package main

func CountIslands(grid [][]int) int {

	var count int = 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
				convertMerge(grid, i, j)
			}
		}
	}
	return count
}

func convertMerge(Map [][]int, i int, j int) {

	if i < 0 || j < 0 || i > len(Map)-1 || j > len(Map[0])-1 {
		return
	}
	if Map[i][j] != 1 {
		return
	}

	Map[i][j] = 0
	convertMerge(Map, i-1, j)
	convertMerge(Map, i+1, j)
	convertMerge(Map, i, j-1)
	convertMerge(Map, i, j+1)
}
