<template>
  <Card title="只有线">
    <HChart only-line :data="[10, 20, 15, 60, 5]"></HChart>
  </Card>
</template>

<script>
import HChart from '@components/h-chart'
export default {
  name: 'OnlyLine',
  components: {
    HChart
  },
  created () {
    this.$bus.$on(this.$bus.homeResizeCharts, () => {
      console.log(this.$bus.homeResizeCharts)
    })
  }
}
</script>

<style lang="less" scoped>

</style>
