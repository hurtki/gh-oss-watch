package services

import (
	"fmt"
	"testing"
)

// Unit tests for ParseRepoString
func TestParseRepoString(t *testing.T) {
	input := []struct {
		name           string
		input          string
		excpectedOwner string
		excpectedRepo  string
		HasError       bool
	}{
		{name: "owner/repo succsess", input: "owner/repo", excpectedOwner: "owner", excpectedRepo: "repo", HasError: false},
		{name: "link succsess", input: "https://github.com/owner/repo", excpectedOwner: "owner", excpectedRepo: "repo", HasError: false},
		{name: "link to .git succsess", input: "https://github.com/owner/repo.git", excpectedOwner: "owner", excpectedRepo: "repo", HasError: false},
		{name: "slashes succses", input: "///owner/repo///", excpectedOwner: "owner", excpectedRepo: "repo", HasError: false},
		{name: "link to owner wrong", input: "https://github.com/owner", excpectedOwner: "", excpectedRepo: "", HasError: true},
		{name: "only owner wrong", input: "owner", excpectedOwner: "", excpectedRepo: "", HasError: true},
		{name: "slashes wrong", input: "///owner////", excpectedOwner: "", excpectedRepo: "", HasError: true},
	}

	for i, tc := range input {
		owner, repo, err := ParseRepoString(tc.input)
		hasError := err != nil
		if owner != tc.excpectedOwner || repo != tc.excpectedRepo || hasError != tc.HasError {
			fatal := fmt.Sprintf("test %d (%s): expected owner=%q repo=%q error=%v, got owner=%q repo=%q error=%v",
				i, tc.name, tc.excpectedOwner, tc.excpectedRepo, tc.HasError, owner, repo, hasError)
			t.Fatal(fatal)
		}
	}
}
