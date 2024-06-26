// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package goaster

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

// GoasterJSOptions represents options for the GoasterJS function.
type GoasterJSOptions struct {
	DismissTimer   int
	RemoveDuration int
}

func GoasterJS(toaster *Toaster, options *GoasterJSOptions) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_GoasterJS_3556`,
		Function: `function __templ_GoasterJS_3556(toaster, options){const { DismissTimer, RemoveDuration } = options ?? {};

    /**
     * @constant
     * @type {string}
     * @default "top"
     * @description Used to specify toast notifications originating from the top of the screen.
     */
    const DIRECTION_TOP = "top";

    /**
     * @constant
     * @type {string}
     * @default "bottom"
     * @description Used to specify toast notifications originating from the bottom of the screen.
     */
    const DIRECTION_BOTTOM = "bottom";

    /**
     * @constant
     * @type {number}
     * @default 3000
     * @description Default time in milliseconds before a toast notification is automatically dismissed.
     */
    const DISMISS_TIMER = DismissTimer ?? 3000;

    /**
     * @constant
     * @type {number}
     * @default 500
     * @description Duration in milliseconds for the exit animation of a toast notification before it is removed from the DOM.
     */
    const REMOVE_DURATION = RemoveDuration ?? 500;

    /**
     * Removes all CSS classes from the specified element that start with the given prefix.
     * @param {HTMLElement} element - The element from which to remove classes.
     * @param {string} prefix - The prefix of classes to be removed.
     */
    function removeClassesWithPrefix(element, prefix) {
        const classesToRemove = [];
        element.classList.forEach(className => {
            if (className.startsWith(prefix)) {
                classesToRemove.push(className);
            }
        });
        classesToRemove.forEach(className => {
            element.classList.remove(className);
        });
    }

    /**
     * Performs the exit animation for a toast element and removes it from the DOM after the animation.
     *
     * @param {HTMLElement} element - The toast element to animate and remove.
     * @param {boolean} animated - Specifies whether to apply an exit animation.
     * @param {string} positionName - The position of the toast used to determine the direction of the exit animation.
     */
    function performExitAnimationAndRemove(element, animated, positionName) {
        let classesToAdd = ["gttClose"];
        // Add animation-specific classes if animate is true
        if (animated) {
            classesToAdd.push(
                positionName.startsWith(DIRECTION_TOP)
                    ? "gttCloseFromTop"
                    : positionName.startsWith(DIRECTION_BOTTOM)
                    ? "gttCloseFromBottom"
                    : null
            );
        }
        element.classList.add(...classesToAdd.filter(Boolean));

        removeClassesWithPrefix(element, "gttShow");
        if(animated){
            removeClassesWithPrefix(element, "gttShowFrom");
        }

        setTimeout(() => element.remove(), REMOVE_DURATION);
    }

    /**
     * Closes the toast by stopping event propagation, clearing auto-dismiss timer, and performing exit animation.
     * @param {Event} e - The event object.
     */
    function closeToast(e) {
        e.stopPropagation();
        const toast = e.target.closest('[class^="gttToast"]');
        if (toast) {
            clearTimeout(parseInt(toast.dataset.dismissTimer));
            performExitAnimationAndRemove(toast, toaster.Animation, toaster.Position.Name, toaster.ProgressBar);
        }
    }

    /**
     * Initializes the automatic dismissal of toasts with an optional progress bar animation.
     * This function searches for all toasts marked for auto-dismissal and sets a timer for each.
     * If the progress bar feature is enabled, it also initializes and starts the progress bar animation.
     *
     * @param {boolean} isWithProgressBar - Indicates whether the progress bar animation should be enabled for auto-dismissing toasts.
     */
    let handleAutoDismiss = {
        init: function (isWithProgressBar) {
            const autoDismissToasts = document.querySelectorAll('[class^="gttToast"][data-auto-dismiss="true"]');
            autoDismissToasts.forEach(toast => {

                if(isWithProgressBar){
                    animateProgressBar(toast)
                }

                toast.dataset.dismissTimer = setTimeout(() => {
                    performExitAnimationAndRemove(toast, toaster.Animation, toaster.Position.Name);
                }, DISMISS_TIMER);
            });
        }
    }

    /**
     * Initializes and animates the progress bar within a toast notification element.
     * The progress bar visually indicates the remaining time until the toast is automatically dismissed.
     *
     * @param {HTMLElement} element - The toast notification element containing the progress bar.
     */
    function animateProgressBar(element) {
        const progressBarElement = element.querySelector('[class^="gttProgressBar"]');
        const duration = DISMISS_TIMER;

        // Initialize progress bar
        if (progressBarElement) {
            progressBarElement.style.transition = ` + "`" + `width ${duration}ms linear` + "`" + `;
            progressBarElement.style.width = '100%';
            // Start the animation by setting width to 0%
            setTimeout(() => { progressBarElement.style.width = '0%'; }, 0);
        }
    }

    document.body.addEventListener('click', function (e) {
        // Using closest to find the nearest ancestor which is a close button, starting from the event target itself.
        // The ` + "`" + `gttCloseBtn` + "`" + ` is a component of type ` + "`" + `templ.CSSClass` + "`" + `.
        // Its name is automatically generated, using the name of the component as a prefix.
        const closeButton = e.target.closest('[class*="gttCloseBtn"]');

        if (closeButton) {
            closeToast(e);
        }
    });

    document.addEventListener("DOMContentLoaded", function () {
        if(toaster) {
            handleAutoDismiss.init(toaster.ProgressBar);
        }
    });

}`,
		Call:       templ.SafeScript(`__templ_GoasterJS_3556`, toaster, options),
		CallInline: templ.SafeScriptInline(`__templ_GoasterJS_3556`, toaster, options),
	}
}
