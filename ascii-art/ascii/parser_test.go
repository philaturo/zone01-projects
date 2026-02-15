package ascii

import "testing"

func TestLoadBanner(t *testing.T) {
	data, err := LoadBannerFile("../banners/standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	banner, err := LoadBanner(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(banner) != 95 {
		t.Fatalf("expected 95 characters, got %d", len(banner))
	}

	if len(banner['A']) != charHeight {
		t.Fatalf("expected %d lines for 'A'", charHeight)
	}
}
