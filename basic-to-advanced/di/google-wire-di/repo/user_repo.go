package repo

type User struct {
	ID   int
	Name string
}

type UserRepo interface {
	GetUser(id int) *User
}

type MockUserRepo struct{}

func NewMockUserRepo() UserRepo {
	return &MockUserRepo{}
}

func (r *MockUserRepo) GetUser(id int) *User {
	return &User{ID: id, Name: "Alice"}
}
