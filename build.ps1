Write-Host "Building Go Agent for multiple platforms..." -ForegroundColor Green

# Create bin directory if it doesn't exist
New-Item -ItemType Directory -Force -Path "bin"



# Linux builds
Write-Host "Building for Linux (amd64)..."
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/agent-linux-amd64 ./main.go

Write-Host "Building for Linux (arm64)..."
$env:GOOS="linux";  $env:GOARCH="arm64"; go build -o bin/agent-linux-arm64 ./main.go

# Windows builds
Write-Host "Building for Windows (amd64)..."
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/agent-windows-amd64.exe ./main.go

Write-Host "Building for Windows (386)..."
$env:GOOS="windows"; $env:GOARCH="386"; go build -o bin/agent-windows-386.exe ./main.go

# macOS builds
Write-Host "Building for macOS (amd64)..."
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o bin/agent-darwin-amd64 ./main.go

Write-Host "Building for macOS (arm64)..."
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o bin/agent-darwin-arm64 ./main.go

Write-Host "All builds complete! Check the 'bin' directory." -ForegroundColor Green