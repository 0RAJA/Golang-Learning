package main

func main() {

}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

//匿名函数很好用
func getImportance(employees []*Employee, id int) int {
	map1 := make(map[int]int)
	for index, i := range employees {
		map1[i.Id] = index
	}
	sum := 0
	var dfs func(int)
	dfs = func(index int) {
		sum += employees[map1[index]].Importance
		for _, index := range employees[map1[index]].Subordinates {
			dfs(index)
		}
	}
	dfs(id)
	return sum
}
