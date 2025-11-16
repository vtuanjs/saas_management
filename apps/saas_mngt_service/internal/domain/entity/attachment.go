package entity

type AttachmentType int

const (
	AttachmentTypePDF AttachmentType = iota + 1
	AttachmentTypeImage
	AttachmentTypeVideo
)

type Attachment struct {
	ID   string
	Path string
	Type AttachmentType
}
