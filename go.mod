module github.com/roguesoftware/tla-story

go 1.14

replace github.com/roguesoftware/tla-proto => ../proto

require (
	github.com/roguesoftware/tla-proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)
