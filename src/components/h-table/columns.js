/**
 * Created by xuwei on 2018/1/15.
 */
import Link from '@components/link'
export default context => [
  {
    type: 'selection',
    width: 60,
    align: 'center'
  },
  {
    title: '序号',
    key: 'rank',
    width: 80
  },
  {
    title: '商品',
    key: 'commodity',
    sortable: 'custom'
  },
  {
    title: '站点',
    key: 'website'
  },
  {
    title: '价格($)',
    key: 'price'
  },
  {
    title: '上架时间',
    key: 'upDate'
  },
  {
    title: '评论数量',
    key: 'comment',
    sortable: 'custom',
    sortType: 'desc'
  },
  {
    title: '操作',
    width: 100,
    render (h, params) {
      return <Link nativeOnClick={() => context.showModal(params)}>
        查看详情
      </Link>
    }
  }
]
