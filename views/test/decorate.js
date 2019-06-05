/**
 * Created by joylee on 2019/4/24.
 * 节流组件
 */
const throttle = function (fn, wait = 50, ctx) {
  let lastCall = 0
  return function (...params) {
    const now = new Date().getTime()
    if (now - lastCall < wait) return
    lastCall = now
    fn.apply(ctx, params)
  }
}

const debounce = function (fn, wait = 50, ctx) {
  return function (...params) {
    clearTimeout(fn.tId)
    fn.tId = setTimeout(function () {
      fn.apply(ctx, params)
    }, wait)
  }
}

const func = (name) => {
  return {
    name: name,
    abstract: true,
    props: {
      time: Number,
      events: String
    },
    created () {
      this.eventKeys = this.events.split(',')
      this.originMap = {}
      this.throttledMap = {}
    },
    // render函数直接返回slot的vnode，避免外层添加包裹元素
    render (h) {
      const vnode = this.$slots.default[0]
      let obj1 = {}
      let obj2 = vnode.data.on || {}
      // 组件
      if (vnode.tag && vnode.componentOptions) {
        obj1 = vnode.componentOptions.listeners || {}
      }
      // 此时组件实例还没有
      console.log('componentInstance:', vnode.componentInstance)
      this.eventKeys.forEach((key) => {
        let obj = obj1
        obj2[key] && (obj = obj2)
        const target = obj[key]
        if (target === this.originMap[key] && this.throttledMap[key]) {
          obj[key] = this.throttledMap[key]
        } else if (target) {
          // 将原本的事件处理函数替换成throttle节流后的处理函数
          this.originMap[key] = target
          const fn = name === 'throttle' ? throttle : debounce
          this.throttledMap[key] = fn(target, this.time, vnode)
          obj[key] = this.throttledMap[key]
        }
      })
      return vnode
    }
  }
}

export default func
