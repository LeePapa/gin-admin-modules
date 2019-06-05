<template>
  <div :style="styles" ref="chart"></div>
</template>
<script>
import echarts from 'echarts'
import onlyLine from './onlyLine'
import lineOptions from './lineOptions'
import pieOptions from './pieOptions'
import merge from 'merge'
import debounce from 'lodash/debounce'
import { addListener, removeListener } from 'resize-detector/dist'

export default {
  name: 'HChart',
  props: {
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '300px'
    },
    windowResize: {
      type: Boolean,
      default: true
    },
    options: {
      type: Object,
      default: () => lineOptions
    },
    onlyLine: {
      type: Boolean,
      default: false
    },
    data: {
      type: Array,
      default: () => []
    },
    // 目前有line、pie、original(什么默认配置都不加)
    type: {
      type: String,
      default: 'line'
    }
  },
  data () {
    return {
      chart: null,
      chartData: {}
    }
  },
  computed: {
    styles () {
      return {
        width: this.width,
        height: this.height
      }
    }
  },
  watch: {
    options: {
      handler: function (val) {
        if (this.chart) {
          this.drawChart(val)
        } else {
          this.formatOptions(val)
        }
      },
      deep: true
    },
    data: {
      handler: function (val) {
        if (this.onlyLine) {
          if (this.chart) {
            this.drawChart(val)
          } else {
            this.formatOptions(val)
          }
        }
      },
      deep: true
    }
  },
  mounted () {
    this.$nextTick(() => {
      this.chart = echarts.init(this.$refs.chart)
      if (this.windowResize) {
        window.addEventListener('resize', this.debounceResize)
      }
      this.drawChart(this.options)
      this.autoResize()
    })
  },
  beforeDestroy () {
    removeListener(this.$el, this.__resizeHandler)
    window.removeEventListener('resize', this.debounceResize)
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    drawChart (options) {
      this.formatOptions(options)
      this.reDraw()
    },
    reDraw () {
      this.chart.setOption(this.chartData, true)
    },
    resize () {
      setTimeout(() => {
        this.chart.resize()
      }, 10)
    },
    debounceResize () {
      debounce(() => {
        this.resize()
      }, 150)()
    },
    formatOptions (opt) {
      if (this.onlyLine) {
        return this.onlyLineOptions()
      }
      if (opt) {
        switch (this.type) {
          case 'line':
            this.lineOptions(lineOptions, opt)
            break
          case 'pie':
            this.pieOptions(pieOptions, opt)
            break
          case 'original':
            this.originalOptions(opt)
            break
          default:
            this.lineOptions(lineOptions, opt)
            break
        }
      }
    },
    onlyLineOptions () {
      this.chartData = onlyLine.setData(onlyLine.base, this.data)
    },
    pieOptions (base, opt) {
      // 单个饼图的情况
      if (this.options.series && this.options.series.length === 1) {
        let series = [merge.recursive(true, base.series[0], opt.series[0])]
        let result = merge.recursive(true, base, opt)
        result.series = series
        this.chartData = result
      } else {
        this.chartData = merge.recursive(true, base, opt)
      }
    },
    lineOptions (base, opt) {
      this.chartData = merge.recursive(true, base, opt)
    },
    originalOptions (opt) {
      this.chartData = opt
    },
    // 当$el的长宽变化时，触发图表重绘
    autoResize () {
      this.lastArea = this.getArea()
      this.__resizeHandler = debounce(() => {
        if (this.lastArea === 0) {
          // emulate initial render for initially hidden charts
          this.resize()
        }
        this.lastArea = this.getArea()
      }, 100, { leading: true })
      addListener(this.$el, this.__resizeHandler)
    },
    getArea () {
      return this.$el.offsetWidth * this.$el.offsetHeight
    }
  }
}
</script>

<!--

通过侦听options或者data的变化来刷新图表，所以只需要修改对应传入的options就可以了

onlyLine: 为true时表示图仅仅显示折线图，坐标等任何信息都没有，主要用在table中，此时只需要传入data

data: 目前只用在onlyLine为true的情况

type: line(默认)、pie(目前是单个饼图)、original(什么配置都不加)

windowResize: 是否根据屏幕宽度变化伸缩图表，一般模态框中宽度固定不需要伸缩

-->
