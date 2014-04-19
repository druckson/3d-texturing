../bin/:
	mkdir ../bin/

../bin/viewer: ../bin/ main.go volumetricRenderer.go util.go vtk.go pnm.go scalarField.go gl.go
	go build -o ../bin/viewer


run: ../bin/viewer
	../bin/viewer

clean:
	rm ./bin/
