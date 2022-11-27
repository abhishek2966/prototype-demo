package resource

type employee struct {
	name   string
	email  string
	skills []string
}

func NewEmployee() employee {
	return employee{}
}
func (e employee) Name(name string) employee {
	e.name = name
	return e
}
func (e employee) Email(email string) employee {
	e.email = email
	return e
}
func (e employee) Skills(skillSet []string) employee {
	e.skills = skillSet
	return e
}
func (e employee) CloneEmployee() employee {
	e.skills = append([]string{}, e.skills...)
	return e
}
func (e employee) CloneEmployeeWithNewName(name string) employee {
	e.name = name
	e.skills = append([]string{}, e.skills...)
	return e
}
func (e employee) CloneEmployeeWithNewEmail(email string) employee {
	e.email = email
	e.skills = append([]string{}, e.skills...)
	return e
}
