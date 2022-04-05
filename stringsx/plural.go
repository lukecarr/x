package stringsx

import (
	"fmt"
	"strings"
)

// ES_ENDINGS are word endings/suffixes that indicate "es" should be added to
// pluralize the word, rather than "s".
//
// For example, the plural of "church" is "churches" (not "churchs") because
// it has a suffix in ES_ENDINGS ("ch").
var ES_ENDINGS = []string{"s", "ss", "z", "ch", "sh", "x"}

// Plural converts an (English) word into its plural form, if the supplied
// integer (n) is not equal to 1.
func Plural(word string, n int) string {
	if len(word) == 1 {
		return word
	}

	suffix := "s"

	for _, ending := range ES_ENDINGS {
		if strings.HasSuffix(word, ending) {
			suffix = "es"
			break
		}
	}

	return fmt.Sprintf("%s%s", word, suffix)
}
