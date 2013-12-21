test: main.go scene.go util.go vtk.go
	go build -o test

run: test
	./test

clean:
	rm test
