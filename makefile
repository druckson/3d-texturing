test: main.go volumetricRenderer.go util.go vtk.go
	go build -o test

run: test
	./test

clean:
	rm test
