export NDK_TOOLCHAIN=~/my-ndk-toolchain/
export CC=$NDK_TOOLCHAIN/bin/arm-linux-androideabi-gcc
export GOROOT="/usr/local/go"
export GOPATH="/Users/mega/go"
export GOOS=android
export GOARCH=arm
export GOARM=7
export CGO_ENABLED=1

GO="$GOROOT/bin/go"

$GO build -x main.go
