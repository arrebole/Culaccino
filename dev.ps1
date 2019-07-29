
function goBuild {
    param ($out)
    Start-Process -FilePath "go.exe" -ArgumentList "build -o $out ./service/main.go" -NoNewWindow;
}

function startServer {
    param($exe)
    Start-Process -FilePath $exe 
    
}

function startFE {
    param ($OptionalParameters)
    Set-Location "./client";
    Start-Process -FilePath "yarn" -ArgumentList "run serve";
    Set-Location "../";
}


$service = "./bin/service.exe";
goBuild -out $service;
startServer -exe $service; startFE;


