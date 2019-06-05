<template>
  <div>
    <Input v-model="name" placeholder="Enter something..." style="width: 300px" />
  </div>
</template>

<script>
// import {mapState} from 'vuex'
export default {
  name: 'Child2',
  computed: {
    name: {
      get () {
        return this.$store.state.demo.name
      },
      set (val) {
        this.$store.commit('demo/updateName', { name: val })
      }
    }
  }
}
</script>

<style lang="less" scoped>

</style>
