SET GOARCH=wasm
SET GOOS=js
go build -o app.wasm
MOVE app.wasm ..\serve\html\api