/**
 * Created by xuwei on 2018/5/23.
 * 单个饼图的配置
 */
import common from './common'
// 设置些假数据
const legendData = ['legend1', 'legend2']
const random = common.getRandomArray(2)
const seriesData = random.map((item, index) => ({ value: item, name: legendData[index] }))
export default {
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)'
  },
  legend: {
    show: false,
    orient: 'vertical',
    x: 'left',
    data: legendData
  },
  color: common.color,
  series: [
    {
      name: 'name1',
      type: 'pie',
      radius: ['0%', '80%'],
      avoidLabelOverlap: true,
      label: {
        normal: {
          position: 'outside'
        }
        // emphasis: {
        //   show: false,
        //   textStyle: {
        //     fontSize: '20',
        //     fontWeight: 'bold'
        //   }
        // }
      },
      labelLine: {
        normal: {
          show: true
        }
      },
      data: seriesData
    }
  ]
}
