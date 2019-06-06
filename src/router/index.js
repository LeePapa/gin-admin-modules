import Vue from 'vue'
import Router from 'vue-router'
import { LoadingBar } from 'iview'

const Home = () => import('@views/home/Index.vue')
const Admin = (resolve) => require(['@views/admin/Index.vue'], resolve)

// -------错误类
// 404
const NotFound = () => import('@views/error/notfound/Index.vue')

// -------账户类
const Account = () => import('@views/account/Index.vue')
const Login = (resolve) => require(['@views/account/login/Index.vue'], resolve)

// -------chart
const Chart = (resolve) => require(['@views/chart/Index.vue'], resolve)

// -------table
const Table = (resolve) => require(['@views/table/Index.vue'], resolve)
const Detail = (resolve) => require(['@views/table/Detail.vue'], resolve)
const Ajax = (resolve) => require(['@views/ajax/Index.vue'], resolve)

// -------vuex
const VuexC = (resolve) => require(['@views/vuex/Index.vue'], resolve)

// -------widget
const Link = (resolve) => require(['@views/widget/link/Index.vue'], resolve)
const Image = (resolve) => require(['@views/widget/image/Index.vue'], resolve)
const Grid = (resolve) => require(['@views/widget/grid/Index.vue'], resolve)
const AspectBox = (resolve) => require(['@views/widget/aspect-box/Index.vue'], resolve)
const Drag = (resolve) => require(['@views/widget/drag/Index.vue'], resolve)
const Theme = (resolve) => require(['@views/widget/theme/Index.vue'], resolve)
const Bfc = (resolve) => require(['@views/widget/bfc/Index.vue'], resolve)
const Lottie = (resolve) => require(['@views/widget/lottie/Index.vue'], resolve)
const Editor = (resolve) => require(['@views/widget/editor/Index.vue'], resolve)

// -------icon
const Icon = (resolve) => require(['@views/icon/Index.vue'], resolve)

// -------专用测试
const Test = (resolve) => require(['@views/test/Index.vue'], resolve)

Vue.use(Router)

// 默认页面都是keepAlive的，不需要写meta.keepAlive
const routes = [
  {
    path: '/',
    component: Admin,
    children: [
      {
        path: '',
        name: 'Home',
        component: Home
      },
      {
        path: 'chart',
        name: 'Chart',
        meta: {
          keepAlive: false
        },
        component: Chart
      },
      {
        path: 'table',
        name: 'Table',
        meta: {
          keepAlive: true,
          isBack: false
        },
        component: Table
      },
      {
        path: 'table/detail/:id',
        name: 'Detail',
        meta: {
          keepAlive: false
        },
        component: Detail
      },
      {
        path: 'ajax',
        name: 'Ajax',
        component: Ajax
      },
      {
        path: 'widget/link',
        name: 'Link',
        component: Link
      },
      {
        path: 'widget/image',
        name: 'Image',
        component: Image
      },
      {
        path: 'widget/grid',
        name: 'Grid',
        component: Grid
      },
      {
        path: 'widget/box',
        name: 'AspectBox',
        component: AspectBox
      },
      {
        path: 'widget/drag',
        name: 'Drag',
        component: Drag
      },
      {
        path: 'widget/theme',
        name: 'Theme',
        component: Theme
      },
      {
        path: 'widget/bfc',
        name: 'BFC',
        component: Bfc
      },
      {
        path: 'widget/lottie',
        name: 'Lottie',
        component: Lottie
      },
      {
        path: 'widget/editor',
        name: 'Editor',
        component: Editor
      },
      {
        path: 'vuex',
        name: 'VuexC',
        component: VuexC
      },
      {
        path: 'icon',
        name: 'Icon',
        component: Icon
      },
      {
        path: 'test',
        name: 'Test',
        component: Test
      }
    ]
  },
  {
    path: '/',
    component: Account,
    children: [
      {
        path: 'login',
        name: 'login',
        meta: {
          title: '商家登录'
        },
        component: Login
      }
    ]
  },
  {
    path: '*',
    meta: {
      title: '404-notfound'
    },
    component: NotFound
  }
]

const router = new Router({
  routes,
  linkActiveClass: 'active',
  mode: 'history'
})

router.beforeEach((to, from, next) => {
  console.log(to.path)
  LoadingBar.start()
  window.document.title = to.meta.title || '后台'
  next()
})

router.afterEach(() => {
  LoadingBar.finish()
  // id: layout-main-content,来源于src/views/admin/LayoutMain.vue
  Vue.nextTick(() => {
    document.getElementById('layout-main-content') && document.getElementById('layout-main-content').scrollTo(0, 0)
  })
})

export default router
