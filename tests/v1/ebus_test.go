package v1_test

import (
	"encoding/json"
	"testing"
	"time"

	v1 "github.com/radian-solusi/pubsub-event-types/v1"
)

// Guards the wire contract: ebus-pipeline sends these strings and ebus-order
// switches on them, so a typo or a collapsed duplicate silently skips a cleanup.
func TestEbusEmisiActions_AreDistinct(t *testing.T) {
	actions := []string{
		v1.EbusEmisiActionOfferingStarted,
		v1.EbusEmisiActionAllotmentStarted,
		v1.EbusEmisiActionSeriesCancelled,
		v1.EbusEmisiActionPipelineCancelled,
	}
	want := []string{"offering_started", "allotment_started", "series_cancelled", "pipeline_cancelled"}
	for i, got := range actions {
		if got != want[i] {
			t.Fatalf("action %d = %q, want %q", i, got, want[i])
		}
	}
	seen := make(map[string]struct{}, len(actions))
	for _, a := range actions {
		if _, dup := seen[a]; dup {
			t.Fatalf("duplicate action %q", a)
		}
		seen[a] = struct{}{}
	}
}

func TestEbusEmisiProcessDetail_OmitsUnsetFields(t *testing.T) {
	d := v1.EbusEmisiProcessDetail{ProcessID: "p1", Action: v1.EbusEmisiActionOfferingStarted,
		RequestedAt: time.Date(2026, 11, 11, 0, 0, 0, 0, time.UTC)}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"process_id":"p1","action":"offering_started","requested_at":"2026-11-11T00:00:00Z"}`
	if string(b) != want {
		t.Fatalf("got %s want %s", b, want)
	}
}

func TestEbusEmisiProcessEvent_JSONKeys(t *testing.T) {
	e := v1.EbusEmisiProcessEvent{ProcessID: "p1", EmisiID: 7, Action: v1.EbusEmisiActionAllotmentStarted,
		CreatedAt: time.Date(2026, 11, 11, 0, 0, 0, 0, time.UTC)}
	b, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"process_id":"p1","emisi_id":7,"action":"allotment_started","created_at":"2026-11-11T00:00:00Z"}`
	if string(b) != want {
		t.Fatalf("got %s want %s", b, want)
	}
}

func TestEbusEmisiProcessDetail_RoundTripSetFields(t *testing.T) {
	ts := time.Date(2026, 11, 11, 1, 0, 0, 0, time.UTC)
	d := v1.EbusEmisiProcessDetail{ProcessID: "p1", Action: v1.EbusEmisiActionAllotmentStarted,
		RequestedAt: ts, ClaimedAt: &ts, FinishedAt: &ts, Result: map[string]int{"dropped": 2}, Error: "x"}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	var got v1.EbusEmisiProcessDetail
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.ClaimedAt == nil || got.FinishedAt == nil || got.Result["dropped"] != 2 || got.Error != "x" {
		t.Fatalf("round trip lost fields: %s", b)
	}
}
