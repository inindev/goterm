# goterm
serial terminal written in go, supports speeds such as 1.5e6 (rockchip)

### build
```
go mod tidy
go build -o goterm
```

### run
```
./goterm /dev/ttyUSB0
./goterm /dev/cu.usbserial-10 1.5e6
```
