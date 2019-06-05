/**
 * Created by joylee on 2019/5/9.
 * 最简单的nextTick，基于Promise微任务实现
 */
export const nextTick = (function () {
  const callbacks = []
  let pending = false
  const nextTickHandler = function () {
    pending = false
    const copies = callbacks.slice(0)
    callbacks.length = 0
    copies.map(cb => cb())
  }
  const timerFunc = function () {
    Promise.resolve().then(nextTickHandler)
  }
  return function (cb, ctx) {
    let _resolve
    if (cb) {
      callbacks.push(cb)
    } else if (_resolve) {
      _resolve(ctx)
    }
    if (!pending) {
      pending = true
      timerFunc()
    }
    if (!cb) {
      return new Promise((resolve, reject) => {
        _resolve = resolve
      })
    }
  }
})()
