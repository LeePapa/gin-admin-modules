<template>
  <div>
    <Editor v-model="content" style="width: 800px;" ref="editor"></Editor>
    <Button type="primary" @click.native="onClick">输出</Button>
  </div>
</template>

<script>
import Editor from '@components/vue-editor'

export default {
  name: 'Index',
  components: { Editor },
  data () {
    return {
      content: ''
    }
  },
  created () {
    setTimeout(() => {
      this.content = 'ttttt'
      this.$refs.editor.setHtml('ttttt')
    }, 3000)
  },
  methods: {
    onClick () {
      console.log(this.content)
    }
  }
}
</script>

<style lang="less" scoped>

</style>
