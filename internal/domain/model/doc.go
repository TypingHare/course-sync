package model

import "time"

// Doc describes a documentation entry.
type Doc struct {
	// Name is the identifier of the documentation.
	Name string `json:"name"`

	// Title is the human-readable title of the documentation.
	Title string `json:"title"`

	// Version is the version of the documentation.
	Version string `json:"version"`

	// ReleasedAt is the release date of the documentation.
	ReleasedAt time.Time `json:"released_at"`

	// Path is the filesystem path to the documentation.
	Path string `json:"path"`

	// UpdatedAt is the time the documentation was last updated.
	UpdatedAt string `json:"updated_at"`

	// IsDefault reports whether this documentation is the default.
	IsDefault bool `json:"is_default"`
}
