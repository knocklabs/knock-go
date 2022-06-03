# Knock Go Library

Knock API access for applications written in Go.

## Documentation

See the [documentation](https://docs.knock.app/) for Python usage examples.

## Installation

```sh
go get github.com/knocklabs/knock-go/knock
```

## Configuration

To use the library you must provide a secret API key, provided in the Knock dashboard. You must pass it into the client:

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

    // any example code from the rest of the readme can be run here

```

### Identifying Users

```go
    user, _ := client.Users.Identify(ctx, &knock.IdentifyUserRequest{
        User: &knock.User{
            ID:   "fun-user",
            Name: "John Hammond",
            CustomProperties: map[string]interface{}{
                "welcome":     "to jurassic park",
                "middle-name": "alfred",
            },
        },
    })
    fmt.Printf("user-name: %+v\n", user)
```

### Sending notifies (triggering workflows)

```go
    workflow, _ := client.Workflows.Trigger(ctx, &knock.TriggerWorkflowRequest{
        Workflow:   "test",
        Recipients: []string{"tim", "lex"},
        Data: map[string]interface{}{
            "life":      "found a way",
            "dinosaurs": "loose",
        },
    })
    fmt.Printf("workflow: %+v\n", workflow)
```

### Retrieving users

```go
    user, _ = client.Users.Get(ctx, &knock.GetUserRequest{
        ID: "fun-user",
    })

    fmt.Printf("retrieved-user: %+v\n", user)
```

### Deleting users

```go
    client.Users.Delete(ctx, &knock.DeleteUserRequest{
        ID: "fun-user",
    })
```

### Getting and setting channel data

```go
    setUserChannelDataResponse, _ := client.Users.SetChannelData(ctx, &knock.SetUserChannelDataRequest{
        UserID:    "test-123",
        ChannelID: "5d2377a0-92fb-4616-8315-eee843556566",
        Data: map[string]interface{}{
            "tokens": []string{"a", "b"},
        },
    })
    fmt.Printf("channel-data-set: %+v\n", setUserChannelDataResponse)

    getUserChannelDataResponse, _ := client.Users.GetChannelData(ctx, &knock.GetUserChannelDataRequest{
        UserID:    "test-123",
        ChannelID: "5d2377a0-92fb-4616-8315-eee843556566",
    })
    fmt.Printf("channel-data-found: %+v\n", getUserChannelDataResponse)
```

### Canceling workflows

```go
    client.Workflows.Cancel(ctx, &knock.CancelWorkflowRequest{
        Workflow:        "test",
        CancellationKey: "user-123",
    })
}
```
