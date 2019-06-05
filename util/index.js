/**
 * Created by joylee on 2018/6/14.
 */
// 点击页面（除当前需要隐藏的块）隐藏当前块
export function documentTapHide (func) {
  const bindHide = function () {
    func && func()
    document.removeEventListener('click', bindHide, false)
  }
  setTimeout(function () {
    document.addEventListener('click', bindHide, false)
  }, 500)
}
