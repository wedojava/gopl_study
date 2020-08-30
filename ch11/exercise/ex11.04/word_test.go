package word

import (
	"math/rand"
	"testing"
	"time"
)

var puncspaces = []rune{
	' ', ',', '.', '!', '?', '　', '、', '。', '！', '？',
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}

	if n == 0 {
		return string(runes)
	}
	// runes' len is n, so pos is less than n, cause runes is palindromed
	// so, runes splice any position will still palidromed.
	pos := rng.Intn(n)
	runes = append(runes[:pos], append([]rune{
		puncspaces[rng.Intn(len(puncspaces))]}, runes[pos:]...)...)
	return string(runes)

	// m := rng.Intn(n + 1)
	// var temp []rune
	// temp = append(temp, runes[:m]...)
	// temp = append(temp, ' ') // Insert space
	// temp = append(temp, runes[m:]...)
	//
	// m = rng.Intn(n + 1)
	// runes = append(make([]rune, 0), temp[:m]...)
	// runes = append(runes, ',') // Insert punctuation
	// runes = append(runes, temp[m:]...)
	// return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
