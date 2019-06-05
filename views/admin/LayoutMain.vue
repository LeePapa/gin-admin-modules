<template>
  <div id="layout-main-content" class="layout-main-content">
    <div class="layout-main-content-wrap">
      <!--页面默认都是缓存的，只有当meta.keepAlive===false时不缓存-->
      <transition name="main">
        <keep-alive>
          <router-view v-if="$route.meta.keepAlive===undefined||$route.meta.keepAlive===true"></router-view>
        </keep-alive>
      </transition>
      <transition name="main">
        <router-view v-if="$route.meta.keepAlive===false"></router-view>
      </transition>
      <div class="layout-copyright">
        2016-{{latest}} &copy; heroxiao
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LayoutMain',
  computed: {
    latest () {
      return new Date().getFullYear()
    }
  },
  mounted () {
    this.$nextTick(() => {
      console.log('keep:', this.$refs.keep)
    })
  }
}
</script>

<style lang="less" scoped>
</style>
