package entity

type AttachmentType int

const (
	AttachmentTypePDF AttachmentType = iota + 1
	AttachmentTypeImage
	AttachmentTypeVideo
)

type Attachment struct {
	// Unique identifier for the attachment. Useful for when multiple attachments are associated with a single entity.
	ID   string `json:"id"`
	Path string `json:"path"`
	// FromInternal indicates the attachment was created within the system.
	// If true, ReturnUrl should be baseURL + Path; otherwise ReturnUrl is just Path.
	FromInternal bool           `json:"from_internal"`
	Type         AttachmentType `json:"type"`
	Ref          string         `json:"ref"`
	ReturnUrl    string         `json:"return_url"`
}

type Device struct {
	Name      string `json:"name"`
	OS        string `json:"os"`
	Browser   string `json:"browser"`
	UserAgent string `json:"user_agent"`
}
