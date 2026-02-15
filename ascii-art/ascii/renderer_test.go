package ascii

import "testing"

// mockBanner creates a fake banner with height = 8 for all ASCII chars
func mockBanner() map[rune][]string {
	banner := make(map[rune][]string)

	for r := rune(32); r <= 126; r++ {
		banner[r] = []string{
			string(r),
			string(r),
			string(r),
			string(r),
			string(r),
			string(r),
			string(r),
			string(r),
		}
	}

	return banner
}

func TestRenderEmptyInput(t *testing.T) {
	banner := mockBanner()
	result := RenderText("", banner)

	if len(result) != 0 {
		t.Fatalf("expected 0 lines, got %d", len(result))
	}
}

func TestRenderNewlineOnly(t *testing.T) {
	banner := mockBanner()
	result := RenderText("\n", banner)

	if len(result) != 2 {
		t.Fatalf("expected 2 line, got %d", len(result))
	}
	if result[0] != "" {
		t.Fatalf("expected empty line")
	}
}


func TestRenderSingleWord(t *testing.T) {
	banner := mockBanner()
	result := RenderText("Hi", banner)

	if len(result) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(result))
	}
	for i, line := range result {
		if line == "" {
			t.Fatalf("line %d is empty, expected ASCII art", i)
		}
	}
}

func TestRenderTrailingNewline(t *testing.T) {
	banner := mockBanner()
	result := RenderText("Hi\n", banner)

	if len(result) != 10 {
		t.Fatalf("expected 10 lines, got %d", len(result))
	}
}

func TestRenderTwoLines(t *testing.T) {
	banner := mockBanner()
	result := RenderText("Hi\nThere", banner)

	if len(result) != 17 {
		t.Fatalf("expected 17 lines, got %d", len(result))
	}
}

func TestRenderDoubleNewline(t *testing.T) {
	banner := mockBanner()
	result := RenderText("Hi\n\nThere", banner)

	if len(result) != 18 {
		t.Fatalf("expected 18 lines, got %d", len(result))
	}
}
