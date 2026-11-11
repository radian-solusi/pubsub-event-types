package v1

import "time"

const (
	EbusEmisiActionOfferingStarted  = "offering_started"
	EbusEmisiActionAllotmentStarted = "allotment_started"
)

// EbusEmisiProcessEvent: ebus-pipeline -> ebus-order, topic ebus-emisi-process-events.
type EbusEmisiProcessEvent struct {
	ProcessID string    `json:"process_id"` // UUIDv7; sama dengan EbusEmisiProcessDetail.ProcessID
	EmisiID   uint64    `json:"emisi_id"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
}

// EbusEmisiProcessDetail is the JSON stored in emisi.process_state_detail.
type EbusEmisiProcessDetail struct {
	ProcessID   string         `json:"process_id"`
	Action      string         `json:"action"`
	RequestedAt time.Time      `json:"requested_at"`
	ClaimedAt   *time.Time     `json:"claimed_at,omitempty"`
	FinishedAt  *time.Time     `json:"finished_at,omitempty"`
	Result      map[string]int `json:"result,omitempty"` // offering: dropped, pending_confirm; allotment: dropped
	Error       string         `json:"error,omitempty"`
}
