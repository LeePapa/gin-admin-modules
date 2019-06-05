<template>
  <div>
    <Table v-bind="tableAttrs"
           v-on="tableListeners"
           :columns="columnsNative"
           :data="dataSource"
           :loading="loading">
    </Table>
    <Row type="flex" justify="end" style="margin-top: 10px;" v-if="showPage">
      <Page :total="totalNumber"
            :current="page"
            :page-size="pageSize"
            v-bind="pageAttrs"
            v-on="pageListeners">
      </Page>
    </Row>
  </div>
</template>
<script>
/*
* page开头传入的属性或者事件作为Page组件的属性传入
* */
import listMock from './list'
import getColumns from './columns'

export default {
  name: 'HTable',
  inheritAttrs: false,
  props: {
    // 是否显示分页器
    showPage: {
      type: Boolean,
      default: true
    },
    // 除去page,pageSize两个查询参数之外的其他查询条件参数
    searchParamsDefault: {
      type: Object,
      default () {
        return {}
      }
    },
    // 接口调用的方法
    listApiMethod: {
      type: Function
    },
    formatDataSource: {
      type: Function
    },
    columns: {
      type: Array
    }
  },
  data () {
    return {
      loading: true,
      // 当前请求的页码
      page: 1,
      // 每页的记录数量
      pageSize: 10,
      // 总记录数
      totalNumber: 0,
      // 数据列表
      dataSource: [],
      columnsNative: this.columns,
      // 除去page,pageSize两个查询参数之外的其他查询条件参数
      searchParams: this.searchParamsDefault,
      tableAttrs: {},
      pageAttrs: {
        'show-total': true,
        'show-sizer': true,
        'show-elevator': true,
        'page-size-opts': [10, 20, 30, 50]
      },
      tableListeners: {},
      pageListeners: {}
    }
  },
  created () {
    console.log('this.$attrs:', this.$attrs)
    console.log('this.$listeners:', this.$listeners)
    Object.keys(this.$attrs).forEach(key => {
      if (key.startsWith('page-')) {
        this.pageAttrs[key.slice(5)] = this.$attrs[key]
      } else {
        this.tableAttrs[key] = this.$attrs[key]
      }
      key === 'loading' && (this.loading = this.$attrs[key])
      key === 'page-current' && (this.page = this.$attrs[key])
    })
    Object.keys(this.$listeners).forEach(key => {
      if (key.startsWith('page-')) {
        this.pageListeners[key.slice(5)] = this.$listeners[key]
      } else {
        this.tableListeners[key] = this.$listeners[key]
      }
    })
    this.pageListeners['on-page-size-change'] = this.pageListeners['on-page-size-change'] || this.onPageSizeChange
    this.pageListeners['on-change'] = this.pageListeners['on-change'] || this.onPageChange
    this.tableListeners['on-sort-change'] = this.tableListeners['on-sort-change'] || this.onSortChange
    this.doSearch()
  },
  methods: {
    onPageChange (page) {
      this.doSearch({}, page)
    },
    onPageSizeChange (size) {
      this.doSearch({}, 1, size)
    },
    onSortChange ({ key, order }) {
      if (order === 'normal') {
        this.searchParams.sortName = ''
        this.searchParams.sortType = ''
      } else {
        this.searchParams.sortName = key
        this.searchParams.sortType = order
      }
      this.getList()
    },
    // 列表获取数据统一在此处方法调接口
    getList () {
      if (!this.listApiMethod) {
        console.warn('请传入列表获取接口方法')
        this.columnsNative = this.columnsNative || getColumns(this)
        this.dataSource = this.formatDataSource ? this.formatDataSource(listMock) : listMock
        this.totalNumber = listMock.length
        this.loading = false
        this.$emit('complete', { data: listMock, success: true })
        return
      }
      this.loading = true
      let opt = {
        currentPage: this.page,
        pageSize: this.pageSize
      }
      opt = Object.assign(opt, this.searchParams)
      this.listApiMethod(opt).then((data) => {
        let list = (data.value.list || [])
        this.dataSource = this.formatDataSource ? this.formatDataSource(list) : list
        this.totalNumber = data.value.total
        this.loading = false
        this.$emit('complete', { data, success: true })
      }).catch(() => {
        this.loading = false
        this.$emit('complete', { success: false })
      })
    },
    doSearch (params, page, pageSize) {
      this.searchParams = Object.assign({}, this.searchParams, params || {})
      page && (this.page = page)
      pageSize && (this.pageSize = pageSize)
      this.getList()
    }
  }
}
</script>
