/**
 * Created by xuwei on 2018/4/9.
 * 默认是折线图的统一配置
 */
import common from './common'
// 设置些假数据
const legendData = ['类目1', '类目2']
const series = legendData.map(item => ({ name: item, type: 'line', smooth: true, data: common.getRandomArray() }))
export default {
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    // show: false,
    type: 'scroll',
    data: legendData,
    top: 0
  },
  grid: {
    left: '30px',
    right: '40px',
    bottom: '30px',
    top: '40px',
    containLabel: true
  },
  toolbox: {
    feature: {
      saveAsImage: {}
    }
  },
  xAxis: [{
    type: 'category',
    boundaryGap: false,
    data: common.getMockXAxis()
  }],
  yAxis: [{
    type: 'value'
  }],
  color: common.color,
  series: series
}
