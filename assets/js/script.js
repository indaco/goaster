(function () {
  /** Toast direction constants */
  const DIRECTION_TOP = "top";
  const DIRECTION_BOTTOM = "bottom";

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
    return { dismissTimer, removeDuration };
  }

  /**
   * Removes known "show" classes from the toast element.
   * @param {HTMLElement} element - The toast element.
   */
  function removeShowClasses(element) {
    const showClasses = ["gttShow", "gttShowFromTop", "gttShowFromBottom"];
    element.classList.remove(...showClasses);
  }

  /**
   * Performs the exit animation for a toast element and removes it from the DOM
   * after the animation ends. Uses the `animationend` event when animation is
   * applied, with a fallback timeout in case the event never fires.
   *
   * @param {HTMLElement} element - The toast element to animate and remove.
   * @param {boolean} animated - Whether to apply an exit animation.
   * @param {string} positionName - The position of the toast used to determine animation direction.
   * @param {number} duration - The fallback duration before removal (in ms).
   */
  function performExitAnimationAndRemove(
    element,
    animated,
    positionName,
    duration = 500,
  ) {
    const classesToAdd = ["gttClose"];

    if (animated) {
      classesToAdd.push(
        positionName.startsWith(DIRECTION_TOP)
          ? "gttCloseFromTop"
          : "gttCloseFromBottom",
      );
    }

    element.classList.add(...classesToAdd);
    removeShowClasses(element);

    if (animated) {
      var removed = false;
      var fallbackTimer = setTimeout(function () {
        if (!removed) {
          removed = true;
          element.remove();
        }
      }, duration);

      element.addEventListener(
        "animationend",
        function () {
          if (!removed) {
            removed = true;
            clearTimeout(fallbackTimer);
            element.remove();
          }
        },
        { once: true },
      );
    } else {
      element.remove();
    }
  }

  /**
   * Initializes and animates the progress bar within a toast notification element.
   *
   * @param {HTMLElement} element - The toast notification element containing the progress bar.
   * @param {number} duration - Duration in milliseconds for the progress bar animation.
   */
  function animateProgressBar(element, duration = 3000) {
    const progressBarElement = element.querySelector(".gttProgressBar");

    if (progressBarElement) {
      progressBarElement.style.transition = "width " + duration + "ms linear";
      progressBarElement.style.width = "100%";

      requestAnimationFrame(function () {
        progressBarElement.style.width = "0%";
      });
    }
  }

  /**
   * Sets up automatic dismissal for a single toast element with an optional
   * progress bar animation. Stores the timer ID on `dataset.timerId` so
   * the original `dataset.dismissTimer` config value is preserved.
   *
   * @param {HTMLElement} toast - The toast element to auto-dismiss.
   * @param {boolean} isWithProgressBar - Whether to animate the progress bar.
   * @param {number} dismissTimer - Duration before auto-dismiss (in ms).
   * @param {number} removeDuration - Duration for removal animation (in ms).
   */
  function handleAutoDismiss(
    toast,
    isWithProgressBar,
    dismissTimer,
    removeDuration,
  ) {
    if (isWithProgressBar) animateProgressBar(toast, dismissTimer);

    toast.dataset.timerId = setTimeout(function () {
      performExitAnimationAndRemove(
        toast,
        toast.dataset.animation === "true",
        toast.dataset.position,
        removeDuration,
      );
    }, dismissTimer);
  }

  /**
   * Closes the toast by stopping propagation, clearing timers, and performing
   * exit animation. Reads the configured `removeDuration` from the toast's
   * data attribute instead of relying on the default.
   *
   * @param {Event} e - The event object.
   */
  function closeToast(e) {
    e.stopPropagation();

    const toast = e.target.closest('.gttToast[role="alert"]');
    if (toast) {
      clearTimeout(Number(toast.dataset.timerId));

      var removeDuration = Number(toast.dataset.removeDuration) || 500;

      performExitAnimationAndRemove(
        toast,
        toast.dataset.animation === "true",
        toast.dataset.position,
        removeDuration,
      );
    }
  }

  /**
   * Dismisses the most recently shown (last in DOM order) visible toast.
   * Used by the Escape key handler.
   */
  function dismissLastToast() {
    var toasts = document.querySelectorAll(
      '.gttToast[role="alert"]:not(.gttClose)',
    );
    if (toasts.length === 0) return;

    var lastToast = toasts[toasts.length - 1];
    clearTimeout(Number(lastToast.dataset.timerId));

    var removeDuration = Number(lastToast.dataset.removeDuration) || 500;

    performExitAnimationAndRemove(
      lastToast,
      lastToast.dataset.animation === "true",
      lastToast.dataset.position,
      removeDuration,
    );
  }

  /** Close toast on button click */
  document.body.addEventListener("click", function (e) {
    var closeButton = e.target.closest(".gttCloseBtn");
    if (closeButton) closeToast(e);
  });

  /** Dismiss the most recent toast when Escape is pressed */
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape") {
      dismissLastToast();
    }
  });

  /** Initialize auto-dismiss on page load */
  document.addEventListener("DOMContentLoaded", function () {
    document
      .querySelectorAll('.gttToast[data-auto-dismiss="true"]')
      .forEach(function (el) {
        var config = initGoaster({
          dismissTimer: Number(el.dataset.dismissTimer) || undefined,
          removeDuration: Number(el.dataset.removeDuration) || undefined,
        });

        var hasProgressBar = !!el.querySelector(".gttProgressBar");
        handleAutoDismiss(
          el,
          hasProgressBar,
          config.dismissTimer,
          config.removeDuration,
        );
      });
  });
})();
