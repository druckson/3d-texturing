test: main.go volumetricRenderer.go util.go vtk.go pnm.go scalarField.go gl.go
	go build -o test

run: test
	./test

clean:
	rm test
