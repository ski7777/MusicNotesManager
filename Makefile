.DEFAULT_GOAL := mnm

.PHONY: mnm
mnm: internal/web/pkged.go
	go build cmd/mnm.go

.PHONY: generate
generate:
	go generate ./...

internal/web/pkged.go:
	echo "package dummy" > dummy.go
	pkger -o internal/web
	rm -f dummy.go

.PHONY: clean
clean:
	rm -f mnm
	rm -f internal/web/pkged.go