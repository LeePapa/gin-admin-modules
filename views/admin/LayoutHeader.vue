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
  created () {
    /* this.$on('global:username', (name) => {
      this.userName = name
    }) */
    /* this.$bus.$on(this.$bus.changeUserName, (name) => {
      this.userName = name
    }) */
  },
  methods: {
    onToggleHeader (name) {
      this.$emit('update', name)
      if (name === '1') {
        // this.loginOut()
      }
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
