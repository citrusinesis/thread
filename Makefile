wire:
	cd ./pkg/domain && wire .

new:
	cd ./pkg && go run -mod=mod entgo.io/ent/cmd/ent new $(NAME)

generate:
	cd ./pkg && go generate ./ent

run:
	go run cmd/main.go

dev: wire run

build:
	go build -o runThread cmd/main.go