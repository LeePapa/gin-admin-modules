<template>
  <div>
    <div @click="onClick(0)">
      <div class="hairline-border">text</div>
      <div class="trigger">
        Hover me to see a height transition.
        <div class="el">content</div>
      </div>
      <Input type="text" v-model="number" @on-blur="onBlur"/>
      <p>{{tweenedNumber.toFixed(0)}}</p>
      <p>冒泡捕获测试：</p>
      <div @click.capture="onClick(1)">
        console.log(1)
        <div @click.capture="onClick(2)">
          console.log(2)
          <div @click="onClick(3)">
            console.log(3)
            <div @click="onClick(4)">
              console.log(4)
            </div>
          </div>
        </div>
      </div>
      <AspectBox :ratio="4" style="background-color: #ff9900" ref="box">
        <div><Input placeholder="请输入"/></div>
        <!--我是注释-->
        煮熟吗？
        <Input placeholder="请输入"/>
        <Input placeholder="请输入"/>
        <Input placeholder="请输入"/>
      </AspectBox>
    </div>
    <Throttle events="click,mouseenter" :time="1000">
      <div @click="onThrottle" @mouseenter="onttt">klsdfkjaflkj</div>
      <!--<Button type="primary" @click="onThrottle" @mouseenter.native="onttt">测试Throttle</Button>-->
    </Throttle>
    <Debounce events="click,mouseenter" :time="1000">
      <Button type="primary" @click="onThrottle" @mouseenter.native="onttt">测试Throttle</Button>
    </Debounce>
  </div>
</template>

<script>
import AspectBox from '@components/aspect-box'
import decorate from './decorate'

export default {
  name: 'Index',
  components: { AspectBox, Throttle: decorate('throttle'), Debounce: decorate('debounce') },
  data () {
    return {
      number: 0,
      tweenedNumber: 0
    }
  },
  created () {
    this.$nextTick(() => {
      let el = document.querySelector('.el')// xss,csrf,钓鱼网站，网络劫持，控制台
      let height = el.scrollHeight
      el.style.setProperty('--max-height', height + 'px')
      console.log('box:', this.$refs.box.$children)
    })
  },
  methods: {
    onBlur () {
      // window.TweenLite.to(this.$data, 0.5, { tweenedNumber: this.number })
    },
    onClick (val) {
      console.log(val)
    },
    onThrottle () {
      console.log('onThrottle')
    },
    onttt () {
      console.log('mouseenter')
    }
  }
}
</script>

<style lang="less" scoped>
  .hairline-border {
    box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.1);
  }

  @media (min-resolution: 2dppx) {
    .hairline-border {
      box-shadow: 0 0 0 0.5px;
    }
  }

  @media (min-resolution: 3dppx) {
    .hairline-border {
      box-shadow: 0 0 0 0.33333333px;
    }
  }

  @media (min-resolution: 4dppx) {
    .hairline-border {
      box-shadow: 0 0 0 0.25px;
    }
  }

  .el {
    transition: max-height 0.5s;
    overflow: hidden;
    max-height: 0;
  }

  .trigger {
    background-color: rgba(0, 0, 0, .2);
    padding: 20px;
  }

  .trigger:hover > .el {
    max-height: var(--max-height);
  }
</style>
