$env:PASSWORD = "dev"

function goBuild {
    #go build -o "./bin/service.exe"  "./service/main.go"
    Start-Process -FilePath "go.exe" -ArgumentList "build -o ./bin/service.exe ./service/main.go" -NoNewWindow -Wait;
}

function startServer {
    $env:mode = "debug";
    Start-Process -FilePath "./bin/service.exe" -WorkingDirectory "./" -NoNewWindow 
    
}

function startFE {
    Start-Process -FilePath "yarn" -ArgumentList "run serve" -WorkingDirectory "./web-app";
}

goBuild;  startFE; startServer;


