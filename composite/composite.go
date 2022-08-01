package composite

type IHumanResource interface {
	CalculateSalary() float64
}

type Employee struct {
	id     int64
	salary float64
}

func (employee *Employee) CalculateSalary() float64 {
	return employee.salary
}

type Department struct {
	id       int64
	subNodes []IHumanResource
}

func (department *Department) CalculateSalary() float64 {
	var sum float64
	for _, node := range department.subNodes {
		sum += node.CalculateSalary()
	}
	return sum
}

func (department *Department) AddSubNode(hr IHumanResource) {
	department.subNodes = append(department.subNodes, hr)
}
