git rev-parse HEAD > sha.txt
set /p sha=<sha.txt

echo sha = %sha%
go build -ldflags "-X main.version=1.0 -X main.sha=%sha%"
