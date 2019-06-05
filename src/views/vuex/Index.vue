<template>
  <div>
    测试vuex的使用
    <Child1 :data.sync="data"></Child1>
    <Child2></Child2>
    <Button @click.native="onClick">change</Button>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Child1 from './Child1'
import Child2 from './Child2'

export default {
  name: 'Index',
  components: { Child1, Child2 },
  data () {
    return {
      data: ''
    }
  },
  computed: {
    ...mapState({
      isLoading: state => state.vux.isLoading
    })
  },
  methods: {
    ...mapActions(['updateName']),
    onClick () {
      this.data = Math.random().toString()
    }
  },
  watch: {
    data: {
      handler (val) {
        console.log('ddd:', val)
      }
    }
  }
}
</script>

<style lang="less" scoped>

</style>
