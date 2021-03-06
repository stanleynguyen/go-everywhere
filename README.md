# go-everywhere

### Compiling apk

- [Download ndk-bundle](https://developer.android.com/ndk/downloads)
- `go get -u godoc.org/golang.org/x/mobile/cmd/gomobile`
- Run this command inside this project root

```bash
ANDROID_HOME=$(pwd) gomobile build -ldflags "-X main.serverURL=<your server url>" -o light.apk github.com/stanleynguyen/go-everywhere/mobo
```

### Flash to arduino

```bash
tinygo flash -target arduino uno/main.go
```

### Compiling binary for RPi

```bash
GOOS=linux GOARCH=arm GOARM=5 go build -o pi.out -ldflags "-X main.serverURL=<your server url> -X main.pinNumber=<output pin number>" pi/main.go
```

Then copy to pi with `scp pi.out pi@raspberrypi.local:~/`
