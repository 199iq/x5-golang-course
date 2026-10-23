package main

import "fmt"

func StudentScores() map[string]int {
	scores := make(map[string]int)
	scores["Анна"] = 3
	scores["Борис"] = 0
	scores["Вера"] = 5

	scores["Анна"] += 2
	delete(scores, "Вера")

	return scores
}

func main() {
	scores := StudentScores()

	fmt.Printf("Баллы Анны: %d\n", scores["Анна"])

	score, ok := scores["Борис"]
	fmt.Printf("Результат Бориса: %d %t\n", score, ok)
	score, ok = scores["Глеб"]
	fmt.Printf("Результат Глеба: %d %t\n", score, ok)

	fmt.Println(len(scores))
}
