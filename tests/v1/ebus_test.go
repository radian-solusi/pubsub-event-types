package v1_test

import (
	"encoding/json"
	"testing"
	"time"

	v1 "github.com/radian-solusi/pubsub-event-types/v1"
)

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
