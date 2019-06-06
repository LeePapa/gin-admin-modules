<template>
  <Card title="饼状图">
    <HChart type="pie" :options="chartOptions"></HChart>
  </Card>
</template>

<script>
import HChart from '@components/h-chart'

export default {
  name: 'Pie',
  components: {
    HChart
  },
  data () {
    return {
      chartOptions: {
        legend: {
          data: ['新访客数', '老访客数']
        },
        series: [
          {
            name: '新老访客',
            data: [
              { value: 335, name: '新访客数' },
              { value: 310, name: '老访客数' }
            ]
          }
        ]
      }
    }
  }
}
</script>

<style lang="less" scoped>

</style>
