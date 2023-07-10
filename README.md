# httpbin-http-wasm

`httpbin-http-wasm` is a httpbin server which allows you to load and use [http-wasm](https://http-wasm.io) compliant Wasm middlewares
to try them out.

## Getting started

Running a local middleware:

```bash
# terminal 1
go run github.com/jcchavezs/httpbin-http-wasm/cmd/httpbin-http-wasm@main --middleware my-middleware.wasm

# terminal 2
curl -i localhost:8080
```

Or using a remote one:

```bash
# terminal 1
go run github.com/jcchavezs/httpbin-http-wasm/cmd/httpbin-http-wasm@main --middleware https://github.com/http-wasm/http-wasm-host-go/raw/main/examples/auth.wasm

# terminal 2
curl -i localhost:8080
```
