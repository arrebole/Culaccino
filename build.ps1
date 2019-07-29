function buildFE {
    Start-Process -FilePath "yarn" -ArgumentList "run build"-WorkingDirectory "./client" -NoNewWindow;
}

buildFE;