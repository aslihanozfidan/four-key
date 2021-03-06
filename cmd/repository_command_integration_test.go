package cmd

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestAddCommand_WhenRunsWithEmptyArgs_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "add",
		"--cloneAddress", "",
		"--team", "",
		"--releaseTagPattern", "",
		"--fixCommitPatterns", "")

	out, err := cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "You must specify a repository to clone"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}

func TestAddCommand_WhenRunsWithCorrectCloneAddressButWrongArgs_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "add",
		"--cloneAddress", "https://github.com/Trendyol/four-key.git",
		"--teams=")

	var got bytes.Buffer
	cmd.Stderr = &got
	_, _ = cmd.Output()

	want := "unknown flag: --teams"
	if !strings.Contains(got.String(), want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got.String(), want)
	}
}

func TestAddCommand_WhenRunsWithCorrectArgs_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "add",
		"--cloneAddress", "https://github.com/Trendyol/four-key.git",
		"--team", "trendyol-team",
		"--releaseTagPattern", "release-v",
		"--fixCommitPatterns", "fix", "-f", "hotfix")

	out, err := cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "successfully added your repository to config file"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}

func TestRemoveCommand_WhenRunsWithWrongArg_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "remove",
		"--repositorysss", "")

	var got bytes.Buffer
	cmd.Stderr = &got
	_, _ = cmd.Output()

	want := "unknown flag: --repositorysss"
	if !strings.Contains(got.String(), want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got.String(), want)
	}
}

func TestRemoveCommand_WhenRunsWithCorrectArg_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "add",
		"--cloneAddress", "https://github.com/Trendyol/four-key.git",
		"--team", "trendyol-team",
		"--releaseTagPattern", "release-v",
		"--fixCommitPatterns", "fix", "-f", "hotfix")

	out, err := cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	cmd = exec.Command("go", "run", "../main.go", "remove",
		"--repository", "four-key")

	out, err = cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "successfully removed four-key repository from the config file"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}

func TestRemoveCommand_WhenRunsWithDoesNotExistRepository_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "remove",
		"--repository", "four-key")

	out, err := cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "The four-key repository does not exist!"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}

func TestListCommand_WhenRunsCorrectly_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cmd := exec.Command("go", "run", "../main.go", "list")

	out, err := cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "repository/repositories has been found"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}
