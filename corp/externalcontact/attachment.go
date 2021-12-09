package externalcontact

type AttachmentType string

const (
	AttachmentImage AttachmentType = "image"
	AttachmentLink  AttachmentType = "link"
	AttachmentMinip AttachmentType = "miniprogram"
	AttachmentVideo AttachmentType = "video"
	AttachmentFile  AttachmentType = "file"
)
