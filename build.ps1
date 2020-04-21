$env:DB_NAME = "./culaccino.db"
Start-Process go.exe -ArgumentList "build -o culaccino.exe ./src/main.go" -NoNewWindow -Wait