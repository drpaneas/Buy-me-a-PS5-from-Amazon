# How to use it

## 1. Put your data

Go to `main.go` and modify 3 *const* values, with your data:

```go
// Changes these const values with yours:
lastname           = "Wayne"                // Change this with your lastname
yourAmazonEmail    = "brucewayne@gmail.com" // Change this with your Amazon E-mail
yourAmazonPassword = "doYouBl33d?"          // Change this with your Amazon Password
```

Yes, I know, hardcoded password is lame.
You can probably send a PR to use an ENV variable -- I'm too lazy to make any modification now and I am not actively maintinaing this, so do not open issues.

## 2. Run it

Compile it the same way you would with any Go program.

```shell
go run main.go  # to run it
go build.go     # or, if you prefer a binary
```

This should work fine.
