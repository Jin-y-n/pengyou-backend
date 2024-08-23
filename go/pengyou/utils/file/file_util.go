package file

import (
	"os"
	"pengyou/utils/log"
)

func Close(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Logger.Error("file close failed")
		return
	}
}
