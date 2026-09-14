.DEFAULT_GOAL := run # defaut target if not arg is supplied after the make keyword

.PHONY:fmt vet build # helps make to not confuse the targets below with other files in our directory.
# The lines below define the targets (commands) make will run.
fmt:
		go fmt ./...
# Having a word after a target definition tells make to execute a specified target before the current one.
vet: fmt
		go vet ./...
build: vet
		go build -o build/go_app
run: build
		./build/go_app
clean:
		echo "Deleting compiled binary!"
		rm -rf build/go_app
