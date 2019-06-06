<template>
  <div class="body-wrap">
    <header class="header">
      <div class="header-left">
        运营后台
      </div>
      <div class="header-right">
      </div>
    </header>
    <router-view></router-view>
    <footer class="footer">
      <div class="footer-left">
        <span class="txt1">咨询：</span><br/>
        <span class="txt2">XXX-XXX-XXXX</span>
      </div>
      <div class="footer-right">
        <a href="javascript:;">关于我们</a>&nbsp;|&nbsp;
        <a href="javascript:;">服务条款</a>&nbsp;|&nbsp;
        <a href="javascript:;">&copy;XXX 备案号</a>
      </div>
    </footer>
  </div>
</template>
<style lang="less" scoped>
  @import "~@base/fn";

  @header-height: 80px;
  @bottom-height: 100px;
  .body-wrap {
    position: relative;
    min-height: 100%;
    min-width: @min-width;
    padding-bottom: @bottom-height;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #fff;
    background-color: #333;
    height: @header-height;
    width: 100%;
    padding: 0 16px;
    font-size: 15px;
    a {
      color: #fff;
    }
    &-left {
      font-size: 24px;
    }
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    height: @bottom-height;
    background: #f5f7f9;
    &-left {
      padding-right: 60px;
      border-right: 1px solid @border-color;
    }
    &-right {
      a {
        color: #333;
      }
      font-size: 12px;
      padding-left: 60px;
    }
    .txt1 {
      font-size: 16px;
    }
    .txt2 {
      font-size: 18px;
      font-weight: 700;
    }
  }
</style>
<script>
export default {
  data () {
    return {}
  }
}
</script>
