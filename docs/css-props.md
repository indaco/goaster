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

### Variants

[Source Code](../assets/css/variants/).

All variants (except `accent-light` and `accent-dark`) automatically adapt to light/dark mode via `@media (prefers-color-scheme)`.

| Variant        | Light Mode                               | Dark Mode                                       | Key Visual                 |
| :------------- | :--------------------------------------- | :---------------------------------------------- | :------------------------- |
| `accent`       | White bg, dark text, colored left border | Dark bg, light text, colored left border        | Auto-adaptive left accent  |
| `accent-light` | White bg, dark text, colored left border | - (forced light)                                | Explicit light mode        |
| `accent-dark`  | - (forced dark)                          | Dark bg, light text, colored left border        | Explicit dark mode         |
| `filled`       | Base-color bg, white text                | Dark-color bg, lightest text                    | Bold solid background      |
| `outlined`     | Transparent bg, base border, dark text   | Transparent bg, light border, light text        | Full colored border        |
| `soft`         | Lightest bg, dark text, no border        | Subtle tinted dark bg, lightest text, no border | Subtle tinted background   |
| `minimal`      | Transparent bg, dark text, no border     | Transparent bg, light text, no border           | Just colored text          |
| `brutalist`    | White bg, thick border, offset shadow    | Near-black bg, thick border, offset shadow      | Raw, heavy, blocky         |
| `retro`        | Cream bg, warm tones, inset shadow       | Warm dark bg, warm tones                        | Vintage, nostalgic         |
| `neon`         | Falls back to default                    | Dark bg, neon glow borders                      | Cyberpunk glow (dark only) |

#### Accent

Background tokens used by the accent family of variants.

| Variant        | CSS Variable            | Light Default | Dark Default                    |
| :------------- | :---------------------- | :------------ | :------------------------------ |
| `accent`       | `--gtt-accent-bg`       | `white`       | `hsla(220.9, 39.3%, 11%, 0.85)` |
| `accent-light` | `--gtt-accent-light-bg` | `white`       | -                               |
| `accent-dark`  | `--gtt-accent-dark-bg`  | -             | `hsla(220.9, 39.3%, 11%, 0.85)` |

Per-level accent color overrides (apply to `accent`, `accent-light`, `accent-dark`):

- `--gtt-<level>-accent-color` - text color
- `--gtt-<level>-accent-border-color` - left border color

#### Filled

Per-level overrides:

- `--gtt-filled-<level>-bg` - background color
- `--gtt-filled-<level>-color` - text color

#### Outlined

Per-level overrides:

- `--gtt-outlined-<level>-color` - text color
- `--gtt-outlined-<level>-border` - border color
- `--gtt-outlined-bg` - shared background (default: `transparent`)

#### Soft

Per-level overrides:

- `--gtt-soft-<level>-bg` - background color (light: lightest tint; dark: 15% base tint)
- `--gtt-soft-<level>-color` - text color

#### Minimal

Per-level overrides:

- `--gtt-minimal-<level>-color` - text color

Background and border are always `transparent`.

#### Brutalist

Per-level overrides:

- `--gtt-brutalist-<level>-color` - text color
- `--gtt-brutalist-<level>-border` - border color
- `--gtt-brutalist-bg` - shared background (light: `white`; dark: `hsl(0 0% 8%)`)

Structural: `border-width: 3px`, `border-radius: 0`, `box-shadow: 4px 4px 0 0` with level base color.

#### Retro

Per-level overrides:

- `--gtt-retro-<level>-color` - text color
- `--gtt-retro-<level>-border` - border color
- `--gtt-retro-bg` - shared background (light: `hsl(45 40% 96%)`; dark: `hsl(35 25% 18%)`)

Structural: `border-width: 1px`, `box-shadow: inset 0 1px 2px` for paper-like depth.

#### Neon (dark mode only)

All styles are wrapped in `@media (prefers-color-scheme: dark)`. In light mode, toasts fall back to default appearance.

Uses level base colors directly - no custom override properties. Structural: `border-width: 1px`, dual `box-shadow` glow via `color-mix()` at 40% and 15% opacity.

## Theming

To customize any level (e.g., change the _success_ color), override the base variable:

```css
:root {
  --gtt-color-success-base: oklch(
    0.72 0.25 150
  ); /* your custom success color */
}
```

All derived shades for `--gtt-color-success-*` will automatically update based on your base value.

## Component Styles

[Source Code](../assets/css/styles.css)

These custom properties control spacing, layout, typography, and animation behavior for the toast component.

### Layout

| Property              | Default                        | Description                     |
| :-------------------- | :----------------------------- | :------------------------------ |
| `--gtt-py`            | `0`                            | Vertical padding for the toast. |
| `--gtt-px`            | `0.75em`                       | Horizontal padding.             |
| `--gtt-bg`            | `--gtt-color-<level>-lightest` | Background color.               |
| `--gtt-border-color`  | `--gtt-color-<level>-lighter`  | Border color.                   |
| `--gtt-border-width`  | `1px`                          | Border width.                   |
| `--gtt-border-style`  | `solid`                        | Border style.                   |
| `--gtt-border-radius` | `0.375rem`                     | Border radius.                  |

### Typography

| Property            | Default                    | Description  |
| :------------------ | :------------------------- | :----------- |
| `--gtt-color`       | `--gtt-color-<level>-dark` | Text color.  |
| `--gtt-font`        | `inherit`                  | Font family. |
| `--gtt-font-size`   | `1rem`                     | Font size.   |
| `--gtt-line-height` | `1rem`                     | Line height. |

### Progress Bar

| Property               | Default                     | Description            |
| :--------------------- | :-------------------------- | :--------------------- |
| `--gtt-progress-color` | `--gtt-color-<level>-light` | Color of progress bar. |

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
