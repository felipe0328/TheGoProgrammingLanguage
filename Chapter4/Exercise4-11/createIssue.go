package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type CreateTemplate struct {
	Title string
	Body  string
}

func createIssue(owner, repo string) *IssueGet {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		return nil
	}

	tempFile, err := os.CreateTemp("", "issue_crud")
	if err != nil {
		return nil
	}

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	encoder := json.NewEncoder(tempFile)
	err = encoder.Encode(map[string]string{
		"title": "",
		"body":  "",
	})
	if err != nil {
		return nil
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempFile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()
	if err != nil {
		return nil
	}

	tempFile.Seek(0, 0)

	req, err := http.NewRequest(CreateIssue.Method, fmt.Sprintf(CreateIssue.Endpoint, owner, repo), tempFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp, err := doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	var result IssueGet
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	return &result
}
