import {Menu, Icon, Input} from 'ant-design-vue'

const {Item, ItemGroup, SubMenu} = Menu
const {Search} = Input

export default {
  name: 'Tree',
  props: {
    dataSource: {
      type: Array,
      required: true
    },
    openKeys: {
      type: Array,
      default: () => []
    },
    search: {
      type: Boolean,
      default: false
    }
  },
  created() {
    this.localOpenKeys = this.openKeys.slice(0)
  },
  data() {
    return {
      localOpenKeys: []
    }
  },
  methods: {
    handlePlus(item) {
      this.$emit('add', item)
    },
    handleSet(key, item) {
      console.log(item)
      console.log(key)
    },
    handleTitleClick(...args) {
      this.$emit('titleClick', {args})
    },

    renderSearch() {
      return (
        <Search
          placeholder="input search text"
          style="width: 100%; margin-bottom: 1rem"
        />
      )
    },
    renderIcon(icon) {
      return icon && (<Icon type={icon}/>) || null
    },
    renderMenuItem(item) {
      console.log(item)
      return (
        <Item key={item.key}>
          {this.renderIcon(item.icon)}
          {item.title}
          <a-dropdown>
            <a class="btn">
              <a-icon type="ellipsis"/>
            </a>
            <a-menu slot="overlay" {...{on: {click: (key) => this.handleSet(key, item)}}}>
              <a-menu-item key="1">编辑</a-menu-item>
              {item.level < 3 ? <a-menu-item key="2">新增子菜单</a-menu-item> : ''}
              {item.sort !== 0 ? <a-menu-item key="3">上移</a-menu-item> : ''}
              {item.sort !== -1 ? <a-menu-item key="4">下移</a-menu-item> : ''}
            </a-menu>
          </a-dropdown>
        </Item>
      )
    },
    renderItem(item) {
      return item.children ? this.renderSubItem(item, item.key) : this.renderMenuItem(item, item.key)
    },
    renderItemGroup(item) {
      const childrenItems = item.children.map(o => {
        return this.renderItem(o, o.key)
      })

      return (
        <ItemGroup key={item.key}>
          <template slot="title">
            <span>{item.title}</span>
            <a-dropdown>
              <a class="btn">
                <a-icon type="ellipsis"/>
              </a>
              <a-menu slot="overlay">
                <a-menu-item key="1">编辑</a-menu-item>
                <a-menu-item key="2">新增子菜单</a-menu-item>
                {item.sort !== 0 ? <a-menu-item key="3">上移</a-menu-item> : ''}
                {item.sort !== -1 ? <a-menu-item key="4">下移</a-menu-item> : ''}
              </a-menu>
            </a-dropdown>
          </template>
          {childrenItems}
        </ItemGroup>
      )
    },
    renderSubItem(item, key) {
      const childrenItems = item.children && item.children.map(o => {
        return this.renderItem(o, o.key)
      })

      const title = (
        <span slot="title">
          {this.renderIcon(item.icon)}
          <span>{item.title}</span>
        </span>
      )

      if (item.group) {
        return this.renderItemGroup(item)
      }
      // titleClick={this.handleTitleClick(item)}
      return (
        <SubMenu key={key}>
          {title}
          {childrenItems}
        </SubMenu>
      )
    }
  },
  render() {
    const {dataSource, search} = this.$props

    // this.localOpenKeys = openKeys.slice(0)
    const list = dataSource.map(item => {
      return this.renderItem(item)
    })

    return (
      <div class="tree-wrapper">
        {search ? this.renderSearch() : null}
        <Menu mode="inline" class="custom-tree" {...{
          on: {
            click: item => this.$emit('click', item),
            'update:openKeys': val => {
              this.localOpenKeys = val
            }
          }
        }} openKeys={this.localOpenKeys}>
          {list}
        </Menu>
      </div>
    )
  }
}
