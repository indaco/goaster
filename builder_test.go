package goaster

import (
	"testing"
)

func TestToasterBuilder_Defaults(t *testing.T) {
	toaster := NewToasterBuilder().Build()

	if toaster == nil {
		t.Fatal("expected toaster instance, got nil")
	}
	if toaster.Position != BottomRight {
		t.Errorf("expected default position to be BottomRight, got %v", toaster.Position)
	}
	if len(toaster.Icons) == 0 {
		t.Errorf("expected default icons to be set")
	}
	if toaster.Queue() == nil {
		t.Errorf("expected queue to be initialized")
	}
}

func TestToasterBuilder_CustomConfig(t *testing.T) {
	toaster := NewToasterBuilder().
		WithPosition(TopLeft).
		WithBorder(false).
		WithAutoDismiss(false).
		Build()

	if toaster.Position != TopLeft {
		t.Errorf("expected position TopLeft, got %v", toaster.Position)
	}
	if toaster.Border != false {
		t.Errorf("expected Border=false, got %v", toaster.Border)
	}
	if toaster.AutoDismiss != false {
		t.Errorf("expected AutoDismiss=false, got %v", toaster.AutoDismiss)
	}
}

func TestToasterBuilder_WithAllOptions(t *testing.T) {
	toaster := NewToasterBuilder().
		WithVariant(AccentLight).
		WithBorder(false).
		WithRounded(true).
		WithShowIcon(false).
		WithButton(false).
		WithAutoDismiss(false).
		WithAnimation(false).
		WithProgressBar(false).
		WithPosition(TopCenter).
		WithIcon(ErrorLevel, "<svg>Error</svg>").
		Build()

	if toaster.Variant != AccentLight {
		t.Errorf("expected Variant to be AccentLight, got %v", toaster.Variant)
	}
	if toaster.Border {
		t.Errorf("expected Border=false")
	}
	if !toaster.Rounded {
		t.Errorf("expected Rounded=true")
	}
	if toaster.ShowIcon {
		t.Errorf("expected ShowIcon=false")
	}
	if toaster.Button {
		t.Errorf("expected Button=false")
	}
	if toaster.AutoDismiss {
		t.Errorf("expected AutoDismiss=false")
	}
	if toaster.Animation {
		t.Errorf("expected Animation=false")
	}
	if toaster.ProgressBar {
		t.Errorf("expected ProgressBar=false")
	}
	if toaster.Position != TopCenter {
		t.Errorf("expected Position=TopCenter, got %v", toaster.Position)
	}
	if icon, ok := toaster.Icons[ErrorLevel]; !ok || icon != "<svg>Error</svg>" {
		t.Errorf("expected custom ErrorLevel icon to be set, got %q", icon)
	}
}

func TestToasterBuilder_WithOptionsMerge(t *testing.T) {
	builder := NewToasterBuilder().
		WithBorder(false)

	// Apply additional options
	toaster := builder.WithOptions(
		WithAutoDismiss(false),
		WithAnimation(false),
	).Build()

	if toaster.Border != false {
		t.Errorf("expected Border=false after builder chain")
	}
	if toaster.AutoDismiss != false {
		t.Errorf("expected AutoDismiss=false via WithOptions")
	}
	if toaster.Animation != false {
		t.Errorf("expected Animation=false via WithOptions")
	}
}

func TestToasterBuilder_WithIconNilMap(t *testing.T) {
	// Manually simulate icon map being nil
	builder := NewToasterBuilder()
	builder.Build().Icons = nil // forcibly set to nil

	// Set new icon — this should initialize the map
	toaster := builder.WithIcon(WarningLevel, "<svg>Warning</svg>").Build()

	icon, ok := toaster.Icons[WarningLevel]
	if !ok || icon != "<svg>Warning</svg>" {
		t.Errorf("expected WarningLevel icon to be set correctly, got %q", icon)
	}
}

func TestToasterBuilder_Build_FallbackToDefaultPosition(t *testing.T) {
	// Simulate missing Position
	builder := NewToasterBuilder()
	builder.Build().Position = "" // forcibly unset
	toaster := builder.Build()

	if toaster.Position != BottomRight {
		t.Errorf("expected fallback Position to be BottomRight, got %v", toaster.Position)
	}
}

func TestToasterBuilder_Build_FallbackToDefaultIcons(t *testing.T) {
	builder := NewToasterBuilder()
	builder.Build().Icons = nil // forcibly unset
	toaster := builder.Build()

	if len(toaster.Icons) == 0 {
		t.Errorf("expected fallback Icons to be set")
	}
}

func TestToasterBuilder_Build_FallbackToNewQueue_Clean(t *testing.T) {
	builder := NewToasterBuilder()
	toaster := builder.Build()

	// Attempt to dequeue from an empty queue
	_, err := toaster.Queue().Dequeue()
	if err == nil {
		t.Errorf("expected error when dequeuing from an empty queue, got nil")
	}

	// Make sure the queue is initialized
	if toaster.Queue() == nil {
		t.Errorf("expected Queue to be initialized")
	}
}

func TestToasterBuilder_Build_SetsDefaultPositionIfEmpty(t *testing.T) {
	builder := NewToasterBuilder()
	builder.Build().Position = "" // simulate unset

	toaster := builder.Build()
	if toaster.Position != BottomRight {
		t.Errorf("expected Position to default to BottomRight, got %v", toaster.Position)
	}
}

func TestToasterBuilder_Build_SetsDefaultIconsIfNil(t *testing.T) {
	builder := NewToasterBuilder()
	builder.Build().Icons = nil // simulate nil

	toaster := builder.Build()
	if len(toaster.Icons) == 0 {
		t.Errorf("expected Icons to be initialized with defaults")
	}
}

func TestToasterBuilder_Build_InitializesQueueIfNil(t *testing.T) {
	builder := NewToasterBuilder()

	// Manually set queue to nil to simulate uninitialized queue
	builder.Build().queue = nil

	// Call Build again to trigger fallback
	toaster := builder.Build()

	if toaster.Queue() == nil {
		t.Errorf("expected Queue to be initialized, got nil")
	}

	// Ensure it functions properly
	_, err := toaster.Queue().Dequeue()
	if err == nil {
		t.Errorf("expected error when dequeuing from empty queue, got nil")
	}
}
