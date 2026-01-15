package tokenizer

import "testing"

func TestTokenizer_Apostrophes(t *testing.T) {
	input := "I don't know ' why '"
	tokens := Tokenize(input)

	got := []string{}
	for _, tok := range tokens {
		got = append(got, tok.Value)
	}

	want := []string{"I", "don't", "know", "'", "why", "'"}

	if len(got) != len(want) {
		t.Fatalf("expected %v tokens, got %v", want, got)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("token %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}
