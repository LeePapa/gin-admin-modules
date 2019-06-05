<template>
  <Card title="Tabs切换隐藏导致图表显示问题">
    <Tabs @on-click="onClickTab" type="card" :animated="false">
      <TabPane label="标签1">
        <HChart ref="chart0"></HChart>
      </TabPane>
      <TabPane label="标签2">
        <HChart ref="chart1"></HChart>
      </TabPane>
      <TabPane label="标签3">
        <HChart ref="chart2"></HChart>
      </TabPane>
    </Tabs>
    <div v-show="show">
      <HChart ref="chart3"></HChart>
    </div>
  </Card>
</template>

<script>
import HChart from '@components/h-chart'
export default {
  name: 'TabsChart',
  components: {
    HChart
  },
  data () {
    return {
      show: false
    }
  },
  methods: {
    onClickTab (tab) {
      // this.$refs[`chart${tab}`].resize()
      this.show = !this.show
    }
  }
}
</script>

<style lang="less" scoped>

</style>
