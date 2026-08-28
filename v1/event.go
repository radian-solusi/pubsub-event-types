package v1

import "time"

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
}

// MetaData is the flat, producer-supplied value bag NotificationEvent.Metadata
// carries — the named values a consumer binds into the message it renders.
//
// It is flat and already display-formatted by design: the producer decides how a
// number or date reads, the renderer only substitutes. Values may be PII
// (investor name, order code, old/new contact details) and MUST NOT be logged.
//
// Two keys are RESERVED and are not ordinary values — a consumer lifts them into
// the dedicated link slots of whatever it renders, so the set of inputs that can
// inject a clickable destination stays fixed and reviewable:
//
//	action_url       the primary call to action
//	unsubscribe_url  the footer opt-out
//
// Do not use either name for ordinary copy.
type MetaData map[string]string

// Reserved MetaData keys. See MetaData.
const (
	MetaKeyActionURL      = "action_url"
	MetaKeyUnsubscribeURL = "unsubscribe_url"
)

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
