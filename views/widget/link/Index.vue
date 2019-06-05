<template>
  <div>
    <div>
      <HLink @click.native="onShowModal">去grid</HLink>
    </div>
    <div>
      <HLink to="https://www.baidu.com">跳转到外部链接在当前页</HLink>
    </div>
    <div>
      <HLink to="https://www.baidu.com" blank :underline="false">跳转到外部链接打开新页</HLink>
    </div>
    <div>
      <HLink :to="{path: '/widget/grid'}" plain>基本的原始文字样式</HLink>
    </div>
  </div>
</template>

<script>
import HLink from '@components/link'
export default {
  name: 'Index',
  components: {
    HLink
  },
  methods: {
    onShowModal () {
      this.$Modal.success({ title: '提示下' })
    }
  }
}
</script>

<style lang="less" scoped>

</style>
