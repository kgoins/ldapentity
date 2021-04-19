package entity

import (
	"strings"

	hashset "github.com/kgoins/hashset/pkg"
)

// BuildEntity is a wrapper for BuildEntityFromAttrList
// that doesn't require an attribute filter and returns all entities.
func BuildEntity(entityLines []string) Entity {
	return BuildEntityFromAttrList(
		entityLines,
		nil,
	)
}

// BuildEntityFromAttrList will construct an Entity from a list of attribute strings
// and filter out all attributes not in `includeAttrs`.
// Either a null or empty HashSetStr value in `includeAttrs` will include all attributes.
// The `includeAttrs` argument must contain lowercase string values.
func BuildEntityFromAttrList(entityLines []string, includeAttrs *hashset.StrHashset) Entity {
	entity := NewEntity()
	hasAttrFilter := (includeAttrs != nil) && !includeAttrs.IsEmpty()

	// Ensure that we always pull a DN if possible
	if hasAttrFilter {
		includeAttrs.Add("dn")
		includeAttrs.Add("distinguishedname")
	}

	for _, line := range entityLines {
		attr, err := BuildAttributeFromLine(line)
		if err != nil {
			continue
		}

		if !hasAttrFilter {
			entity.AddAttribute(attr)
		} else {
			if includeAttrs.Contains(strings.ToLower(attr.Name)) {
				entity.AddAttribute(attr)
			}
		}
	}

	if dn, found := entity.GetDN(); found == true {
		Logger.Info("Built entity: " + dn.Value.GetSingleValue())
	}

	return entity
}
