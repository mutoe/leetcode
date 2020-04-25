package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// BFS Solution
// 12ms 6.3M

// leetcode submit region begin(Prohibit modification and deletion)

func getImportance(employees []*Employee, id int) (sum int) {
	employeeMap := make(map[int]Employee)
	for _, employee := range employees {
		employeeMap[employee.Id] = *employee
	}

	queue := []int{id}
	for len(queue) > 0 {
		employeeId := queue[0]
		sum += employeeMap[employeeId].Importance
		queue = append(queue[1:], employeeMap[employeeId].Subordinates...)
	}

	return
}

// leetcode submit region end(Prohibit modification and deletion)
