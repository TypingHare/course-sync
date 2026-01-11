package doc

// Doc describes a documentation entry.
type Doc struct {
	// Name is the identifier of the documentation.
	Name string `json:"name"`

	// Title is the human-readable title of the documentation.
	Title string `json:"title"`

	// Path is the filesystem path to the documentation.
	Path string `json:"path"`

	// UpdatedAt is the time the documentation was last updated.
	UpdatedAt string `json:"updatedAt"`

	// IsDefault reports whether this documentation is the default.
	IsDefault bool `json:"isDefault"`
}
