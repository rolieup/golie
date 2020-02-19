# Golie - A Golang ROLIE Implementation
A client/server implementation of the Resource-Oriented Lightweight Information Exchange (ROLIE) specification [RFC8322](https://tools.ietf.org/html/rfc8322)

## Building

Change directory to the `build` directory:
```
cd build
```

To build the Golie client, run the following from the `build` directory:

```
go build -o golie ../cmd/golie/main.go
```

To build the Golie server run the following from the `build` directory:

```
go build -o golied ../cmd/golied/main.go
```

You can then run the Golie client `golie` and Golie server `golied` from the
`build` directory.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.