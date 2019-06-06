<template>
  <div>
    <AspectBox :ratio="4" width="200px">
      <div class="box1">
        宽度为200px,宽高比为4:1的盒子
      </div>
    </AspectBox>
    <br>
    <AspectBox :ratio="20">
      <div class="box1">
        宽度为100%,宽高比为20:1的盒子
      </div>
    </AspectBox>
  </div>
</template>

<script>
import AspectBox from '@components/aspect-box'

export default {
  name: 'Index',
  components: {
    AspectBox
  }
}
</script>

<style lang="less" scoped>
  .box1 {
    width: 100%;
    height: 100%;
    background: #2D8CF0;
  }
</style>
