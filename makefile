viewer: main.go volumetricRenderer.go util.go vtk.go pnm.go scalarField.go gl.go
	go build -o viewer

run: viewer
	./viewer

clean:
	rm viewer
