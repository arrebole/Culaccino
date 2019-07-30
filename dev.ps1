
function goBuild {
    #go build -o "./bin/service.exe"  "./service/main.go"
    Start-Process -FilePath "go.exe" -ArgumentList "build -o ./bin/service.exe ./service/main.go" -NoNewWindow -Wait;
}

function startServer {
    $env:mode = "dev";
    Start-Process -FilePath "./bin/service.exe" -WorkingDirectory "./" 
    
}

function startFE {
    Start-Process -FilePath "yarn" -ArgumentList "run serve" -WorkingDirectory "./client";
}



goBuild; startServer; startFE;


