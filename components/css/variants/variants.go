package variants

import (
	"github.com/a-h/templ"
)

func UseVariant(variant string) templ.Component {
	switch variant {
	case "accent":
		return variantAccent()
	case "accent-light":
		return variantAccentLight()
	case "accent-dark":
		return variantAccentDark()
	default:
		return templ.NopComponent
	}
}
