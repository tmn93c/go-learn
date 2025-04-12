package main

import (
	"basic-to-advanced/di/google-wire-di/di"
	"fmt"
)

// go install github.com/google/wire/cmd/wire@latest
// wire

// ✅ Ưu điểm: Đỡ phải viết wiring tay, dễ mở rộng
// ❌ Nhược điểm: Phải học cách dùng wire, debug không trực tiếp

func main() {
	service := di.InitializeUserService()
	fmt.Println("User name:", service.GetUserName(1))
}
