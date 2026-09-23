package v1

import "time"

// DownloadEvent tracks the lifecycle of an async download/export job.
// Status is one of: pending, in_progress, completed, failed.
type DownloadEvent struct {
	DownloadID   string    `json:"download_id"`
	UserID       string    `json:"user_id"`
	Status       string    `json:"status"`
	Progress     int       `json:"progress"`
	PathDownload *string   `json:"path_download,omitempty"`
	ErrorMessage *string   `json:"error_message,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
