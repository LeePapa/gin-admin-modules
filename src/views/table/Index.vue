<template>
  <div>
    <Tag type="border" closable color="primary">标签一</Tag>
    <MTable
      @page-on-page-size-change="onPageSizeChange"
      @on-selection-change="onSelectionChange"
      :format-data-source="formatDataSource"
      @complete="onLoaded"
      :show-page="true"
      :columns="columns"></MTable>
    <MTable
      @complete="onLoaded"
      :show-page="false"
      :columns="columns"></MTable>
    <MTable
      @complete="onLoaded"></MTable>
  </div>
</template>

<script>
import getColumns from './columns'
import MTable from '@components/h-table'

export default {
  name: 'Index',
  components: { MTable },
  data () {
    const columns = getColumns(this)
    return {
      columns: columns
    }
  },
  beforeCreate () {
    console.time('start')
  },
  // 第一次进入keep-alive路由组件时
  created () {
    console.log('created')
    this.isFirstEnter = true
    // 只有第一次进入或者刷新页面后才会执行此钩子函数，使用keep-alive后（2+次）进入不会再执行此钩子函数
  },
  mounted () {
    console.log('mounted')
    console.timeEnd('start')
  },
  beforeRouteEnter (to, from, next) {
    console.log('beforeRouteEnter')
    // 判断是从哪个路由过来的，若是detail页面不需要刷新获取新数据，直接用之前缓存的数据即可
    if (from.name === 'Detail') {
      to.meta.isBack = true
    }
    next()
  },
  activated () {
    console.log('activated')
    if (!this.$route.meta.isBack || this.isFirstEnter) {
      // 如果isBack是false，表明需要获取新数据，否则就不再请求，直接使用缓存的数据
      // 如果isFirstEnter是true，表明是第一次进入此页面或用户刷新了页面，需获取新数据
      // this.data = ''// 把数据清空，可以稍微避免让用户看到之前缓存的数据
      // this.getData() // ajax获取数据方法
      console.log('ajax获取数据...')
    }
    // 恢复成默认的false，避免isBack一直是true，导致下次无法获取数据
    this.$route.meta.isBack = false
    // 恢复成默认的false，避免isBack一直是true，导致每次都获取新数据
    this.isFirstEnter = false
  },
  deactivated () {
    console.log('deactivated')
  },
  methods: {
    onSelectionChange (selection) {
      // this.columns = Object.freeze(getColumns(this))
      // v-if不与v-for一起使用，数据量大字段时，不需要双向绑定的数据只传入初始默认值，一些不需要响应的数据可以考虑用Object.freeze锁定不让proxy,数据的遍历查找可以考虑使用map字典方式
      // keep-alive的使用，js分包
      console.log(selection)
    },
    onPageSizeChange () {
      console.log('.....')
    },
    onLoaded (result) {
      console.log('result:', result)
    },
    formatDataSource (list) {
      return [{
        reasonOption: '1'
      }, {
        reasonOption: '2'
      }]
    }
  }
}
</script>

<style lang="less" scoped>

</style>
