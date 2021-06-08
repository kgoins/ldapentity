package entity

import (
	"strings"

	hashset "github.com/kgoins/hashset/pkg"
)

func BuildAttributeFilter(filterParts []string) hashset.StrHashset {
	set := hashset.NewStrHashset()

	for _, attr := range filterParts {
		set.Add(strings.ToLower(attr))
	}

	return set
}
