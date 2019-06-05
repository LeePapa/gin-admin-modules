/**
 * Created by xuwei on 2018/3/30.
 */
const paths = [
  // -1为假设的特殊id，表示一级菜单,其余为带有子菜单
  { id: '-1', urls: [{ title: '首页', path: '/', icon: 'ios-home' }] },
  {
    id: '0',
    title: '图表',
    icon: 'logo-buffer',
    urls: [{ title: '图表', path: '/chart' }]
  },
  {
    id: '1',
    title: '表格',
    icon: 'ios-grid',
    urls: [{ title: '表格', path: '/table' }]
  },
  {
    id: '2',
    title: '组件',
    icon: 'ios-keypad',
    urls: [
      { title: '链接', path: '/widget/link' },
      { title: '栅格', path: '/widget/grid' },
      { title: '图片', path: '/widget/image' },
      { title: '拖拽', path: '/widget/drag' },
      { title: '主题', path: '/widget/theme' },
      { title: 'BFC', path: '/widget/bfc' },
      { title: 'lottie', path: '/widget/lottie' },
      { title: 'editor', path: '/widget/editor' },
      { title: '比例盒子', path: '/widget/box' }
    ]
  },
  {
    id: '3',
    title: '图标',
    icon: 'ios-grid',
    urls: [{ title: '图标', path: '/icon' }]
  },
  {
    id: '4',
    title: '测试',
    icon: 'ios-grid',
    urls: [{ title: '测试', path: '/test' }]
  }
]

const activeMenu = (path) => {
  let obj = {}
  try {
    paths.forEach(item => {
      item.urls.forEach(im => {
        if (path === im.path) {
          obj.openNames = [item.id]
          obj.activeName = im.path
          throw new Error('break')
        }
      })
    })
  } catch (e) {}
  return obj
}

export default {
  paths,
  activeMenu
}
