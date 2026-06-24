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

type NotificationEvent struct {
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	IPAddress   string    `json:"ip_address"`
	Type        string    `json:"type"`
	UserID      string    `json:"user_id"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	Metadata    *string   `json:"metadata,omitempty"`
	Phone       *string   `json:"phone,omitempty"`
	Email       *string   `json:"email,omitempty"`
}
