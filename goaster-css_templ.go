// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.639
package goaster

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func GoasterCSS() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\n        *\n        {\n          box-sizing: border-box;\n          border-width: 0;\n          border-style: solid;\n        }\n\n        :root {\n            /* Padding Y and X for the toast element */\n            --gtt-py: 0rem;\n            --gtt-px: 0.75rem;\n\n            --gtt-bg: var(--gtt-default-bg);\n            --gtt-progress-bar-color: var(--gtt-default-progress-bar-color);\n\n            /* Font styling for the toast text */\n            --gtt-color: var(--gtt-default-color);\n            --gtt-font-family: inherit;\n            --gtt-font-size: 1rem;\n            --gtt-line-height: 1rem;\n\n            /* Border styling for the toast */\n            --gtt-border-width: 1px;\n            --gtt-border-style: solid;\n            --gtt-border-color: var(--gtt-default-border-color);\n            --gtt-border-radius: 0.375rem;\n\n            /* Accent toast theme colors */\n            --gtt-accent-light-bg: white;\n\n            /* Dark accent toast theme colors */\n            --gtt-accent-dark-bg: hsla(220.9,39.3%,11%, 0.85);\n\n            /* Default toast theme colors (border, background, text color) */\n            --gtt-default-border-color: #f3f4f6; /* gray-100 */\n            --gtt-default-bg: #f9fafb; /* gray-50 */\n            --gtt-default-color: #1f2937; /* gray-800 */\n            --gtt-default-progress-bar-color: #e5e7eb; /* gray-200 */\n            --gtt-default-accent-color: var(--gtt-default-color);\n            --gtt-default-accent-dark-color: #94a3b8; /* gray-400 */\n\n            /* Success toast theme colors (border, background, text color) */\n            --gtt-success-border-color: #dcfce7; /* green-100 */\n            --gtt-success-bg: #f0fdf4; /* green-50 */\n            --gtt-success-color: #166534; /* green-800 */\n            --gtt-success-progress-bar-color: #bbf7d0; /* green-200 */\n            --gtt-success-accent-color: var(--gtt-success-color);\n            --gtt-success-accent-dark-color: #22c55e; /* green-400 */\n\n            /* Error toast theme colors (border, background, text color) */\n            --gtt-error-border-color: #fee2e2; /* red-100 */\n            --gtt-error-bg: #fef2f2; /* red-50 */\n            --gtt-error-color: #991b1b; /* red-800 */\n            --gtt-error-progress-bar-color: #fecaca; /* red-200 */\n            --gtt-error-accent-color: var(--gtt-error-color);\n            --gtt-error-accent-dark-color: #f87171; /* red-400 */\n\n            /* Warning toast theme colors (border, background, text color) */\n            --gtt-warning-border-color: #ffedd5; /* orange-100 */\n            --gtt-warning-bg: #fff7ed; /* orange-50 */\n            --gtt-warning-color: #9a3412; /* orange-800 */\n            --gtt-warning-progress-bar-color: #fed7aa; /* orange-200 */\n            --gtt-warning-accent-color: var(--gtt-warning-color);\n            --gtt-warning-accent-dark-color: #fb923c; /* orange-400 */\n\n            /* Info toast theme colors (border, background, text color) */\n            --gtt-info-border-color: #dbeafe; /* blue-100 */\n            --gtt-info-bg: #eff6ff; /* blue-50 */\n            --gtt-info-color: #1e40af; /* blue-800 */\n            --gtt-info-progress-bar-color: #bfdbfe; /* blue-200 */\n            --gtt-info-accent-color: var(--gtt-info-color);\n            --gtt-info-accent-dark-color: #60a5fa; /* blue-400 */\n\n            /* Default entrance animation properties */\n            --gtt-animation-entrance-duration: 0.5s;\n            --gtt-animation-entrance-direction: normal;\n            --gtt-animation-entrance-timing-function: ease;\n            --gtt-animation-entrance-delay: 0s;\n\n            /* Default exit animation properties */\n            --gtt-animation-exit-duration: 0.5s;\n            --gtt-animation-exit-direction: normal;\n            --gtt-animation-exit-timing-function: ease;\n            --gtt-animation-exit-delay: 0s;\n\n            /* Animation properties (names) when entrance direction is from the top */\n            --gtt-animation-entrance-name-top: gttFadeInDown;\n            --gtt-animation-exit-name-top: gttFadeOutUp;\n\n            /* Animation properties (names) when entrance direction is from the bottom */\n            --gtt-animation-entrance-name-bottom: gttFadeInUp;\n            --gtt-animation-exit-name-bottom: gttFadeOutDown;\n        }\n\n        button[class^=\"gttCloseBtn\"]:hover,\n        button[class^=\"gttCloseBtn\"]:focus {\n            background-color: var(--gtt-border-color);\n        }\n\n        /* Define the animation properties for a toast element's exit animation */\n        .gttClose {\n            animation-duration: var(--gtt-animation-exit-duration);\n            animation-direction: var(--gtt-animation-exit-direction);\n            animation-timing-function: var(--gtt-animation-exit-timing-function);\n            animation-delay: var(--gtt-animation-exit-delay);\n        }\n\n        .gttCloseFromTop {\n            animation-name: var(--gtt-animation-exit-name-top);\n        }\n\n        .gttCloseFromBottom {\n            animation-name: var(--gtt-animation-exit-name-bottom);\n        }\n\n        /* Animations */\n        @keyframes gttFadeInUp {\n            from {\n                opacity: 0;\n                transform: translate3d(0, 100%, 0);\n            }\n\n            to {\n                opacity: 1;\n                transform: translate3d(0, 0, 0);\n            }\n        }\n\n        @keyframes gttFadeInDown {\n            from {\n                opacity: 0;\n                transform: translate3d(0, -100%, 0);\n            }\n\n            to {\n                opacity: 1;\n                transform: translate3d(0, 0, 0);\n            }\n        }\n\n        @keyframes gttFadeOutUp {\n            from {\n                opacity: 1;\n            }\n\n            to {\n                opacity: 0;\n                transform: translate3d(0, -100%, 0);\n            }\n        }\n\n        @keyframes gttFadeOutDown {\n            from {\n                opacity: 1;\n            }\n\n            to {\n                opacity: 0;\n                transform: translate3d(0, 100%, 0);\n            }\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func gttContainer() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`position:fixed;`)
	templ_7745c5c3_CSSBuilder.WriteString(`z-index:1000;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column-reverse;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:flex-end;`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-height:90%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`overflow-y:auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:var(--gtt-font-family);`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:var(--gtt-font-size);`)
	templ_7745c5c3_CSSBuilder.WriteString(`line-height:var(--gtt-line-height);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainer`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttContainerBottomRight() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`right:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`bottom:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:flex-end;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainerBottomRight`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttContainerBottomLeft() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`left:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`bottom:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:flex-start;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainerBottomLeft`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttContainerBottomCenter() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`left:50%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`bottom:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`transform:translateX(-50%);`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainerBottomCenter`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttContainerTopRight() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`right:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`top:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:flex-end;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainerTopRight`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttContainerTopLeft() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`left:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`top:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:flex-start;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainerTopLeft`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttContainerTopCenter() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`left:50%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`top:15px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`transform:translateX(-50%);`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttContainerTopCenter`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttToast() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`position:relative;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-width:250px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin-top:0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--gtt-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:var(--gtt-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-color:var(--gtt-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:var(--gtt-py) var(--gtt-px);`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttToast`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttRounded() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:var(--gtt-border-radius);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttRounded`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttBordered() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-style:var(--gtt-border-style);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-width:var(--gtt-border-width);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttBordered`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccent() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-width:4px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-style:solid;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccent`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAnimated() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`animation-duration:var(--gtt-animation-entrance-duration);`)
	templ_7745c5c3_CSSBuilder.WriteString(`animation-direction:var(--gtt-animation-entrance-direction);`)
	templ_7745c5c3_CSSBuilder.WriteString(`animation-timing-function:var(--gtt-animation-entrance-timing-function);`)
	templ_7745c5c3_CSSBuilder.WriteString(`animation-delay:var(--gtt-animation-entrance-delay);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAnimated`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttShow() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`visibility:visible;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttShow`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttShowFromTop() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`animation-name:var(--gtt-animation-entrance-name-top);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttShowFromTop`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttShowFromBottom() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`animation-name:var(--gtt-animation-entrance-name-bottom);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttShowFromBottom`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttProgressBar() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`position:absolute;`)
	templ_7745c5c3_CSSBuilder.WriteString(`overflow:hidden;`)
	templ_7745c5c3_CSSBuilder.WriteString(`top:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`left:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:4px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:var(--gtt-progress-bar-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`transition:width linear;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttProgressBar`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttDefaultLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-default-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-default-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-default-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-progress-bar-color:var(--gtt-default-progress-bar-color);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttDefaultLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentDefaultLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-default-accent-border-color, #6b7280);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentDefaultLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentLightDefaultLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-default-accent-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-light-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-default-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-default-accent-border-color, #6b7280);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentLightDefaultLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentDarkDefaultLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-default-accent-dark-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-dark-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-default-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-success-accent-border-color, #6b7280);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentDarkDefaultLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttSuccessLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-success-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-success-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-success-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-progress-bar-color:var(--gtt-success-progress-bar-color);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttSuccessLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentSuccessLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-success-accent-dark-color, #22c55e);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentSuccessLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentLightSuccessLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-success-accent-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-light-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-success-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-success-accent-dark-color, #22c55e);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentLightSuccessLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentDarkSuccessLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-success-accent-dark-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-dark-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-success-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-success-accent-border-color, #22c55e);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentDarkSuccessLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttErrorLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-error-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-error-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-error-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-progress-bar-color:var(--gtt-error-progress-bar-color);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttErrorLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentErrorLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-error-accent-border-color, #ef4444);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentErrorLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentLightErrorLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-error-accent-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-light-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-error-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-error-accent-border-color, #ef4444);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentLightErrorLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentDarkErrorLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-error-accent-dark-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-dark-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-error-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-error-accent-border-color, #ef4444);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentDarkErrorLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttWarningLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-warning-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-warning-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-warning-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-progress-bar-color:var(--gtt-warning-progress-bar-color);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttWarningLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentWarningLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-warning-accent-border-color, #f97316);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentWarningLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentLightWarningLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-warning-accent-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-light-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-warning-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-warning-accent-border-color, #f97316);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentLightWarningLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentDarkWarningLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-warning-accent-dark-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-dark-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-warning-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-warning-accent-border-color, #f97316);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentDarkWarningLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttInfoLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-info-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-info-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-info-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-progress-bar-color:var(--gtt-info-progress-bar-color);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttInfoLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentInfoLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-info-accent-border-color, #3b82f6);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentInfoLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentLightInfoLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-info-accent-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-light-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-info-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-info-accent-border-color, #3b82f6);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentLightInfoLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttAccentDarkInfoLevel() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-color:var(--gtt-info-accent-dark-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-bg:var(--gtt-accent-dark-bg);`)
	templ_7745c5c3_CSSBuilder.WriteString(`--gtt-border-color:var(--gtt-info-border-color);`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-left-color:var(--gtt-info-accent-border-color, #3b82f6);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttAccentDarkInfoLevel`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttIcon() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`display:inline-flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-shrink:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:1.125rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:1.125rem;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttIcon`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttMessage() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`margin:1em 0 1em 0.75em;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttMessage`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttCloseBtn() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`margin-left:auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:inline-flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:transparent;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:none;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:1e5px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`aspect-ratio:1;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:inherit;`)
	templ_7745c5c3_CSSBuilder.WriteString(`cursor:pointer;`)
	templ_7745c5c3_CSSBuilder.WriteString(`transition:background 0.3s cubic-bezier(0.4, 0, 0.2, 1);`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttCloseBtn`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func gttSrOnly() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`position:absolute;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:1px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:1px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:-1px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`overflow:hidden;`)
	templ_7745c5c3_CSSBuilder.WriteString(`clip:rect(0, 0, 0, 0);`)
	templ_7745c5c3_CSSBuilder.WriteString(`white-space:nowrap;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-width:0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`gttSrOnly`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}
