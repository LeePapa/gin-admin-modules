<template>
  <Card title="没有任何默认配置">
    <HChart :options="chartOptions"></HChart>
  </Card>
</template>

<script>
import HChart from '@components/h-chart'

export default {
  name: 'OriginalChart',
  components: {
    HChart
  },
  data () {
    return {
      chartOptions: {
        color: ['#5a6580', '#1f95e7', '#57c7d3'],
        tooltip: {
          trigger: 'axis',
          axisPointer: { // 坐标轴指示器，坐标轴触发有效
            type: 'shadow' // 默认为直线，可选为：'line' | 'shadow'
          }
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: [
          {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
            axisTick: {
              alignWithLabel: true
            }
          }
        ],
        yAxis: [
          {
            type: 'value'
          }
        ],
        series: [
          {
            name: '柱状1',
            type: 'bar',
            barWidth: '40%',
            data: [10, 52, 200, 334, 390, 330, 220]
          },
          {
            name: '柱状2',
            type: 'bar',
            barWidth: '40%',
            data: [10, 32, 100, 334, 123, 330, 234]
          },
          {
            name: '折线1',
            type: 'line',
            data: [10, 52, 44, 55, 390, 330, 232]
          }
        ]
      }
    }
  }
}
</script>

<style lang="less">

</style>
