<template>
  <div class="layout-header">
    <div class="layout-header-logo">
      后台模板
    </div>
    <div class="layout-header-main">
      <Dropdown placement="bottom-end" @on-click="onToggleHeader">
        <img class="avatar" :src="avatar" alt="头像">
        <DropdownMenu slot="list">
          <DropdownItem name="0">{{userName}}</DropdownItem>
          <DropdownItem name="1" divided>退出</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </div>
  </div>
</template>

<script>
import storage from '@util/storage'
// import loginApi from '@service/system'
import { mapState } from 'vuex'

export default {
  name: 'LayoutHeader',
  data () {
    return {
      // userName: 'heroxiao',
      avatar: require('@assets/logo.png')
    }
  },
  computed: {
    ...mapState(['userName'])
  },
  methods: {
    onToggleHeader (name) {
      if (name === '1') {
        this.loginOut()
      }
    },
    loginOut () {
      storage.setUserName()
      this.$router.replace('/login')
      // loginApi.loginOut()
      //   .then(data => {
      //     storage.setUserName()
      //     this.$router.replace('/login')
      //   })
    }
  }
}
</script>

<style lang="less" scoped>
  .avatar {
    width: 44px;
    height: 44px;
    border-radius: 50%;
    cursor: pointer;
  }
</style>
