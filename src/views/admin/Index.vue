<template>
  <div class="layout">
    <LayoutHeader @update="onUpdate"></LayoutHeader>
    <div class="layout-content">
      <LayoutSider :path="$route.path"></LayoutSider>
      <LayoutMain></LayoutMain>
    </div>
  </div>
</template>

<script>
import LayoutHeader from './LayoutHeader'
import LayoutSider from './LayoutSider'
import LayoutMain from './LayoutMain'
export default {
  name: 'Index',
  components: {
    LayoutHeader,
    LayoutSider,
    LayoutMain
  },
  methods: {
    onUpdate (val) {
      console.log('val:', val)
    }
  }
}
</script>

<style lang="less">
  @import "./style";
</style>
