<template>
  <div>
    <Row>
      <Col span="4">
        <Donut></Donut>
      </Col>
      <Col span="4">
        <BouncingLoader></BouncingLoader>
      </Col>
    </Row>
  </div>
</template>

<script>
import Donut from './Donut'
import BouncingLoader from './BouncingLoader'

export default {
  name: 'Index',
  components: { Donut, BouncingLoader }
}
</script>

<style lang="less" scoped>

</style>
