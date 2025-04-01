// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package variants

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

var toasterAccentLightVariantHandle = templ.NewOnceHandle()

func variantAccentLight() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style type=\"text/css\">\n/* [tempo] BEGIN - Do not edit! This section is auto-generated. */\n:root{--gtt-accent-light-bg: white;--gtt-default-accent-color: var(--gtt-color-default-dark);--gtt-success-accent-color: var(--gtt-color-success-dark);--gtt-error-accent-color: var(--gtt-color-error-dark);--gtt-warning-accent-color: var(--gtt-color-warning-dark);--gtt-info-accent-color: var(--gtt-color-info-dark)}.gttToast[role=alert][data-variant=accent-light]{border:0;border-left-width:4px;border-left-style:solid}.gttToast[role=alert][data-level=default][data-variant=accent-light]{--_gtt-color: var(--gtt-default-accent-color);--_gtt-bg: var(--gtt-accent-light-bg);--_gtt-border-color: var(--gtt-default-border-color);border-left-color:var( --gtt-default-accent-border-color, var(--gtt-color-default-base) )}.gttToast[role=alert][data-level=success][data-variant=accent-light]{--_gtt-color: var(--gtt-success-accent-color);--_gtt-bg: var(--gtt-accent-light-bg);--_gtt-border-color: var(--gtt-success-border-color);border-left-color:var( --gtt-success-accent-dark-color, var(--gtt-color-success-base) )}.gttToast[role=alert][data-level=error][data-variant=accent-light]{--_gtt-color: var(--gtt-error-accent-color);--_gtt-bg: var(--gtt-accent-light-bg);--_gtt-border-color: var(--gtt-error-border-color);border-left-color:var( --gtt-error-accent-border-color, var(--gtt-color-error-base) )}.gttToast[role=alert][data-level=warning][data-variant=accent-light]{--_gtt-color: var(--gtt-warning-accent-color);--_gtt-bg: var(--gtt-accent-light-bg);--_gtt-border-color: var(--gtt-warning-border-color);border-left-color:var( --gtt-warning-accent-border-color, var(--gtt-color-warning-base) )}.gttToast[role=alert][data-level=info][data-variant=accent-light]{--_gtt-color: var(--gtt-info-accent-color);--_gtt-bg: var(--gtt-accent-light-bg);--_gtt-border-color: var(--gtt-info-border-color);border-left-color:var( --gtt-info-accent-border-color, var(--gtt-color-info-base) )}\n\n/* [tempo] END */\n</style>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = toasterAccentLightVariantHandle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
