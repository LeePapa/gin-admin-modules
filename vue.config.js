/**
 * Created by joylee on 2019/4/17.
 */
const path = require('path')
function resolve (dir) {
  return path.join(__dirname, dir)
}
module.exports = {
  lintOnSave: false,
  configureWebpack: {
    resolve: {
      alias: {
        '@src': resolve('src'),
        // 'vue$': 'vue/dist/vue.esm.js',
        '@service': resolve('src/service'),
        '@components': resolve('src/components'),
        '@views': resolve('src/views'),
        '@util': resolve('src/util'),
        '@mixin': resolve('src/mixin'),
        '@store': resolve('src/store'),
        '@router': resolve('src/router'),
        '@assets': resolve('src/assets'),
        '@base': resolve('src/assets/css/base')
      }
    }
  },
  pluginOptions: {
    'style-resources-loader': {
      patterns: [
        resolve('src/assets/styles/base/vars.less'),
        resolve('src/assets/styles/base/mixins.less')
      ],
      preProcessor: 'less'
    }
  }
}
