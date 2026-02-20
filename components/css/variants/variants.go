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
	case "filled":
		return variantFilled()
	case "outlined":
		return variantOutlined()
	case "soft":
		return variantSoft()
	case "minimal":
		return variantMinimal()
	case "brutalist":
		return variantBrutalist()
	case "retro":
		return variantRetro()
	case "neon":
		return variantNeon()
	default:
		return templ.NopComponent
	}
}
