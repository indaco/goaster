<h1 align="center" style="font-size: 2.5rem;">
  goaster
</h1>
<h2 align="center" style="font-size: 1.5rem;">
A configurable, themeable and non-intrusive server-rendered toast notification component for Go web applications.
</h2>
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
</p>

Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

### Features

- **No External Dependencies**: Built with native Go and the `templ` library, requiring no external frontend dependencies.
- **Multiple Toasts**: Support to display multiple toast notifications.
- **Configurable**: Customize appearance, behavior, and position.
- **Themeable**: Use CSS variables to theme your toasts to match your application's design.
- **Icon Support**: Include default SVG icons for various toast levels (such as success, error, info, etc.), allowing you to use your preferred icons.
- **Positioning**: Flexible positioning of toast messages (top-right, bottom-left, etc.).
- **Progress Bar for Auto-Dismiss**: A visual progress bar indicates the countdown until the toast will automatically dismiss (with auto-dismiss enabled only).
- **Animations**: Entrance and exit animations for toast notifications.
  - _Default Animations_: built-in animations for toasts appearing and disappearing.
  - _Custom Animations_: Customize animations via CSS variables or integrate with external CSS animation libraries.
  - _Animation Control_: Full control over animation timing, easing, delay and effects, via CSS variables.
  - _Disable Animations Option_: Ability to completely disable animations, providing flexibility for different application needs and user preferences.

<div style="display: flex; justify-content: center;">
   <img src="statics/demo.gif" alt="Image" >
</div>

## Installation

Ensure your project is using Go Modules.

To install the module, use the `go get` command:

```sh
go get github.com/indaco/goaster
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

### Add goester CSS and Javascript

`goaster` leverages the `templ` library's features, including CSS Components and JavaScript Templates, to encapsulate all necessary styling and functionality without relying on external dependencies.

- `GoasterCSS`: it supplies the required CSS, encapsulating the visual design and layout specifics of the toast notifications.
- `GoasterJS`: it provides the JavaScript logic essential for dynamic behaviors such as displaying, hiding, and managing toast notifications.

To facilitate integration with Go's `template/html` standard library, `goaster` includes a dedicated `HTMLGenerator` type to seamlessly integrate toast notifications into web applications built with Go's `html/template` standard library. 

3 methods, acting as wrappers to the templ's `templ.ToGoHTML`, generate the necessary HTML to be embedded them into server-rendered pages:

- `GoasterCSSToGoHTML`: render the `GoasterCSS` component into a `template.HTML` value.
- `GoasterJSToGoHTML`: render the `GoasterJS`component into a `template.HTML` value.

> **Note**: refer to the [Examples](#examples) section to see how to use `goaster` with `templ` and `html/template`.

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

## Theming

Toast notifications are themeable using CSS variables (prefix `gtt`) to customize the appearance according to your design.

Here below is the list of all CSS variables defined and their default values:

```css
:root {
  /* Padding Y and X for the toast element */
  --gtt-py: 0;
  --gtt-px: 0.75rem;

  /* Font styling for the toast text */
  --gtt-font-family: inherit;
  --gtt-font-size: 1rem;
  --gtt-line-height: 1rem;

  /* Border styling for the toast */
  --gtt-border-radius: 0.375rem;
  --gtt-border-style: solid;
  --gtt-border-width: 1px;

  /* Default toast theme colors (border, background, text color) */
  --gtt-default-border-color: #f3f4f6; /* gray-100 */
  --gtt-default-bg: #f9fafb; /* gray-50 */
  --gtt-default-color: #1f2937; /* gray-800 */

  /* Success toast theme colors (border, background, text color) */
  --gtt-success-border-color: #dcfce7; /* green-100 */
  --gtt-success-bg: #f0fdf4; /* green-50 */
  --gtt-success-color: #166534; /* green-800 */

  /* Error toast theme colors (border, background, text color) */
  --gtt-error-border-color: #fee2e2; /* red-100 */
  --gtt-error-bg: #fef2f2; /* red-50 */
  --gtt-error-color: #991b1b; /* red-800 */

  /* Warning toast theme colors (border, background, text color) */
  --gtt-warning-border-color: #ffedd5; /* orange-100 */
  --gtt-warning-bg: #fff7ed; /* orange-50 */
  --gtt-warning-color: #9a3412; /* orange-800 */

  /* Info toast theme colors (border, background, text color) */
  --gtt-info-border-color: #dbeafe; /* blue-100 */
  --gtt-info-bg: #eff6ff; /* blue-50 */
  --gtt-info-color: #1e40af; /* blue-800 */

  /* Default entrance animation properties */
  --gtt-animation-entrance-duration: 0.5s;
  --gtt-animation-entrance-direction: normal;
  --gtt-animation-entrance-timing-function: ease;
  --gtt-animation-entrance-delay: 0s;

  /* Default exit animation properties */
  --gtt-animation-exit-duration: 0.5s;
  --gtt-animation-exit-direction: normal;
  --gtt-animation-exit-timing-function: ease;
  --gtt-animation-exit-delay: 0s;

  /* Animation properties (names) when entrance direction is from the top */
  --gtt-animation-entrance-name-top: gttFadeInDown;
  --gtt-animation-exit-name-top: gttFadeOutUp;

  /* Animation properties (names) when entrance direction is from the bottom */
  --gtt-animation-entrance-name-bottom: gttFadeInUp;
  --gtt-animation-exit-name-bottom: gttFadeOutDown;
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

## Examples

- [use with `a-h/templ` template](_examples/a-h-templ-single-toast)
- [multiple messages with `a-h/templ`](_examples/a-h-templ-multiple-toasts)
- [use with `template/html`](_examples/go-html-template-single-toast)
- [multiple  messages with `template/html`](_examples/go-html-template-multiple-toasts)
- [theming](_examples/theming)
- [custom icons](_examples/custom-icons)
- [custom animations](_examples/custom-animations)

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
