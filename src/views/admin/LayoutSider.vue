<template>
  <div class="layout-sider">
    <Menu theme="dark" :active-name="activeMenu.activeName" :open-names="activeMenu.openNames" width="auto"
          @on-select="onTapMenu" accordion>
      <div v-for="(item, index) in paths" :key="index">
        <template v-if="item.id === '-1'">
          <MenuItem :name="im.path" v-for="(im,idx) in item.urls" :key="idx">
            <Icon :type="im.icon" :size="iconSize" :title="isFold?im.title:''"></Icon>
            <span v-show="!isFold">{{im.title}}</span>
          </MenuItem>
        </template>
        <template v-else>
          <Submenu :name="item.id">
            <template slot="title">
              <Icon :type="item.icon" :size="iconSize"></Icon>
              {{item.title}}
            </template>
            <MenuItem :name="im.path" v-for="(im, idx) in item.urls" :key="idx">{{im.title}}</MenuItem>
          </Submenu>
        </template>
      </div>
    </Menu>
  </div>
</template>

<script>
import menu from './menu'

export default {
  name: 'LayoutSider',
  props: {
    path: {
      type: String,
      default: '/'
    }
  },
  data () {
    return {
      iconSize: 18,
      paths: menu.paths,
      isFold: false,
      activeMenu: {}
    }
  },
  methods: {
    onTapMenu (path) {
      this.$router.push(path)
    }
  },
  watch: {
    path: {
      handler (val) {
        this.activeMenu = menu.activeMenu(val)
      },
      immediate: true
    }
  }
}
</script>

<style lang="less" scoped>
</style>
