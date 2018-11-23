package logging

import (
	"PushServer/pkg/file"
	"fmt"
	"os"
)

func IsNotExistMkDir(filePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err :%v", err)
	}

	src := dir + "/" + filePath
	perm := file.HasPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	return nil
}
