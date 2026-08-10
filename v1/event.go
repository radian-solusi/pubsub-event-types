package v1

import "time"

type ActivityEvent struct {
	Category    ActivityCategory    `json:"category"`
	Subcategory ActivitySubcategory `json:"subcategory"`
	UserID      string              `json:"user_id"`
	IPAddress   string              `json:"ip_address"`
	Message     string              `json:"message"`
	CreatedAt   time.Time           `json:"created_at"`
	Metadata    *string             `json:"metadata,omitempty"`
	Phone       *string             `json:"phone,omitempty"`
	Email       *string             `json:"email,omitempty"`
}

// MetaData is the structured payload NotificationEvent.Metadata carries.
type MetaData struct {
	ActionUrl *string `json:"action_url,omitempty"`
	Code      *string `json:"code,omitempty"`
}

type NotificationEvent struct {
	Category    NotificationCategory    `json:"category"`
	Subcategory NotificationSubcategory `json:"subcategory"`
	IPAddress   *string                 `json:"ip_address,omitempty"`
	Type        string                  `json:"type"`
	UserID      string                  `json:"user_id"`
	Message     string                  `json:"message"`
	CreatedAt   time.Time               `json:"created_at"`
	Metadata    *MetaData               `json:"metadata,omitempty"`
	Phone       *string                 `json:"phone,omitempty"`
	Email       *string                 `json:"email,omitempty"`
}
