<template>
  <div>

  </div>
</template>

<script>
import api from '@service/home'

export default {
  name: 'Index',
  methods: {
    async initData () {
      try {
        const data = await api.getCategoryList()
        console.log('data:', data)
      } catch (e) {
        console.log('e:', e)
      }
    }
  },
  mounted () {
    console.log('dddd:', this.initData().then(e => console.log('eeeee:', e)))
  }
}
</script>

<style lang="less" scoped>

</style>
