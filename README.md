# Knock Go Library

Knock API access for applications written in Go.

## Documentation

See the [documentation](https://docs.knock.app/) for Python usage examples.

## Installation

```sh
go get github.com/knocklabs/knock-go/knock
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/knocklabs/knock-go/knock"
)

func main() {
    token := os.Getenv("KNOCK_TOKEN")

    ctx := context.Background()

    // create a new Knock API client with the given access token
    client, _ := knock.NewClient(
        knock.WithAccessToken(token),
    )

    user, _ := client.Users.Identify(ctx, &knock.IdentifyUserRequest{
        User: &knock.User{
            ID:   "user-126",
            Name: "John Hammond",
            CustomProperties: map[string]interface{}{
                "welcome":     "to jurassic park",
                "middle-name": "alfred",
            },
        },
    })

    fmt.Printf("message: %+v\n", user)

    messages, _, _ := client.Messages.List(ctx, &knock.ListMessagesRequest{})

    for _, message := range messages {
        fmt.Printf("message: %+v\n", message)
    }
}

```
