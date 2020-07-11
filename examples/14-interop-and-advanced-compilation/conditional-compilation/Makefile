debug:
	go build -tags debug

release:
	mv main_release.go main_release.go.bak
	m4 main_release.go.bak > main_release.go
	go build main_release.go
	mv main_release.go.bak main_release.go
