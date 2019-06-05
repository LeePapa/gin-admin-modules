<template>
  <div>
    <Row :gutter="24">
      <Col span="3">
        <HImage :src="src" width="80px" height="80px"></HImage>
        adapt
      </Col>
      <Col span="3">
        <HImage :src="src" type="scaleToFill" width="180px" height="80px"></HImage>
        scaleToFill
      </Col>
      <Col span="3">
        <HImage :src="src" type="widthFix" width="180px" height="80px"></HImage>
        widthFix
      </Col>
      <Col span="3">
        <HImage :src="src" type="heightFix" width="180px" height="80px"></HImage>
        heightFix
      </Col>
      <Col span="12">
        <ImageGallery :urls="urls"></ImageGallery>
      </Col>
    </Row>
  </div>
</template>

<script>
import HImage from '@components/image'
import ImageGallery from '@components/image-gallery'

export default {
  name: 'Index',
  components: {
    HImage,
    ImageGallery
  },
  data () {
    return {
      src: 'http://169.55.42.92:9999/group1/M03/EF/51/Co6qsVscxuyAHcuQAABTDcVCZU8830_60x60.jpg',
      urls: ['http://169.55.42.92:9999/group1/M03/EF/51/Co6qsVscxuyAHcuQAABTDcVCZU8830_60x60.jpg',
        'http://169.55.42.92:9999/group1/M03/EF/51/Co6qsVscxuyAHcuQAABTDcVCZU8830_60x60.jpg']
    }
  }
}
</script>

<style lang="less" scoped>

</style>
