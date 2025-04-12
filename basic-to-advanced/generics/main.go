package main

import "fmt"

// Điều cần nhớ	Giải thích
// Không abuse Generics	Go vẫn ưa đơn giản, rõ ràng
// Compile chặt chẽ hơn Java/C#	Không có reflection runtime nhiều
// Generics không thay thế interface	Dùng đúng mục đích khác nhau

func Min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Box[T any] struct {
	Value T
}

type Repository[T any] interface {
	Save(entity T) error
	FindByID(id int) (T, error)
}

type User struct {
	ID   int
	Name string
}

type UserRepo struct{}

func (r *UserRepo) Save(u User) error { return nil }
func (r *UserRepo) FindByID(id int) (User, error) {
	return User{ID: id, Name: "Alice"}, nil
}

type Number interface {
	int | int64 | float64
}

func Sum[T Number](arr []T) T {
	var total T
	for _, v := range arr {
		total += v
	}
	return total
}

func main() {
	fmt.Println(Min(3, 5))     // int
	fmt.Println(Min(3.2, 2.9)) // float64

	b1 := Box[int]{Value: 42}
	b2 := Box[string]{Value: "hello"}
	fmt.Println(b1.Value, b2.Value)

	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{1.5, 2.5, 3.0}))

}
