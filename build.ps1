$AppName = "stu-net"
$AppVer = "v0.3.1"

Remove-Item dist\ -Force -Recurse

Write-Output "Start Building..."

Write-Output "Building for Linux amd64..."
$ENV:CGO_ENABLED = 0; $ENV:GOOS = "linux"; $ENV:GOARCH = "amd64"; go build --ldflags "-s -w" -o dist/$($AppName)_$($AppVer)_linux_amd64/$($AppName)
tar -zcvf dist/$($AppName)_$($AppVer)_linux_amd64.tar.gz dist/$($AppName)_$($AppVer)_linux_amd64
Write-Output "Building for darwin amd64..."
$ENV:CGO_ENABLED = 0; $ENV:GOOS = "darwin"; $ENV:GOARCH = "amd64"; go build --ldflags "-s -w" -o dist/$($AppName)_$($AppVer)_darwin_amd64/$($AppName)
tar -zcvf dist/$($AppName)_$($AppVer)_darwin_amd64.tar.gz dist/$($AppName)_$($AppVer)_darwin_amd64
Write-Output "Building for windows amd64..."
$ENV:CGO_ENABLED = 0; $ENV:GOOS = "windows"; $ENV:GOARCH = "amd64"; go build --ldflags "-s -w" -o dist/$($AppName)_$($AppVer)_windows_amd64/$($AppName).exe
Write-Output ".\stu-net.exe -i" >> "dist/$($AppName)_$($AppVer)_windows_amd64/start.bat" 
$compress = @{
    Path             = "dist/$($AppName)_$($AppVer)_windows_amd64"
    CompressionLevel = "Fastest"
    DestinationPath  = "dist/$($AppName)_$($AppVer)_windows_amd64.zip"
    Force            = $true
}
Compress-Archive @compress
Write-Output "Building for Linux arm (such as Raspberry Pi) ... "
$ENV:CGO_ENABLED = 0; $ENV:GOOS = "linux"; $ENV:GOARCH = "arm"; go build --ldflags "-s -w" -o dist/$($AppName)_$($AppVer)_linux_arm64/$($AppName)
tar -zcvf dist/$($AppName)_$($AppVer)_linux_arm64.tar.gz dist/$($AppName)_$($AppVer)_linux_arm64

Write-Output "Build and Compress Complete. Cleaning up..."
# force remove
Remove-Item "dist/$($AppName)_$($AppVer)_linux_amd64" -Force -Recurse
Write-Output "rm dist/$($AppName)_$($AppVer)_linux_amd64" 
Remove-Item "dist/$($AppName)_$($AppVer)_darwin_amd64" -Force -Recurse
Write-Output "rm dist/$($AppName)_$($AppVer)_darwin_amd64"
Remove-Item "dist/$($AppName)_$($AppVer)_windows_amd64" -Force -Recurse
Write-Output "rm dist/$($AppName)_$($AppVer)_windows_amd64"
Remove-Item "dist/$($AppName)_$($AppVer)_linux_arm64" -Force -Recurse 
Write-Output "rm dist/$($AppName)_$($AppVer)_linux_arm64"