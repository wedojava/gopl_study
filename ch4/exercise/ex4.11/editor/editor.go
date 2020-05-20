package editor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

func Edit(value map[string]string) error {
	editor := getEditorName()

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}

	tempFileName := tempFile.Name()
	defer os.Remove(tempFileName)

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(value); err != nil {
		return err
	}

	if err = tempFile.Close(); err != nil {
		return err
	}

	cmd := exec.Command(editor, tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return err
	}

	edited, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		return err
	}

	// Some editors, such as Windows Notepad, always add a BOM when saving UTF-8.
	// json.Unmarshal does not support data with UTF-8 BOM, so delete it beforehand.
	if err = json.Unmarshal(removeUTF8BOM(edited), &value); err != nil {
		return err
	}
	return nil
}

func getEditorName() string {
	editor := os.Getenv("GIT_EDITOR")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		editor = "vi"
	}
	return editor
}

func removeUTF8BOM(s []byte) []byte {
	utf8BOM := []byte{239, 187, 191}
	return bytes.TrimPrefix(s, utf8BOM)
}
