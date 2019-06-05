<template>
  <div class="gallery">
    <div class="swiper-container">
      <div class="swiper-wrapper">
        <div class="swiper-slide" v-for="(item, index) in urls" :key="index" @click="onClickThumb(index)">
          <j-image
            width="100%"
            height="100%"
            :class="{imgactive: currentIndex === index}"
            :src="item"></j-image>
        </div>
      </div>
      <!-- 如果需要导航按钮 -->
      <div class="gallery-button-prev">
        <Icon type="ios-arrow-up"></Icon>
      </div>
      <div class="gallery-button-next">
        <Icon type="ios-arrow-down"></Icon>
      </div>
    </div>
    <slot>
      <div class="gallery-bigimg-wrap">
        <j-image
          width="400px"
          height="400px"
          show-loading
          :src="currentImg"></j-image>
      </div>
    </slot>
    <div style="display: none;">
      <img :src="item" alt="预加载" v-for="(item, index) in preUrls" :key="index">
    </div>
  </div>
</template>
<style lang="less">
  @import "~swiper/dist/css/swiper.css";

  .gallery {
    display: inline-block;
  }

  .gallery-bigimg-wrap {
    display: inline-block;
    margin-left: 10px;
    vertical-align: top;
  }

  .imgactive {
    border-color: #ff9900 !important;
  }

  .swiper-container {
    display: inline-block;
    border: 1px solid #eee;
    width: 60px;
    height: 400px;
    padding: 20px 0;
  }

  .swiper-slide {
    cursor: pointer;
  }

  .gallery-button-prev,
  .gallery-button-next {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    left: 0;
    right: 0;
    height: 20px;
    background: #fff;
    z-index: 100;
    cursor: pointer;
  }

  .gallery-button-prev {
    top: 0;
  }

  .gallery-button-next {
    bottom: 0;
  }

  .swiper-button-disabled {
    opacity: .2;
  }
</style>
<script>
import Swiper from 'swiper/dist/js/swiper.min'
import Jimage from '../image'

export default {
  name: 'j-image-gallery',
  components: {
    [Jimage.name]: Jimage
  },
  props: {
    urls: {
      type: Array,
      default () {
        return []
      }
    },
    current: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      swiper: null,
      currentIndex: this.current
    }
  },
  watch: {
    urls: {
      handler () {
        this.$nextTick(() => {
          this.initSwiper()
        })
      },
      immediate: true
    }
  },
  computed: {
    preUrls () {
      return this.urls.map(item => item.slice(0, item.lastIndexOf('_')) + item.slice(item.lastIndexOf('.')))
    },
    currentImg () {
      return this.preUrls[this.currentIndex]
    }
  },
  methods: {
    onClickThumb (index) {
      this.currentIndex = index
    },
    initSwiper () {
      this.swiper = new Swiper('.swiper-container', {
        direction: 'vertical',
        slidesPerView: 5, // 一容器显示多少个
        spaceBetween: 10, // 每个之间的间距
        slidesPerGroup: 5, // 一次换页切换多少个单元
        navigation: {
          nextEl: '.gallery-button-next',
          prevEl: '.gallery-button-prev'
        }
      })
    }
  }
}
</script>
