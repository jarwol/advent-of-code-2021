package day13

func Part2() {
	paper, folds := Initialize("day13/input.txt")
	for _, fold := range folds {
		paper = doFold(paper, fold)
	}
	
	printPaper(paper)
}
