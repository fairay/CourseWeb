package dbuilder

import "api/recipes/objects"

type AccountBuilder struct {
	Login          string 
	Role           string 
	Salt           string
	HashedPassword string
}
func newAccountBuilder() *AccountBuilder { 
	return new(AccountBuilder) 
}
func (this *AccountBuilder) Build() *objects.Account { 
	return &objects.Account{Login: this.Login, Role: this.Role, 
		Salt: this.Salt, HashedPassword: this.HashedPassword} 
}

type AccountMother struct{}
func (AccountMother) Obj0() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test1"
	b.Role = "admin"
	b.Salt = ""
	b.HashedPassword = ""
	return b.Build()
}

func (AccountMother) Obj1() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test2"
	b.Role = "user"
	b.Salt = ""
	b.HashedPassword = ""
	return b.Build()
}

func (AccountMother) ObjUdp() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test2"
	b.Role = "admin"
	b.Salt = ""
	b.HashedPassword = ""
	return b.Build()
}
