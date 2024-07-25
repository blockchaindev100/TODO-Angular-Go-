package models

type Roles struct {
	Id        string
	Role_type string
}

type Todo struct {
	Id        string
	Task_name string
	Status    string
	Active    bool
	User_id   int
}

type Users struct {
	Id                string
	Name              string
	Email             string
	Roleid            int
	Verification_code string
	Is_verified       bool
	Active            bool
}
