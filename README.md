# `graceful`

[![GoDoc](https://godoc.org/github.com/recoilme/graceful?status.svg)](https://godoc.org/github.com/recoilme/graceful)

Golang has ignore for signals, but unignore don't. Graceful catch passed signals and call graceful function.

## Usage

There are only one function `Unignore()`. This will ignore all signals except passed

### Example


```go

quit := make(chan os.Signal, 1)
graceful.Unignore(quit, fallback, graceful.Terminate...)
//graceful.Unignore(quit, nil, syscall.SIGINT, syscall.SIGQUIT)

```

## Contact

Vadim Kulibaba [@recoilme](http://github.com/recoilme)

## License

`graceful` source code is available under the MIT [License](/LICENSE).