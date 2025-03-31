<h1 align="center" style="font-size: 2.5rem;">
  goaster
</h1>
<p align="center">
  <a href="https://github.com/indaco/goaster/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
  </a>
  &nbsp;
  <a href="https://goreportcard.com/report/github.com/indaco/goaster/" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/goaster" alt="go report card" />
  </a>
  &nbsp;
  <a href="https://pkg.go.dev/github.com/indaco/goaster/" target="_blank">
      <img src="https://pkg.go.dev/badge/github.com/indaco/goaster/.svg" alt="go reference" />
  </a>
  &nbsp;
  <a href="https://www.jetify.com/devbox/docs/contributor-quickstart/">
    <img src="https://www.jetify.com/img/devbox/shield_moon.svg" alt="Built with Devbox" />
  </a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#custom-icons">Custom Icons</a> •
  <a href="#theming">Theming</a> •
  <a href="#examples">Examples</a> •
  <a href="#contributing">Contributing</a>
</p>

A configurable, themeable and non-intrusive server-rendered toast notification component for Go web applications. Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

<div style="display: flex; justify-content: center;">
   <img src="https://raw.githubusercontent.com/indaco/goaster/gh-assets/demo.gif" alt="Image" >
</div>

### Features

- **No External Dependencies**: Built with native Go and the [`templ`](https://templ.guide) library—no JavaScript or frontend dependencies required.
- **Multiple Toasts**: Supports displaying multiple toast notifications simultaneously.
- **Highly Configurable**: Customize appearance (bordered, rounded), behavior, and position to fit your UI.
- **Variants**: Includes style variants such as `Accent`, `AccentLight`, and `AccentDark`.
- **Themeable**: Easily theme your toasts using CSS variables to match your app’s design.
- **Icon Support**: Comes with default SVG icons for common toast types (success, error, info, etc.)—or use your own.
- **Flexible Positioning**: Display toast messages in any corner (top-right, bottom-left, etc.).
- **Auto-Dismiss Progress Bar**: Visual progress indicator for toast duration (when auto-dismiss is enabled).
- **Smooth Animations**:
  - _Built-in Animations_: Default entrance and exit transitions.
  - _Custom Animations_: Define your own via CSS variables or external animation libraries.
  - _Animation Control_: Fine-tune timing, easing, delay, and effects with CSS variables.
  - _Disable Option_: Disable all animations when needed (e.g., for accessibility or testing).

## Installation

Ensure your project is using Go Modules.

To install the module, use the `go get` command:

```sh
go get github.com/indaco/goaster@latest
```

## Usage

Import `goaster` module into your project:

```go
import "github.com/indaco/goaster"
```

### Creating a Toaster

Initialize a `Toaster` with default settings:

```go
toaster := goaster.NewToaster()
```

or customize the toaster with options:

```go
toaster := goaster.NewToaster(
   goaster.WithBorder(false),
   goaster.WithPosition(goaster.TopRight),
   goaster.WithAutoDismiss(false)
   // ...
)
```

### Displaying Toast Messages

Display different levels of toast messages:

```go
toaster.Default("Sample message.")
toaster.Success("Operation completed successfully.")
toaster.Error("An error occurred.")
toaster.Info("Here's some information.")
toaster.Warning("This is a warning message.")
```

### Add Toast Messages to the queue and display them

Multiple toast messages in the queue:

```go
toaster.PushDefault("Sample message.")
toaster.PushSuccess("Operation completed successfully.")
toaster.PushError("An error occurred.")
toaster.PushInfo("Here's some information.")
toaster.PushWarning("This is a warning message.")

toaster.RenderAll()
```

## Custom Icons

Specify custom SVG icons for each toast level:

```go
toaster := goaster.NewToaster(
   goaster.WithIcon(toast.SuccessLevel, "<svg>...</svg>"),
   goaster.WithIcon(toast.ErrorLevel, "<svg>...</svg>"),
)
```

## Theming

Customizing the appearance of `goaster` notifications to align with your design preferences is both straightforward and flexible, accomplished by using CSS custom properties (CSS variables) prefixed with `gtt`.

See the [CSS Custom Properties](./docs/css-props.md) reference for full details.

## Use with Go's `template/html`

To facilitate integration with Go's `template/html` standard library, `goaster` includes a dedicated `HTMLGenerator` type to seamlessly integrate toast notifications into web applications built with Go's `html/template` standard library.

> **Note**: See the [Examples](#examples) section to learn how to use `goaster` with `templ` and `html/template`.

## Examples

👉 [Check out the examples with setup instructions](examples/)

## Contributing

Contributions are welcome!

See the [Contributing Guide](/CONTRIBUTING.md) for setup instructions.

## License

This project is licensed under the MIT License – see the [LICENSE](./LICENSE) file for details.
