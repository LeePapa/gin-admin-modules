/**
 * Created by xuwei on 2018/1/15.
 */
// import Link from '@components/link'
// import {Tag} from 'iview'

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
    width: 160,
    render (h, params) {
      // const redirect = () => {
      //   context.$router.push({path: `/table/detail/${params.id}`})
      // }
      // return <Link nativeOnClick={redirect}>
      //   查看详情
      // </Link>
      // return h('InputNumber', {
      //   props: {
      //     size: 'small',
      //     max: 100,
      //     min: 1
      //   },
      //   on: {
      //     'on-change': (val) => {
      //       console.log(val)
      //       // params.row.appliedVal = event.target.value
      //     },
      //     'on-focus' (event) {
      //       console.log(event.target.value)
      //     }
      //   }
      // })
      return <i-select
        on-on-change={val => console.log(val)}
        clearable
        on-on-open-change={val => console.log('on-open-change')}
        size='small'>
        <i-option value="1">nnn</i-option>
        <i-option value="2">nnn</i-option>
      </i-select>
      // return <Tag type="border" closable color="primary" on-on-close={val => console.log('close')}>标签一</Tag>
      // return <InputNumber size='small' max={100} min={1} onOn-change={val => console.log(val)} onOnBlur={val => console.log('blur')}></InputNumber>
    }
  }
]
