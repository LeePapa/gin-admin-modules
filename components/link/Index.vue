<template>
  <div :class="classes">
    <a v-if="outLink || blank" :href="typeof to === 'string' ? to : to.path " :target="blank ? '_blank' : '_self'">
      <slot></slot>
    </a>
    <router-link :to="to" :replace="replace" v-else>
      <slot></slot>
    </router-link>
  </div>
</template>
<style lang="less">
  @text-color: #666;
  @link-color: rgba(16, 142, 233, 0.87);
  @link-hover-color: rgba(16, 142, 233, 0.87);
  @link-active-color: rgba(16, 142, 233, 0.87);
  .h-link {
    display: inline-block;
    a {
      display: inline-block;
      width: 100%;
      height: 100%;
      color: @link-color;
      text-decoration: none;
      &:link {
        color: @link-color;
      }
      &:hover {
        color: @link-hover-color;
      }
      &:active,
      &:visited,
      &:focus {
        color: @link-active-color;
      }
    }
    &-plain {
      a {
        color: @text-color;
        text-decoration: none;
        &:link,
        &:hover,
        &:active,
        &:visited,
        &:focus {
          color: @text-color;
        }
        &:hover {
          text-decoration: none;
        }
      }
    }
    &-underline {
      a {
        &:hover {
          text-decoration: underline !important;
        }
      }
    }
  }
</style>
<script>
export default {
  name: 'h-link',
  props: {
    to: {
      type: [Object, String],
      default () {
        return {}
      }
    },
    plain: {
      type: Boolean,
      default: false
    },
    underline: {
      type: Boolean,
      default: true
    },
    blank: {
      type: Boolean,
      default: false
    },
    replace: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    classes () {
      return [
        'h-link',
        { 'h-link-plain': this.plain },
        { 'h-link-underline': this.underline }
      ]
    },
    outLink () {
      let src = typeof this.to === 'string' ? this.to : (this.to.path || '')
      return src.startsWith('http')
    }
  }
}
</script>
