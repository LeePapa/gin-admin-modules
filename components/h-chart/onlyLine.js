/**
 * Created by xuwei on 2018/3/26.
 * 仅仅只显示单个折线，没有横纵坐标，适用于表格中的折线图项
 */
import common from './common'
let base = {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      lineStyle: {
        opacity: 0
      }
    },
    formatter: '&nbsp;{c0}&nbsp;'
  },
  grid: {
    show: false,
    left: '5px',
    right: '5px',
    top: '5px',
    bottom: '5px',
    containLabel: false
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    show: false
    // data: []
  },
  yAxis: {
    type: 'value',
    show: false
  },
  series: [
    {
      name: 'onlyLine',
      type: 'line',
      data: [], // 只修改该字段，其他统一定义好，不要做太多修改
      showSymbol: false,
      symbolSize: 1,
      animationEasing: 'easing',
      animationDuration: 1000,
      areaStyle: { normal: {} },
      smooth: true
    }
  ],
  color: common.color
}

const setData = (options, data) => {
  options.xAxis.data = data // 保证xAxis.data的长度与series[0].data一致，值随意因为不显示横坐标
  options.series[0].data = data
  return options
}

export default {
  base,
  setData
}
