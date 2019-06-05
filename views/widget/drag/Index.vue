<template>
  <div>
    <Drag v-model="items">
      <div class="list-item" v-for="item in items" :key="item.id">{{item.title}}</div>
      <button slot="footer" @click="onClick">添加</button>
    </Drag>
  </div>
</template>

<script>
import Drag from 'vuedraggable'

export default {
  name: 'Index',
  components: {
    Drag
  },
  data () {
    return {
      index: 10,
      items: [
        { id: '0', title: 'title1' },
        { id: '1', title: 'title2' },
        { id: '2', title: 'title3' },
        { id: '3', title: 'title4' }
      ]
    }
  },
  watch: {
    items (val) {
      console.log('val:', val)
    }
  },
  methods: {
    onClick () {
      this.items.push({ id: (this.index++) + '', title: this.index })
      console.log('items:', this.items)
    }
  }
}
</script>

<style lang="less" scoped>
  .list-item {

    height: 50px;
    width: 200px;
    border: 1px solid #eee;
    margin-bottom: 5px;
    background: aqua;
    cursor: pointer;
    border-radius: 5px;
  }
</style>
