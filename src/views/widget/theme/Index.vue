<template>
  <div class="theme1">
    <div class="nav">
      <div class="nav-item">12</div>
      <div class="nav-item">22</div>
      <div class="nav-item">23</div>
      <div class="nav-item">334</div>
    </div>
    <Button type="primary" @click.native="onClick">自定义css属性，主题切换</Button>
    <div>
      测试下：
      <hx-alert>
        alert测试内容
      </hx-alert>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Index',
  methods: {
    onClick () {
      const html = document.documentElement
      html.style.setProperty('--bgcolor', '#0f0')
      html.style.setProperty('--color', '#000')
    }
  }
}
</script>

<style lang="less">
  :root {
    --bgcolor: #f00;
    --color: #fff;
  }
</style>
<style lang="less" scoped>
  .nav {
    display: flex;
    height: 50px;
    line-height: 50px;
    background: var(--bgcolor);
    color: var(--color);
    &-item {
      flex: 1;
    }
  }
</style>
