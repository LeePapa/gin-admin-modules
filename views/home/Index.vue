<template>
  <div>
    我是首页
    <Input v-model="userName" placeholder="异步接口修改用户信息" @input.native="onTryUpdateUserName"/>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      userName: ''
    }
  },
  methods: {
    onTryUpdateUserName () {
      // this.$bus.$emit(this.$bus.changeUserName, this.userName)
      // this.$emit('global:username', this.userName)
      this.$store.commit('updateUserName', { userName: this.userName })
    }
  }
}
</script>

<style lang="less">
</style>
