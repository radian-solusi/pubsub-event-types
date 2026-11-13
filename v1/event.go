package v1

import "time"

// Wire statuses of an ActivityEvent. RUNNING and TIMEOUT are not wire statuses;
// they exist only on the consumer's stored document.
const (
	ActivityStatusBegin     = "BEGIN"
	ActivityStatusCompleted = "COMPLETED"
	ActivityStatusFailed    = "FAILED"
)

// ActivityEvent records something a service did.
//
// A process is one BEGIN event plus one terminal event (COMPLETED or FAILED)
// sharing the same ActivityID. An event without an ActivityID is a one-off
// occurrence, treated as COMPLETED. The lifecycle fields are all omitempty, so
// the legacy JSON shape is unchanged.
//
// Message and Metadata MUST NOT contain PII.
type ActivityEvent struct {
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	UserID      string    `json:"user_id"`
	IPAddress   string    `json:"ip_address"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	Metadata    *string   `json:"metadata,omitempty"`
	Phone       *string   `json:"phone,omitempty"`
	Email       *string   `json:"email,omitempty"`

	EventID    string     `json:"event_id,omitempty"`    // UUIDv7, unique per published message (dedupe)
	ActivityID string     `json:"activity_id,omitempty"` // UUIDv7, shared by BEGIN and terminal event
	Status     string     `json:"status,omitempty"`      // ActivityStatus*
	Name       string     `json:"name,omitempty"`        // human-readable process name
	Source     string     `json:"source,omitempty"`      // web|console|subscriber|scheduler|system
	Service    string     `json:"service,omitempty"`
	Pod        string     `json:"pod,omitempty"`
	StartedAt  *time.Time `json:"started_at,omitempty"` // BEGIN time; repeated on the terminal event
	EndedAt    *time.Time `json:"ended_at,omitempty"`   // terminal event only
	Error      string     `json:"error,omitempty"`      // FAILED only
	TraceID    string     `json:"trace_id,omitempty"`   // OTel trace id, 32 hex chars
}

// MetaData is the flat, producer-supplied value bag NotificationEvent.Metadata
// carries — the named values a consumer binds into the message it renders.
//
// It is flat and already display-formatted by design: the producer decides how a
// number or date reads, the renderer only substitutes. Values may be PII
// (investor name, order code, old/new contact details) and MUST NOT be logged.
//
// This module defines the envelope, not the vocabulary: no key name is reserved,
// enumerated, or given a constant here. Which names a notification requires is
// the consuming service's template contract — for notification-otp, each
// template's manifest declares them and the dispatcher rejects an event that
// omits one. That includes the keys carrying links (action_url,
// unsubscribe_url); they are ordinary entries in this bag, and the rule that
// keeps other values out of a URL position lives in that service's CI, not here.
type MetaData map[string]string

type NotificationEvent struct {
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	IPAddress   *string   `json:"ip_address,omitempty"`
	Type        string    `json:"type"`
	UserID      string    `json:"user_id"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	// Locale is the BCP-47 short language tag for the message copy — "id" or
	// "en". It is a first-class rendering dimension, not "additional context",
	// so it sits alongside Type rather than inside Metadata. An absent or
	// unrecognized value falls back to the consumer's default.
	Locale string `json:"locale,omitempty"`
	// Metadata is the flat value bag for this notification. Wire-compatible with
	// the earlier structured form: {"metadata":{"action_url":"…"}} decodes
	// unchanged, with action_url now a reserved key rather than a struct field.
	Metadata MetaData `json:"metadata,omitempty"`
	Phone    *string  `json:"phone,omitempty"`
	Email    *string  `json:"email,omitempty"`
}
