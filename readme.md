step to run agent
1-> go mod vendor

2-> in powershell
    -- .\build.ps1

3-> exe and sh file create in bin folder
    based on your os run application
        --> windows-amd64
            ./bin/agent-windows-amd64.exe

        --> windows-arm64
            ./bin/agent-linux-arm64

---> for other os , referes build.ps1 file, you can fetch ready made windows .exe on
    https://drive.google.com/file/d/1XlwqC6-q8F27mj4Qv6NhgyjnPXHq3TvJ/view?usp=sharing

---> this will create agent.log file  this will use by backend to fetch log of the server on demand