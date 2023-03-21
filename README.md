# goterm
serial terminal written in go, supports speeds such as 1.5e6 (rockchip)

### build
```
go mod tidy
go build -o goterm
```
or
```
sh make.sh
```


### run
```
./goterm /dev/ttyUSB0                   <- default speed of 115200, n81
./goterm /dev/cu.usbserial-10 1.5e6     <- 1500000 baud, n81
./goterm /dev/cu.usbserial-10 1500000   <- 1500000 baud, n81 (for those who like typing zeros)
```

### exit
```
<ctrl> + ]
```
