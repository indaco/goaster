package goaster

// Option type for functional options pattern, allowing customization of the Toaster instance.
type Option func(*Toaster)

// WithVariant configures the style variant for the toast.
func WithVariant(variant Variant) Option {
	return func(tp *Toaster) {
		tp.Variant = variant
	}
}

// WithBorder configures the presence of a border around the toast.
func WithBorder(border bool) Option {
	return func(tp *Toaster) {
		tp.Border = border
	}
}

// WithShowIcon configures whether the toast should display the icon.
func WithShowIcon(icon bool) Option {
	return func(tp *Toaster) {
		tp.ShowIcon = icon
	}
}

// WithButton configures whether the toast should display the close button.
func WithButton(button bool) Option {
	return func(tp *Toaster) {
		tp.Button = button
	}
}

// WithAutoDismiss configures whether the toast should auto-dismiss.
func WithAutoDismiss(autoDismiss bool) Option {
	return func(tp *Toaster) {
		tp.AutoDismiss = autoDismiss
	}
}

// WithAnimation configures whether the toast should use animations.
func WithAnimation(animation bool) Option {
	return func(tp *Toaster) {
		tp.Animation = animation
	}
}

// WithProgressBar configures whether the toast should display a progressbar.
func WithProgressBar(progressbar bool) Option {
	return func(tp *Toaster) {
		tp.ProgressBar = progressbar
	}
}

// WithPosition sets the position of the toast notifications.
func WithPosition(position Position) Option {
	return func(t *Toaster) {
		t.Position = position
	}
}

// WithIcon sets the icon for a specific toast level.
func WithIcon(level Level, iconSVG string) Option {
	return func(t *Toaster) {
		t.Icons[level] = iconSVG
	}
}
