package entity

type AttachmentType int

const (
	AttachmentTypePDF AttachmentType = iota + 1
	AttachmentTypeImage
	AttachmentTypeVideo
)

type Attachment struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	// FromInternal indicates the attachment was created within the system.
	// If true, ReturnUrl should be baseURL + Path; otherwise ReturnUrl is just Path.
	FromInternal bool           `json:"from_internal"`
	Type         AttachmentType `json:"type"`
	ReturnUrl    string         `json:"return_url"`
}
