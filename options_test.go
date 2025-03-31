package goaster

import (
	"testing"
)

func TestToasterOptions(t *testing.T) {
	tests := []struct {
		name     string
		options  []Option
		expected Toaster
	}{
		{
			name:    "Default options",
			options: nil,
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithVariant accent",
			options: []Option{
				WithVariant(AccentLight),
			},
			expected: Toaster{
				Variant:     AccentLight,
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithBorder false",
			options: []Option{
				WithBorder(false),
			},
			expected: Toaster{
				Border:      false,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithRounded false",
			options: []Option{
				WithRounded(false),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     false,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithShowIcon false",
			options: []Option{
				WithShowIcon(false),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    false,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithButton false",
			options: []Option{
				WithButton(false),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      false,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithAutoDismiss false",
			options: []Option{
				WithAutoDismiss(false),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: false,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithAnimation false",
			options: []Option{
				WithAnimation(false),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   false,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithProgressBar false",
			options: []Option{
				WithProgressBar(false),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: false,
				Position:    BottomRight,
			},
		},
		{
			name: "WithPosition TopLeft",
			options: []Option{
				WithPosition(TopLeft),
			},
			expected: Toaster{
				Border:      true,
				Rounded:     true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Position:    TopLeft,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toaster := NewToaster(tt.options...)

			// Check each field in expected Toaster struct
			if toaster.Variant != tt.expected.Variant {
				t.Errorf("Variant: expected %v, got %v", tt.expected.Variant, toaster.Variant)
			}
			if toaster.Border != tt.expected.Border {
				t.Errorf("Border: expected %v, got %v", tt.expected.Border, toaster.Border)
			}
			if toaster.Rounded != tt.expected.Rounded {
				t.Errorf("Rounded: expected %v, got %v", tt.expected.Rounded, toaster.Rounded)
			}
			if toaster.ShowIcon != tt.expected.ShowIcon {
				t.Errorf("ShowIcon: expected %v, got %v", tt.expected.ShowIcon, toaster.ShowIcon)
			}
			if toaster.Button != tt.expected.Button {
				t.Errorf("Button: expected %v, got %v", tt.expected.Button, toaster.Button)
			}
			if toaster.AutoDismiss != tt.expected.AutoDismiss {
				t.Errorf("AutoDismiss: expected %v, got %v", tt.expected.AutoDismiss, toaster.AutoDismiss)
			}
			if toaster.Position != tt.expected.Position {
				t.Errorf("Position: expected %v, got %v", tt.expected.Position, toaster.Position)
			}
		})
	}
}

func TestWithVariant(t *testing.T) {
	toaster := NewToaster(WithVariant(AccentDark))
	if toaster.Variant != AccentDark {
		t.Errorf("expected Variant to be %q, got %q", AccentDark, toaster.Variant)
	}
}

func TestWithBorder(t *testing.T) {
	toaster := NewToaster(WithBorder(false))
	if toaster.Border != false {
		t.Errorf("expected Border=false, got %v", toaster.Border)
	}
}

func TestWithRounded(t *testing.T) {
	toaster := NewToaster(WithRounded(false))
	if toaster.Rounded != false {
		t.Errorf("expected Rounded=false, got %v", toaster.Rounded)
	}
}

func TestWithShowIcon(t *testing.T) {
	toaster := NewToaster(WithShowIcon(false))
	if toaster.ShowIcon != false {
		t.Errorf("expected ShowIcon=false, got %v", toaster.ShowIcon)
	}
}

func TestWithButton(t *testing.T) {
	toaster := NewToaster(WithButton(false))
	if toaster.Button != false {
		t.Errorf("expected Button=false, got %v", toaster.Button)
	}
}

func TestWithAutoDismiss(t *testing.T) {
	toaster := NewToaster(WithAutoDismiss(false))
	if toaster.AutoDismiss != false {
		t.Errorf("expected AutoDismiss=false, got %v", toaster.AutoDismiss)
	}
}

func TestWithAnimation(t *testing.T) {
	toaster := NewToaster(WithAnimation(false))
	if toaster.Animation != false {
		t.Errorf("expected Animation=false, got %v", toaster.Animation)
	}
}

func TestWithProgressBar(t *testing.T) {
	toaster := NewToaster(WithProgressBar(false))
	if toaster.ProgressBar != false {
		t.Errorf("expected ProgressBar=false, got %v", toaster.ProgressBar)
	}
}

func TestWithPosition(t *testing.T) {
	toaster := NewToaster(WithPosition(TopCenter))
	if toaster.Position != TopCenter {
		t.Errorf("expected Position=TopCenter, got %v", toaster.Position)
	}
}

func TestWithIcon(t *testing.T) {
	customSVG := `<svg></svg>`
	toaster := NewToaster(WithIcon(ErrorLevel, customSVG))

	icon, exists := toaster.Icons[ErrorLevel]
	if !exists {
		t.Errorf("expected icon for ErrorLevel to exist")
	}
	if icon != customSVG {
		t.Errorf("expected icon=%q, got %q", customSVG, icon)
	}
}
