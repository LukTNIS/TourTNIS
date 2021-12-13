package txt

import (
	"os"
)

func WriteFileByLine(fileName string, textInput string) error {
	var f *os.File
	var err error
	if _, err = os.Stat(fileName); err == nil {
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	} else {
		f, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(textInput + "\n")

	if err2 != nil {
		return err2
	}
	return nil
}
