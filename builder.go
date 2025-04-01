package goaster

// ToasterBuilder builds a Toaster instance using chainable methods.
type ToasterBuilder struct {
	toaster *Toaster
}

// NewToasterBuilder creates a new builder with default values.
func NewToasterBuilder() *ToasterBuilder {
	return &ToasterBuilder{
		toaster: ToasterDefaults(),
	}
}

func (b *ToasterBuilder) WithVariant(variant Variant) *ToasterBuilder {
	b.toaster.Variant = variant
	return b
}

func (b *ToasterBuilder) WithBorder(border bool) *ToasterBuilder {
	b.toaster.Border = border
	return b
}

func (b *ToasterBuilder) WithRounded(rounded bool) *ToasterBuilder {
	b.toaster.Rounded = rounded
	return b
}

func (b *ToasterBuilder) WithShowIcon(showIcon bool) *ToasterBuilder {
	b.toaster.ShowIcon = showIcon
	return b
}

func (b *ToasterBuilder) WithButton(button bool) *ToasterBuilder {
	b.toaster.Button = button
	return b
}

func (b *ToasterBuilder) WithAutoDismiss(autoDismiss bool) *ToasterBuilder {
	b.toaster.AutoDismiss = autoDismiss
	return b
}

func (b *ToasterBuilder) WithAnimation(animation bool) *ToasterBuilder {
	b.toaster.Animation = animation
	return b
}

func (b *ToasterBuilder) WithProgressBar(progressBar bool) *ToasterBuilder {
	b.toaster.ProgressBar = progressBar
	return b
}

func (b *ToasterBuilder) WithPosition(position Position) *ToasterBuilder {
	b.toaster.Position = position
	return b
}

func (b *ToasterBuilder) WithIcon(level Level, iconSVG string) *ToasterBuilder {
	if b.toaster.Icons == nil {
		b.toaster.Icons = make(map[Level]string)
	}
	b.toaster.Icons[level] = iconSVG
	return b
}

func (b *ToasterBuilder) WithOptions(options ...Option) *ToasterBuilder {
	for _, opt := range options {
		opt(b.toaster)
	}
	return b
}

// Build finalizes and returns the configured Toaster instance.
func (b *ToasterBuilder) Build() *Toaster {
	// Ensure Position is not empty (should always be valid)
	if b.toaster.Position == "" {
		b.toaster.Position = BottomRight
	}

	// Ensure Icons map is initialized
	if b.toaster.Icons == nil {
		b.toaster.Icons = defaultIcons()
	}

	// Ensure Queue is initialized
	if b.toaster.queue == nil {
		b.toaster.queue = NewQueue()
	}

	return b.toaster
}
