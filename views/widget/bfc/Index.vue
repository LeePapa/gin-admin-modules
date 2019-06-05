<template>
  <div>
    <div class="container test">
      测试position:relative;
    </div>
    <div class="container test2">
      11
    </div>
    <div class="bfc">
      <div class="bfc-float">
        test1
      </div>
      <div class="bfc-float">
        test1
      </div>
      <div class="bfc-float">
        test1
      </div>
    </div>
    <div>
      <div class="left">
      <pre>
  .left{
    background:pink;
    float: left;
    width:180px;
  }
      </pre>
      </div>
      <div class="right">
       <pre>
  .right{
    background:lightblue;
    width:180px;
    float:right;
  }
      </pre>
      </div>
      <div class="center">
        <div class="pp">pp</div>
    <pre>
  .center{
    background:lightyellow;
    overflow:hidden;
    height:116px;
  }
      </pre>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Index'
}
</script>

<style lang="less" scoped>
  .container {
    display: inline-block;
    width: 200px;
    height: 200px;
    background-color: #f00;
    border: 5px solid #999;
  }
  .test {
    position: relative;
    left: 400px;
  }
  .test2 {
    position: relative;
    left: 100px;
  }

  .bfc {
    overflow: hidden;
    &-float {
      float: left;
      /*overflow: auto;*/
      width: 200px;
      height: 200px;
      background-color: #f00;
      border: 5px solid #999;
      margin: 20px;
    }
  }

  .left{
    background:pink;
    float: left;
    width:180px;
  }
  .center{
    position: relative;
    background:lightyellow;
    overflow:hidden;
  }
  .center .pp {
    position: absolute;
    right: 0;
    top: 0;
    width: 100px;
    height: 100px;
    background: #2D8CF0;
  }
  .right{
    background: lightblue;
    width:180px;
    float:right;
  }
</style>
