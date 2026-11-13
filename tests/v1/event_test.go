package v1_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	v1 "github.com/radian-solusi/pubsub-event-types/v1"
)

func TestActivityEvent_LegacyJSONUnchanged(t *testing.T) {
	e := v1.ActivityEvent{Category: "AUTH", Subcategory: "LOGIN", UserID: "u1", IPAddress: "1.1.1.1",
		Message: "m", CreatedAt: time.Date(2026, 11, 13, 0, 0, 0, 0, time.UTC)}
	b, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"category":"AUTH","subcategory":"LOGIN","user_id":"u1","ip_address":"1.1.1.1","message":"m","created_at":"2026-11-13T00:00:00Z"}`
	if string(b) != want {
		t.Fatalf("got %s want %s", b, want)
	}
}

func TestActivityEvent_LifecycleRoundTrip(t *testing.T) {
	ts := time.Date(2026, 11, 13, 1, 2, 3, 0, time.UTC)
	in := v1.ActivityEvent{
		Category: "SYNC", Subcategory: "KSEI", UserID: "u1", IPAddress: "1.1.1.1", Message: "m", CreatedAt: ts,
		EventID: "e1", ActivityID: "a1", Status: v1.ActivityStatusFailed, Name: "Sync KSEI SID",
		Source: "console", Service: "sid-ksei", Pod: "pod-1", StartedAt: &ts, EndedAt: &ts,
		Error: "boom", TraceID: "0af7651916cd43dd8448eb211c80319c",
	}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{`"activity_id"`, `"status":"FAILED"`, `"started_at"`, `"ended_at"`, `"error"`, `"trace_id"`} {
		if !strings.Contains(string(b), k) {
			t.Errorf("json %s missing %s", b, k)
		}
	}
	var out v1.ActivityEvent
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("round trip mismatch:\n in=%+v\nout=%+v", in, out)
	}
}

func TestActivityStatusConstants(t *testing.T) {
	for got, want := range map[string]string{
		v1.ActivityStatusBegin:     "BEGIN",
		v1.ActivityStatusCompleted: "COMPLETED",
		v1.ActivityStatusFailed:    "FAILED",
	} {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
