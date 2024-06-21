package composite

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (Employee) Count() int {
	return 1
}

type Department struct {
	Name              string
	SubOriganizations []IOrganization
}

func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOriganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOriganizations = append(d.SubOriganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}

	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOriganizations: []IOrganization{&Employee{}}})
	}
	return root
}
