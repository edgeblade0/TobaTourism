package attachment

import "mime/multipart"

type Usecase interface {
	InsertAttachment(files []*multipart.FileHeader, FilePath string, Type int) (int64, error)
	//GetAttachment(AttachmentID string) error
}
