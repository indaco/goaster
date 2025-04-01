<h1 align="center">
  goaster
</h1>
<h2 align="center" style="font-size: 1.5em;">
  toast notification component for Go web applications.
</h2>
<p align="center">
  <a href="https://github.com/indaco/goaster/actions/workflows/ci.yml" target="_blank">
      <img src="https://github.com/indaco/goaster/actions/workflows/ci.yml/badge.svg" alt="CI" />
    </a>
    <a href="https://coveralls.io/github/indaco/goaster?branch=main" target="_blank">
      <img src="https://coveralls.io/repos/github/indaco/goaster/badge.svg?branch=main" alt="Coverage Status" />
    </a>
  <a href="https://goreportcard.com/report/github.com/indaco/goaster/" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/goaster" alt="go report card" />
  </a>
  <a href="https://badge.fury.io/gh/indaco%2Fgoaster">
    <img src="https://badge.fury.io/gh/indaco%2Fgoaster.svg" alt="GitHub version" height="18">
  </a>
  <a href="https://pkg.go.dev/github.com/indaco/goaster/" target="_blank">
      <img src="https://pkg.go.dev/badge/github.com/indaco/goaster/.svg" alt="go reference" />
  </a>
   <a href="https://github.com/indaco/goaster/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
  </a>
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

You can create a `Toaster` instance using one of the following approaches:

#### 1. Use default settings

```go
func HandleSingle(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.ToasterDefaults()
  templ.Handler(pages.HomePage(toaster)).ServeHTTP(w, r)
}
```

#### 2. Use functional options

Customize the toaster behavior on creation:

```go
func HandleSingle(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.NewToaster(
    goaster.WithBorder(false),
    goaster.WithPosition(goaster.TopRight),
    goaster.WithAutoDismiss(false),
    // ...
  )
  templ.Handler(pages.HomePage(toaster)).ServeHTTP(w, r)
}
```

#### 3. Use the builder pattern

More readable when chaining multiple settings:

```go
func HandleSingle(w http.ResponseWriter, r *http.Request) {
  toaster := goaster.NewToasterBuilder().WithAutoDismiss(false).Build()
  templ.Handler(pages.HomePage(toaster)).ServeHTTP(w, r)
}
```

### Displaying Toast Messages

In your _templ_ pages, call the appropriate method on the `Toaster` instance:

```templ
templ UserPage(toaster *goaster.Toaster) {
  @toaster.Success("Operation completed successfully.")
}
```

Each toast level has a corresponding method:

```templ
@toaster.Default("Sample message.")
@toaster.Error("An error occurred.")
@toaster.Info("Here's some information.")
@toaster.Warning("This is a warning message.")
```

> 💡 Toast messages are displayed when `@toaster.<Level>()` is called in your template.

#### Queueing Toast Messages from Go

You can queue multiple toast messages in your **handler**:

```go
toaster.PushDefault("Sample message.")
toaster.PushSuccess("Operation completed successfully.")
toaster.PushError("An error occurred.")
toaster.PushInfo("Here's some information.")
toaster.PushWarning("This is a warning message.")
```

Then render them all in your _templ_ page:

```templ
templ UserPage(toaster *goaster.Toaster) {
  @toaster.RenderAll()
}
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
