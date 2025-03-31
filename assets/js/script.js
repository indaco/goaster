const { DismissTimer = 3000, RemoveDuration = 500 } = options ?? {}

/** Toast direction constants */
const DIRECTION_TOP = 'top'
const DIRECTION_BOTTOM = 'bottom'

/** Auto-dismiss and removal durations */
const DISMISS_TIMER = DismissTimer
const REMOVE_DURATION = RemoveDuration

/**
 * Removes all CSS classes from the specified element that start with the given prefix.
 * @param {HTMLElement} element - The element from which to remove classes.
 * @param {string} prefix - The prefix of classes to be removed.
 */
function removeClassesWithPrefix(element, prefix) {
  element.classList.forEach((className) => {
    if (className.startsWith(prefix)) {
      element.classList.remove(className)
    }
  })
}

/**
 * Performs the exit animation for a toast element and removes it from the DOM after the animation.
 *
 * @param {HTMLElement} element - The toast element to animate and remove.
 * @param {boolean} animated - Specifies whether to apply an exit animation.
 * @param {string} positionName - The position of the toast used to determine the direction of the exit animation.
 */
function performExitAnimationAndRemove(element, animated, positionName) {
  const classesToAdd = ['gttClose']

  if (animated) {
    classesToAdd.push(
      positionName.startsWith(DIRECTION_TOP)
        ? 'gttCloseFromTop'
        : 'gttCloseFromBottom'
    )
  }
  element.classList.add(...classesToAdd)

  removeClassesWithPrefix(element, 'gttShow')
  if (animated) removeClassesWithPrefix(element, 'gttShowFrom')

  setTimeout(() => element.remove(), REMOVE_DURATION)
}

/**
 * Closes the toast by stopping event propagation, clearing auto-dismiss timer, and performing exit animation.
 * @param {Event} e - The event object.
 */
function closeToast(e) {
  e.stopPropagation()
  const toast = e.target.closest('[class^="gttToast"]')
  if (toast) {
    clearTimeout(Number(toast.dataset.dismissTimer))
    performExitAnimationAndRemove(
      toast,
      toaster.Animation,
      toaster.Position.Name
    )
  }
}

/**
 * Initializes the automatic dismissal of toasts with an optional progress bar animation.
 */
const handleAutoDismiss = {
  init(isWithProgressBar) {
    document
      .querySelectorAll('[class^="gttToast"][data-auto-dismiss="true"]')
      .forEach((toast) => {
        if (isWithProgressBar) animateProgressBar(toast)

        toast.dataset.dismissTimer = setTimeout(() => {
          performExitAnimationAndRemove(
            toast,
            toaster.Animation,
            toaster.Position.Name
          )
        }, DISMISS_TIMER)
      })
  }
}

/**
 * Initializes and animates the progress bar within a toast notification element.
 *
 * @param {HTMLElement} element - The toast notification element containing the progress bar.
 */
function animateProgressBar(element) {
  const progressBarElement = element.querySelector('[class^="gttProgressBar"]')

  if (progressBarElement) {
    progressBarElement.style.transition = `width ${DISMISS_TIMER}ms linear`
    progressBarElement.style.width = '100%'

    requestAnimationFrame(() => {
      progressBarElement.style.width = '0%'
    })
  }
}

/** Close toast on button click */
document.body.addEventListener('click', (e) => {
  const closeButton = e.target.closest('[class*="gttCloseBtn"]')
  if (closeButton) closeToast(e)
})

/** Initialize auto-dismiss on page load */
document.addEventListener('DOMContentLoaded', () => {
  if (toaster) handleAutoDismiss.init(toaster.ProgressBar)
})
