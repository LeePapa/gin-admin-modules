<template>
  <div>
    <Row :gutter="24">
      <Col span="12">
        <OnlyLine></OnlyLine>
      </Col>
      <Col span="12">
        <Pie></Pie>
      </Col>
      <Col span="12">
        <HLine></HLine>
      </Col>
      <Col span="12">
        <OriginalChart></OriginalChart>
      </Col>
      <Col span="12">
        <TabsChart></TabsChart>
      </Col>
      <Col span="12">
        <CalendarChart></CalendarChart>
      </Col>
    </Row>
    <Button type="primary" @click.native="onResize">测试下事件总线bus</Button>
  </div>
</template>

<script>
import OnlyLine from './OnlyLine.vue'
import Pie from './Pie.vue'
import HLine from './Line.vue'
import OriginalChart from './OriginalChart.vue'
import TabsChart from './TabsChart.vue'
import CalendarChart from './CalendarChart.vue'
export default {
  name: 'Index',
  components: {
    OnlyLine,
    Pie,
    HLine,
    OriginalChart,
    TabsChart,
    CalendarChart
  },
  methods: {
    onResize () {
      this.$bus.$emit(this.$bus.homeResizeCharts, true)
    }
  }
}
</script>

<style lang="less">

</style>
