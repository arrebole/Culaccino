function buildFE {
    Start-Process -FilePath "yarn" -ArgumentList "run build"-WorkingDirectory "./client" -NoNewWindow;
}

function buildServer {
    $env:GIN_MODE = "release"
    $env:GOOS = "linux"
    Start-Process -FilePath "go.exe" -ArgumentList "build -o ./bin/service.out ./service/main.go" -NoNewWindow -Wait
    $env:GOOS = "windows"
    Start-Process -FilePath "go.exe" -ArgumentList "build -o ./bin/service.exe ./service/main.go" -NoNewWindow -Wait
}
buildFE; buildServer;