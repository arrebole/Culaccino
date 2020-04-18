$env:DB_NAME = "./culaccino.db"
Start-Process go.exe -ArgumentList "build -o culaccion.exe ./src/main.go" -NoNewWindow -Wait