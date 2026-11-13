# pubsub-event-types

Shared Go event types for pub/sub messaging.

## Installation

```bash
go get github.com/radian-solusi/pubsub-event-types/v1
```

## Usage

```go
package main

import (
	"time"

	v1 "github.com/radian-solusi/pubsub-event-types/v1"
)

func main() {
	event := v1.ActivityEvent{
		Category:    "user_activity",
		Subcategory: "login",
		UserID:      "12345",
		IPAddress:   "192.168.1.1",
		Message:     "User logged in",
		CreatedAt:   time.Now(),
	}

	_ = event

	// or notification event
	notification := v1.NotificationEvent{
		Category:    "system_notification",
		Subcategory: "maintenance",
		Message:     "System maintenance scheduled",
		CreatedAt:   time.Now(),
		UserID:      "xsaeaf-sdeaadff-eaf33",
		Type: 	  "email", // "email" or "sms"
		Phone:    "62852344431331" //user phone number to send when Type is "sms"
		Email:    "example@mail.com", //user email to send when Type is "email"
	}
	
}
```

### Activity lifecycle

One process = a `BEGIN` event + one terminal event (`COMPLETED`/`FAILED`) sharing `ActivityID`.
An event without `ActivityID` is a one-off. No PII in `Message`/`Metadata`.

```go
start := time.Now()
begin := v1.ActivityEvent{EventID: id1, ActivityID: aid, Status: v1.ActivityStatusBegin,
	Name: "Sync KSEI SID", Source: "console", Service: "sid-ksei", StartedAt: &start, CreatedAt: start}

end := time.Now()
done := begin
done.EventID, done.Status, done.EndedAt, done.CreatedAt = id2, v1.ActivityStatusFailed, &end, end
done.Error = "timeout" // FAILED only; StartedAt is repeated from BEGIN
```

## Alternate import alias

```go
import pubsub "github.com/radian-solusi/pubsub-event-types/v1"

activity := pubsub.ActivityEvent{}
notification := pubsub.NotificationEvent{}
```
