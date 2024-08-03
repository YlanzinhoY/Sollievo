package processfile

import (
	"io"
	"os"
)

func copyFile(srcFileName, destFile string) error {

	srcFile, err := os.Open(srcFileName)

	if err != nil {
		return err
	}

	defer srcFile.Close()

	dstFile, err := os.Create(destFile)

	if err != nil {
		return err
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)

	if err != nil {
		return err
	}

	return nil

}
