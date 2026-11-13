package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestRecordingFileEncodingContract(t *testing.T) {
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	want := RecordingFile{
		Status:            RecordingStatusEncoding,
		RelativePath:      "show/episode.mp4",
		PartialPath:       ".staging/capture.ts",
		EncodingStartedAt: &now,
		SourceBytes:       1234,
	}
	b, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got RecordingFile
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got.Status != RecordingStatusEncoding || got.SourceBytes != want.SourceBytes || got.PartialPath != want.PartialPath {
		t.Fatalf("round trip: got %+v want %+v", got, want)
	}
}
