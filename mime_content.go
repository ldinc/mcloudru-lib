package mcloudru_ru

import (
	"mime"
)

func GetContentType(ext string) string {
	ctype := mime.TypeByExtension(ext)
	if ctype != "" {
		return ctype
	} else {
		// return default binary type
		return "application/octet-stream"
	}
}
