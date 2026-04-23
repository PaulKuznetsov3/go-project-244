package code

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenDiff_FromTempFiles(t *testing.T) {
	dir := t.TempDir()
	file1 := filepath.Join(dir, "profile1.json")
	file2 := filepath.Join(dir, "profile2.json")

	profile1 := `{
		"name": "John",
		"age": 30
	}`
	
	profile2 := `{
		"name": "Jane",
		"age": 30,
		"email": "jane@example.com"
	}`

	if err := os.WriteFile(file1, []byte(profile1), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(file2, []byte(profile2), 0644); err != nil {
		t.Fatal(err)
	}

	got, _ := GenDiff(file1, file2, "stylish")
	want := "{\n    age: 30\n  + email: jane@example.com\n  - name: John\n  + name: Jane\n}"

	if got != want {
		t.Fatalf("GenDiff:\n got:\n%q\nwant:\n%q", got, want)
	}
}