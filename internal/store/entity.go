package store

import (
	"time"
)

// struct for list of file with their info
type ListFile struct {
	URL        string
	Uploader   string
	UploadedAt time.Time
}
