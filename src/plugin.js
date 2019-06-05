/**
 * Created by xuwei on 2018/4/11.
 */
// import iView from 'iview'
import accounting from 'accounting'
// 最好不要用这个bus，vuex的简单使用就可以
import bus from './bus'
import axios from 'axios'
import {
  Table, Page, Row, Col, Menu, MenuItem, Submenu, Icon, DropdownItem,
  Dropdown, DropdownMenu, Button, Card, Tabs, TabPane, Input, InputNumber, Tag, Select, Option
} from 'iview'

const jolly = {
  Table,
  Page,
  Row,
  Col,
  Menu,
  MenuItem,
  Submenu,
  Icon,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  Button,
  Card,
  Tabs,
  TabPane,
  Input,
  InputNumber,
  Tag,
  'i-select': Select,
  'i-option': Option
}

const doPrototype = (vue) => {
  accounting.settings = {
    currency: {
      symbol: '', // default currency symbol is '$'
      format: '%s%v', // controls output: %s = symbol, %v = value/number (can be object: see below)
      decimal: '.', // decimal point separator
      thousand: ',', // thousands separator
      precision: 2 // decimal places
    },
    number: {
      precision: 0, // default precision on numbers is 0
      thousand: ',',
      decimal: '.'
    }
  }
  vue.accounting = accounting
  vue.formatNumber = accounting.formatNumber
  vue.formatMoney = accounting.formatMoney
  vue.toFixed = accounting.toFixed
  vue.unformat = accounting.unformat
  vue.axios = axios
}

const install = (Vue, opts = {}) => {
  if (install.installed) return
  Vue.prototype.$bus = bus
  // Vue.use(iView)
  Object.keys(jolly).forEach(key => {
    Vue.component(key, jolly[key])
  })
  doPrototype(Vue.prototype)
}

export default {
  install
}
