package jsonstore_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
)

type testRecord struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonStoreReadMissingReturnsZeroValue(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	store := jsonstore.NewJsonStore[testRecord](filepath.Join(tmp, "missing.json"))

	got, err := store.Read()
	if err != nil {
		t.Fatalf("Read returned error for missing file: %v", err)
	}
	if got != (testRecord{}) {
		t.Fatalf("Read returned %#v, want zero value %#v", got, testRecord{})
	}
}

func TestJsonStoreWriteReadRoundTrip(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	path := filepath.Join(tmp, "nested", "records.json")
	store := jsonstore.NewJsonStore[[]testRecord](path)

	want := []testRecord{
		{Name: "alice", Age: 20},
		{Name: "bob", Age: 22},
	}

	if err := store.Write(want); err != nil {
		t.Fatalf("Write returned error: %v", err)
	}

	got, err := store.Read()
	if err != nil {
		t.Fatalf("Read returned error: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("Read returned %d records, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("record[%d] = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func TestJsonStoreReadRejectsUnknownFields(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	path := filepath.Join(tmp, "data.json")
	if err := os.WriteFile(path, []byte(`{"name":"alice","age":20,"extra":"x"}`), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	store := jsonstore.NewJsonStore[testRecord](path)
	if _, err := store.Read(); err == nil {
		t.Fatalf("Read succeeded for JSON with unknown field; want error")
	}
}

func TestJsonStoreReadRejectsTrailingData(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	path := filepath.Join(tmp, "data.json")
	content := `{"name":"alice","age":20}` + "\n" + `{"name":"bob","age":21}`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	store := jsonstore.NewJsonStore[testRecord](path)
	if _, err := store.Read(); err == nil {
		t.Fatalf("Read succeeded for JSON with trailing data; want error")
	}
}
