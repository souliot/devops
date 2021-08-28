@REM SET CGO_ENABLED=0
SET GOOS=linux
@REM SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w"