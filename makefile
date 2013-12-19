test: main.go scene.go vtk.go
	go build -o test main.go scene.go util.go vtk.go

run: test
	./test

clean:
	rm test
