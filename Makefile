create:
	mkdir -p cmd/$(name)
	printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\" \")\n}\n" > cmd/$(name)/main.go
