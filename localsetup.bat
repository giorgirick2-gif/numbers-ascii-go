@echo off
winget install Git.Git
winget install GoLang.Go
git clone https://github.com/giorgirick2-gif/numbers-ascii-go.git
cd numbers-ascii-go
echo STARTING LOCALHOST SERVER...
echo YOU WILL BE PROMPTED NETWORK ACCESS, PRESS "Allow"
echo THIS WILL BE HOSTED AT: localhost:8080
echo PRESS CONTROL + C THEN "Y" AND ENTER TO EXIT
go run main.go
exit