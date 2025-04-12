package main

import "fmt"

// Ưu điểm: Rõ ràng, dễ test. ❌ Nhược điểm: Dự án lớn sẽ phải "wire" nhiều thứ bằng tay.

type User struct {
	ID   int
	Name string
}

type UserRepo interface {
	GetUser(id int) *User
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetUserName(id int) string {
	user := s.repo.GetUser(id)
	if user == nil {
		return "Not found"
	}
	return user.Name
}

type MockUserRepo struct{}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{}
}

func (r *MockUserRepo) GetUser(id int) *User {
	return &User{ID: id, Name: "Alice"}
}

func main() {
	repo := NewMockUserRepo()
	service := NewUserService(repo)

	fmt.Println("User name:", service.GetUserName(1))
}
