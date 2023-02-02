# matrix-go

## Starting the app
 *	Start app with `go run .`
 *	To run tests, use `go test`

## End points
* curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
* curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
* curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
* curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
* curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
* For files, you can use `file=@/path/matrix.csv`, `file=@/path/matrix-invalid.csv`, or `file=@/path/matrix-empty.csv`
