import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import plugin from './plugin'
// import EventProxy from '@util/plugin'
import 'iview/dist/styles/iview.css'
import 'hx-components/dist/lib/hx.min.css'
import comp from 'hx-components'

Vue.config.productionTip = false

// Vue.config.errorHandler = (err, vm, info) => {
//   console.error('err:', err)
//   console.error('vm:', vm)
//   console.error('info:', info)
//   // 处理错误
//   // `info` 是 Vue 特有的错误信息，例如，错误是在哪个生命周期钩子函数中发现的。
//   // info 只在 2.2.0+ 可访问
// }

console.log('我是test外', process.env.BRANCH_ENV)
if (process.env.BRANCH_ENV === 'test') {
  console.log('我是test里面')
  /* eslint-disable no-unused-vars */
  const vconsole = require('vconsole')
}

// 安装全局组件
Vue.use(plugin)
// 建立事件中心，做简单的全局状态监控，推荐换成vuex统一管理
// Vue.use(EventProxy)

Vue.use(comp)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App)
})
