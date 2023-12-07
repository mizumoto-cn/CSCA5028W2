# GOPER

This is for CU Boulder CSCA5028 week3 assignment.

## How to run

### Setup

Build mysql container

```bash
docker compose up -d
```

### Quick Run

In case you have installed go(>=1.20.5), you may run the following command to run the program.

```bash
go build -o goper.exe
.\goper.exe http://trojan.mizumoto.tech:8081/data
```

If you didn't install go, you can visit [here](https://golang.org/doc/install) to install it.

Or you may run `.\goper.exe http://trojan.mizumoto.tech:8081/data` directly if you are using Windows.

## License

All homeworks are licensed under the [MIT](https://opensource.org/licenses/MIT) license. See ../LICENSE for details.
