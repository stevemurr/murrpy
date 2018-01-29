package fsscan

import (
	"encoding/base32"
	"murrpy/model"
	"path/filepath"

	"github.com/karrick/godirwalk"
)

var (
	supportedFormats = []string{
		".avi",
		".mpg",
		".mov",
		".flv",
		".wmv",
		".asf",
		".mpeg",
		".m4v",
		".divx",
		".mp4",
		".mkv",
	}
)

func endsWith(path string) bool {
	for _, ending := range supportedFormats {
		if filepath.Ext(path) == ending {
			return true
		}
	}
	return false
}

// Scan a directory for files
func Scan(dir string) []*model.Media {
	files := make([]*model.Media, 0)
	godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			_, tail := filepath.Split(osPathname)
			if endsWith(tail) && tail[0] != '.' {
				m := &model.Media{
					Path: osPathname,
					Hash: base32.StdEncoding.EncodeToString([]byte(osPathname)),
				}
				files = append(files, m)
			}
			return nil
		},
	})
	return files
}
