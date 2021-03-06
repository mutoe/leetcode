package employee_importance

// https://leetcode.com/problems/employee-importance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// level: 1
// time: O(n) 12ms 100%
// space: O(n) 6.3M

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
