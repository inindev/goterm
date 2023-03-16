[ -f go.sum ] || go mod tidy
go build -o goterm
