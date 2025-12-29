package config

var AllowedExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".pdf":  true,
	".doc":  true,
	".docx": true,
}

func MaxUploadSizeBytes() int64 {
	return AppConfig.MaxUploadSize << 20
}
