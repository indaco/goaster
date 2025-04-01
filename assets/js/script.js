/** Toast direction constants */
const DIRECTION_TOP = 'top'
const DIRECTION_BOTTOM = 'bottom'

/**
 * @typedef {Object} GoasterJSConfig
 * @property {number} [dismissTimer] - Time (in ms) before the toast auto-dismisses.
 * @property {number} [removeDuration] - Duration (in ms) for the removal animation.
 */

/**
 * Initializes the Goaster system with optional configuration for toast timing.
 *
 * @param {GoasterJSConfig} [config={}] - Configuration options.
 * @returns {{ dismissTimer: number, removeDuration: number }} The resolved configuration.
 */
function initGoaster({ dismissTimer = 3000, removeDuration = 500 } = {}) {
  return { dismissTimer, removeDuration }
}

/**
 * Removes known "show" classes from the toast element.
 * @param {HTMLElement} element - The toast element.
 */
function removeShowClasses(element) {
  const showClasses = ['gttShow', 'gttShowFromTop', 'gttShowFromBottom']
  element.classList.remove(...showClasses)
}

/**
 * Performs the exit animation for a toast element and removes it from the DOM after the animation.
 *
 * @param {HTMLElement} element - The toast element to animate and remove.
 * @param {boolean} animated - Whether to apply an exit animation.
 * @param {string} positionName - The position of the toast used to determine animation direction.
 * @param {number} duration - The duration before removal (in ms).
 */
function performExitAnimationAndRemove(
  element,
  animated,
  positionName,
  duration = 500
) {
  const classesToAdd = ['gttClose']

  if (animated) {
    classesToAdd.push(
      positionName.startsWith(DIRECTION_TOP)
        ? 'gttCloseFromTop'
        : 'gttCloseFromBottom'
    )
  }

  element.classList.add(...classesToAdd)

  removeShowClasses(element)
  setTimeout(() => element.remove(), duration)
}

/**
 * Initializes and animates the progress bar within a toast notification element.
 *
 * @param {HTMLElement} element - The toast notification element containing the progress bar.
 * @param {number} duration - Duration in milliseconds for the progress bar animation.
 */
function animateProgressBar(element, duration = 3000) {
  const progressBarElement = element.querySelector('.gttProgressBar')

  if (progressBarElement) {
    progressBarElement.style.transition = `width ${duration}ms linear`
    progressBarElement.style.width = '100%'

    requestAnimationFrame(() => {
      progressBarElement.style.width = '0%'
    })
  }
}

/**
 * Initializes the automatic dismissal of toasts with an optional progress bar animation.
 *
 * @param {boolean} isWithProgressBar - Whether to animate the progress bar.
 * @param {number} dismissTimer - Duration before auto-dismiss (in ms).
 * @param {number} removeDuration - Duration for removal animation (in ms).
 */
const handleAutoDismiss = {
  init(isWithProgressBar, dismissTimer, removeDuration) {
    document
      .querySelectorAll('.gttToast[role="alert"][data-auto-dismiss="true"]')
      .forEach((toast) => {
        if (isWithProgressBar) animateProgressBar(toast, dismissTimer)

        toast.dataset.dismissTimer = setTimeout(() => {
          performExitAnimationAndRemove(
            toast,
            toast.dataset.animation === 'true',
            toast.dataset.position,
            removeDuration
          )
        }, dismissTimer)
      })
  }
}

/**
 * Closes the toast by stopping propagation, clearing timers, and performing exit animation.
 * @param {Event} e - The event object.
 */
function closeToast(e) {
  e.stopPropagation()

  const toast = e.target.closest('.gttToast[role="alert"]')
  if (toast) {
    clearTimeout(Number(toast.dataset.dismissTimer))
    performExitAnimationAndRemove(
      toast,
      toast.dataset.animation === 'true',
      toast.dataset.position
    )
  }
}

/** Close toast on button click */
document.body.addEventListener('click', (e) => {
  const closeButton = e.target.closest('.gttCloseBtn')
  if (closeButton) closeToast(e)
})

/** Initialize auto-dismiss on page load */
document.addEventListener('DOMContentLoaded', () => {
  document
    .querySelectorAll('.gttToast[data-auto-dismiss="true"]')
    .forEach((el) => {
      const { dismissTimer, removeDuration } = initGoaster({
        dismissTimer: Number(el.dataset.dismissTimer) || undefined,
        removeDuration: Number(el.dataset.removeDuration) || undefined
      })

      const hasProgressBar = !!el.querySelector('.gttProgressBar')
      handleAutoDismiss.init(hasProgressBar, dismissTimer, removeDuration)
    })
})
