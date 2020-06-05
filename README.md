# go-everywhere

### Compiling apk

- [Download ndk-bundle](https://developer.android.com/ndk/downloads)
- `go get -u godoc.org/golang.org/x/mobile/cmd/gomobile`
- Run this command inside this project root

```bash
ANDROID_HOME=$(pwd) gomobile build -ldflags "-X main.serverURL=<your server url>" -o light.apk github.com/stanleynguyen/go-everywhere/mobo
```
