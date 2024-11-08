package utils

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

var allowedFileExtensions = map[string]bool{
	".pdf":  true,
	".docx": true,
}


// ValidateFileExtension checks if the file extension is allowed
func ValidateFileExtension(fileHeader *multipart.FileHeader) error {
    ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
    if !allowedFileExtensions[ext] {
        return errors.New("unsupported file type, only PDF and DOCX allowed")
    }
    return nil
}
