/**
 * Created by xuwei on 2018/1/15.
 */
import data from '@util/data/list'
import getColumns from '@util/data/columns'
export default {
  data () {
    return {
      // 当前请求的页码
      page: 1,
      // 每页的记录数量
      pageSize: 10,
      pageSizes: [10, 20, 30, 50],
      // 总记录数
      totalNumber: 0,
      columns: [],
      // 数据列表
      dataSource: [],
      // 当前接口加载情况
      loading: true,
      // 除去page,pageSize两个查询参数之外的其他查询条件参数
      searchParams: {},
      // 接口调用的方法
      listApiMethod: null,
      // 当前选中的列表项
      selection: []
    }
  },
  created () {
    this.getList()
  },
  methods: {
    onSelectionChange (selection) {
      this.selection = selection
    },
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
    getList (cb) {
      if (!this.listApiMethod) {
        console.warn('请传入列表获取接口方法')
        this.columns = this.columns || getColumns(this)
        this.dataSource = this.formatDataSource ? this.formatDataSource(data) : data
        this.totalNumber = data.length
        this.loading = false
        return
      }
      this.selection = []
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
        cb && cb(data)
      }).catch(() => {
        this.loading = false
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
