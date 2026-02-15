package ascii

import "testing"

func TestLoadBannerFile(t *testing.T) {
	data, err := LoadBannerFile("../banners/standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(data) == 0 {
		t.Fatal("expected banner file content, got empty data")
	}
}
