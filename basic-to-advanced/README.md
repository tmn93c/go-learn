1.Go Concurrency nâng cao
Channel patterns: Fan-in, fan-out, worker pool
Select statement nâng cao
Context (context.Context): để hủy goroutine đúng cách
Mutex, sync.Once, sync.WaitGroup
Race conditions và go run -race

2.Kiến trúc & thiết kế nâng cao
Dependency Injection (không có sẵn, dùng pattern hoặc thư viện như wire)
Clean Architecture / Hexagonal Architecture trong Go
Modular hóa project, tách package, đặt tên hợp lý
Testing nâng cao: table-driven test, mock, benchmark (testing.B, go test -bench)
Generics (Go 1.18+): dùng trong hàm, struct, interface

3. Hiệu năng và tối ưu
Benchmark code (go test -bench .)
Dùng pprof để phân tích CPU, memory
Garbage Collection và cách tránh GC bottlenecks
Dấu hiệu bạn đang bị bottleneck GC
Tăng đột biến CPU → do GC hoạt động liên tục
Latency bất thường, ngắt quãng (pause vài ms)
pprof memory profile thấy nhiều allocs, heap

Checklist “chống GC đau tim”
 Pre-allocate slices/map
 Reuse object qua sync.Pool
 Tránh tạo string tạm quá nhiều
 Dùng []byte thay vì string nếu cần hiệu năng cao
 Profile bằng pprof heap
 Benchmark với testing.B và b.ReportAllocs()

Zero-copy, memory pooling

4. Xây dựng dịch vụ thực tế
Viết REST API với net/http, gorilla/mux, hoặc gin
Middleware, logging, panic recovery
Auth (JWT, OAuth2)
GORM hoặc SQL trực tiếp (tối ưu truy vấn, migration)
Docker + build cross-platform

5.Giao tiếp giữa dịch vụ
gRPC (Protocol Buffers)
Kafka / RabbitMQ (Pub/Sub)
OpenTelemetry để trace logs

6.Viết CLI, tool hoặc lib dùng lại
cobra để viết CLI như kubectl
Viết custom linter với go/ast, go/parser
Viết plugin, hoặc thư viện opensource