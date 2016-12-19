# go-eve

[![Go Report Card](https://goreportcard.com/badge/github.com/killmails/go-eve)](https://goreportcard.com/report/github.com/killmails/go-eve)

go-eve is a library for interacting with EVE Online's XML API.

```sh
$ cat test.go
```
```go
package main

import (
    "fmt"

    "github.com/killamils/go-eve/client"
    "github.com/killamils/go-eve/xmlapi"
)

func main() {
	// No Authentication
	api, err := xmlapi.New(&client.Options{})
	ss, err := api.Server().ServerStatus()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Status:")
	fmt.Println("-----")
	fmt.Printf("Server Open: %t\n", ss.ServerOpen)
	fmt.Printf("Online Players: %d\n", ss.OnlinePlayers)
	fmt.Println("-----")
	fmt.Printf("Cached For: %v\n", ss.GetCachedFor())

	// API Key Authentication
	key := client.ApiKey{
		KeyId:            123,
		VerificationCode: "abcd...",
	}
	opts := &client.Options{
		ApiKey: key,
	}
	api, err = xmlapi.New(opts)
	as, err := api.Account().AccountStatus()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nAccount Status:")
	fmt.Println("-----")
	fmt.Printf("Paid Until: %v\n", as.PaidUntil)
	fmt.Printf("Create Date: %v\n", as.CreateDate)
	fmt.Printf("Logon Count: %d\n", as.LogonCount)
	fmt.Printf("Logon Minutes: %d\n", as.LogonMinutes)
	fmt.Println("-----")
	fmt.Printf("Cached For: %v\n", ss.GetCachedFor())
}
```

```sh
$ go run test.go
Server Status:
-----
Server Open: true
Online Players: 23995
-----
Cached For: 51.203903759s

Account Status:
-----
Paid Until: 2016-11-21 20:05:00 +0000 UTC
Create Date: 2015-10-25 01:35:17 +0000 UTC
Logon Count: 705
Logon Minutes: 26762
-----
Cached For: 50.216473688s
```

## Start using it

1. Go:

    ```sh
    $ go get github.com/killmails/go-eve
    ```

2. Govendor:

    ```sh
    govendor fetch github.com/killmails/go-eve
    ```

## Contributing

- With issues:
  - Use the search tool before opening a new issue.
  - Review existing issues and provide feedback or react to them.
- With pull requests:
  - Open your pull request against develop
  - Your pull request should have no more than two commits, if not you should squash them.
