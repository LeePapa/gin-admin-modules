<template>
  <div>
    我是name:{{name}}
    {{dataNative}}
    <Button @click.native="onClick">改data</Button>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Child1',
  props: {
    data: {
      type: String
    }
  },
  computed: {
    ...mapState({
      name: state => state.demo.name
    }),
    dataNative: {
      get () {
        return this.data
      },
      set (val) {
        this.$emit('update:data', val)
      }
    }
  },
  methods: {
    onClick () {
      this.dataNative = Math.random().toString()
    }
  }
}
</script>

<style lang="less" scoped>

</style>
