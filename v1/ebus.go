package v1

import "time"

const (
	// Sent by ebus-pipeline on waiting_for_offering -> offering and on the new
	// pra_effective -> offering (IsBookBuilding == false); ebus-order consumes it
	// to drop book-building orders whose yield exceeds the offering yield.
	EbusEmisiActionOfferingStarted = "offering_started"
	// Sent by ebus-pipeline on offering -> allotment; ebus-order consumes it to run
	// the allotment cleansing (unconfirmed BB, unapproved, invalid SID/SRE).
	EbusEmisiActionAllotmentStarted = "allotment_started"
	// Sent by ebus-pipeline when UpdatePipeline cancels >= 1 series while the emisi
	// is in book_building/waiting_for_offering/offering; ebus-order drops the active
	// orders of the cancelled series. The cancelled series are derived by ebus-order
	// from series.status, so the event carries no series list.
	EbusEmisiActionSeriesCancelled = "series_cancelled"
	// Sent by ebus-pipeline when the emisi moves to cancelled; ebus-order drops the
	// active orders of every series of that emisi.
	EbusEmisiActionPipelineCancelled = "pipeline_cancelled"
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
