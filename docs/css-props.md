# Goaster CSS custom props

## Semantic Color Levels

[Source Code](../assets/css/colors.css).

All semantic color levels use [`color-mix()`](<https://developer.mozilla.org/en-US/docs/Web/CSS/color-mix()>) in the `oklab` color space, defined by:

```css
--gtt-color-space: oklab;
```

| Level   | Base Value                   |
| :------ | :--------------------------- |
| default | `oklch(0.554 0.046 257.417)` |
| success | `oklch(0.723 0.219 149.579)` |
| error   | `oklch(0.637 0.237 25.331)`  |
| warning | `oklch(0.769 0.188 70.08)`   |
| info    | `oklch(0.623 0.214 259.815)` |

The following variables are generated from the base color:

- `--gtt-color-<level>-lightest`: 10% white
- `--gtt-color-<level>-lighter`: 30% white
- `--gtt-color-<level>-light`: 70% white
- `--gtt-color-<level>-dark`: 70% black
- `--gtt-color-<level>-darker`: 30% black
- `--gtt-color-<level>-darkest`: 10% black

### Accent Variants

These tokens are used by the accent variants and override toast background colors.

| Variant       | CSS Variable           | Value                           |
| :------------ | :--------------------- | :------------------------------ |
| `accent`      | `--gtt-accent-bg`      | `white`                         |
| `accent-dark` | `--gtt-accent-dark-bg` | `hsla(220.9, 39.3%, 11%, 0.85)` |

#### Semantic Accent Color Overrides

Each level can override its accent color using the following properties:

> Note: these are used when applying `variant="accent"` or `variant="accent-dark"`.

## Theming

To customize any level (e.g., change the _success_ color), override the base variable:

```css
:root {
  --gtt-color-success-base: oklch(0.72 0.25 150); /* your custom success color */
}
```

All derived shades for `--gtt-color-success-*` will automatically update based on your base value.

## Component Styles

[Source Code](../assets/css/styles.css)

These custom properties control spacing, layout, typography, and animation behavior for the toast component.

### Layout

| Property              | Default                           | Description                     |
| :-------------------- | :-------------------------------- | :------------------------------ |
| `--gtt-py`            | `0`                               | Vertical padding for the toast. |
| `--gtt-px`            | `0.75em`                          | Horizontal padding.             |
| `--gtt-bg`            | `var(--gtt-default-bg)`           | Background color.               |
| `--gtt-border-color`  | `var(--gtt-default-border-color)` | Border color.                   |
| `--gtt-border-width`  | `1px`                             | Border width.                   |
| `--gtt-border-style`  | `solid`                           | Border style.                   |
| `--gtt-border-radius` | `0.375rem`                        | Border radius.                  |

### Typography

| Property            | Default                    | Description  |
| :------------------ | :------------------------- | :----------- |
| `--gtt-color`       | `var(--gtt-default-color)` | Text color.  |
| `--gtt-font`        | `inherit`                  | Font family. |
| `--gtt-font-size`   | `1rem`                     | Font size.   |
| `--gtt-line-height` | `1rem`                     | Line height. |

### Progress Bar

| Property               | Default                             | Description            |
| :--------------------- | :---------------------------------- | :--------------------- |
| `--gtt-progress-color` | `var(--gtt-default-progress-color)` | Color of progress bar. |

### Animation

| Property                     | Default          | Description                                  |
| :--------------------------- | :--------------- | :------------------------------------------- |
| `--gtt-anim-name-in-top`     | `gttFadeInDown`  | Animation name for entering from the top.    |
| `--gtt-anim-name-in-bottom`  | `gttFadeInUp`    | Animation name for entering from the bottom. |
| `--gtt-anim-name-out-top`    | `gttFadeOutUp`   | Animation name for exiting to the top.       |
| `--gtt-anim-name-out-bottom` | `gttFadeOutDown` | Animation name for exiting to the bottom.    |
| `--gtt-anim-duration-in`     | `0.5s`           | Duration of entrance animation.              |
| `--gtt-anim-delay-in`        | `0s`             | Delay before entrance starts.                |
| `--gtt-anim-direction-in`    | `normal`         | Direction of entrance animation.             |
| `--gtt-anim-timing-in`       | `ease`           | Timing function for entrance.                |
| `--gtt-anim-duration-out`    | `0.5s`           | Duration of exit animation.                  |
| `--gtt-anim-delay-out`       | `0s`             | Delay before exit starts.                    |
| `--gtt-anim-direction-out`   | `normal`         | Direction of exit animation.                 |
| `--gtt-anim-timing-out`      | `ease`           | Timing function for exit.                    |
