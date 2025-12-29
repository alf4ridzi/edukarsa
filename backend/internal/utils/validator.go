package utils

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/domain"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func GetExtension(file *multipart.FileHeader) string {
	extension := strings.ToLower(filepath.Ext(file.Filename))
	return extension
}

func ValidateExtension(extension string) bool {
	return config.AllowedExtensions[extension]
}

func ValidateFileSize(file *multipart.FileHeader) bool {
	return file.Size > config.MaxUploadSizeBytes()
}

func ValidateUpload(file *multipart.FileHeader) error {
	if !ValidateExtension(GetExtension(file)) {
		return domain.ErrInvalidExtension
	}

	if !ValidateFileSize(file) {
		return domain.ErrFileSizeTooBig
	}

	return nil
}
