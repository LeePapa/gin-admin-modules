<template>
  <Card title="折线图">
    <HChart></HChart>
  </Card>
</template>

<script>
import HChart from '@components/h-chart'

const legendResults = ['name1', 'name2', 'name3', 'name4', 'name5']

export default {
  name: 'HLine',
  components: {
    HChart
  },
  data () {
    return {
      resultsOptions: {
        grid: {
          top: '60px'
        },
        legend: {
          data: legendResults
        },
        xAxis: {
          data: []
        },
        yAxis: [{ type: 'value' }],
        series: []
      }
    }
  }
}
</script>

<style lang="less">

</style>
