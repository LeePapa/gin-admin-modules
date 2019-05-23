<template>
  <a-card :bordered="false">
    <a-row :gutter="8">
      <a-col :span="6">
        <a-card>
          <a-button style="width: 250px;" slot="extra" @click="editMenuHandle" v-if="menuInfo.id == 0">添加菜单</a-button>
          <a-button-group style="width: 250px;" v-else slot="extra">
            <a-tooltip v-if="menuInfo.level <= 2" placement="top">
              <template slot="title">
                <span>添加子菜单</span>
              </template>
              <a-button @click="addSubMenuHandle" :style="getButtonClass()">
                <a-icon type="plus-circle" theme="twoTone"/>
              </a-button>
            </a-tooltip>
            <a-tooltip placement="top">
              <template slot="title">
                <span>编辑菜单</span>
              </template>
              <a-button @click="editMenuHandle" :style="getButtonClass">
                <a-icon type="edit" theme="twoTone"/>
              </a-button>
            </a-tooltip>
            <a-tooltip v-if="menuInfo.sort != 1 && menuInfo.sort != 0" placement="top">
              <template slot="title">
                <span>位置上移</span>
              </template>
              <a-button @click="setSortHandle('up')" :style="getButtonClass()">
                <a-icon type="up-circle" theme="twoTone"/>
              </a-button>
            </a-tooltip>
            <a-tooltip v-if="menuInfo.sort != -1 && menuInfo.sort != 0" placement="top">
              <template slot="title">
                <span>位置下移</span>
              </template>
              <a-button @click="setSortHandle('down')" :style="getButtonClass()">
                <a-icon type="down-circle" theme="twoTone"/>
              </a-button>
            </a-tooltip>

            <a-popconfirm placement="topLeft" okText="Yes" cancelText="No" @confirm="deleteMenuHandle">
              <template slot="title">
                <p>是否确定删除"{{menuInfo.name}}"该条菜单记录</p>
              </template>
              <a-tooltip placement="top">
                <template slot="title">
                  <span>删除菜单</span>
                </template>
                <a-button :style="getButtonClass()">
                  <a-icon type="delete" theme="twoTone" twoToneColor="#eb2f96"/>
                </a-button>
              </a-tooltip>
            </a-popconfirm>
          </a-button-group>
          <a-tree showIcon @select="menuSelectHandle">
            <a-tree-node v-if="treeM.children.length > 0" v-for="treeM in this.$store.state.menu.treeMenus"
                         :key="treeM.id.toString()" :title="treeM.title">
              <a-tree-node v-if="treeS.hasOwnProperty('children') && treeS.children.length > 0"
                           v-for="treeS in treeM.children" :title="treeS.title" :key="treeS.id.toString()">
                <a-tree-node v-for="treeI in treeS.children" :title="treeI.title" :key="treeI.id.toString()"/>
              </a-tree-node>
              <a-tree-node v-else :title="treeS.title" :key="treeS.id.toString()"/>
            </a-tree-node>
            <a-tree-node v-else :key="treeM.id.toString()" :title="treeM.title"/>
          </a-tree>
        </a-card>
      </a-col>
      <a-col :span="18" v-if="showMenuModal">
        <a-form :form="form" @submit="handleSubmit">
          <a-form-item v-if="menuInfo.id > 0" label="id" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-input :disabled="true" v-decorator="['id',{initialValue:menuInfo.id}]"/>
          </a-form-item>
          <a-form-item label="名称" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-input
              v-decorator="['name',{rules: [{ required: true, message: '请输入菜单名称' }],initialValue:menuInfo.name}]"/>
          </a-form-item>
          <a-form-item label="上级菜单" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-tree-select v-decorator="['pid',{initialValue:''+menuInfo.pid+''}]" @select="parentMenuSelectHandle">
              <a-tree-select-node value='0' title='无' key='0'/>
              <a-tree-select-node v-if="treeInfo.children.length > 0"
                                  v-for="treeInfo in this.$store.state.menu.treeMenus"
                                  :value="''+treeInfo.id+''"
                                  :title='treeInfo.title' :key='treeInfo.id'>
                <a-tree-select-node v-for="treeSubInfo in treeInfo.children" :value="''+treeSubInfo.id+''"
                                    :title='treeSubInfo.title' :key='treeSubInfo.id'/>
              </a-tree-select-node>
              <a-tree-select-node v-else :value="''+treeInfo.id+''" :title='treeInfo.title' :key='treeInfo.id'/>
            </a-tree-select>
          </a-form-item>
          <a-form-item label="菜单等级" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-select v-decorator="['level',{initialValue:menuInfo.level}]" disabled>
              <a-select-option v-for="levelInfo in levelList" :value="levelInfo.key">
                {{levelInfo.name}}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="icon图标" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-input-group compact>
              <a-input v-decorator="['icon',{initialValue:menuInfo.icon}]" style="width: 80%"/>
              <a-button @click="iconHelpHandle" style="width: 20%" icon="question-circle"/>
            </a-input-group>
          </a-form-item>
          <a-form-item label="唯一标识" :help="checkKeyInfo.msg"
                       :validate-status="checkKeyInfo.status" has-feedback :label-col="{ span: 5 }"
                       :wrapper-col="{ span: 6 }">
            <a-input @change="menuKeyChangeHandle"
                     v-decorator="['key',{rules: [{ required: true, message: '请输入菜单唯一标识，不可重复' }],initialValue:menuInfo.key}]"/>
          </a-form-item>
          <a-form-item label="路由地址" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-input
              v-decorator="['path',{rules: [{ required: true, message:'请输入菜单路由地址' }],initialValue:menuInfo.path}]"/>
          </a-form-item>
          <a-form-item label="是否显示" :label-col="{ span: 5 }" :wrapper-col="{ span: 6 }">
            <a-switch v-decorator="['status', { valuePropName:'checked',initialValue:menuInfo.status }]"/>
          </a-form-item>
          <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
            <a-button :loading="submitLoading" type="primary" html-type="submit">提交</a-button>
          </a-form-item>
        </a-form>
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
  import {getMenuTree, checkMenuKey, getMenuDetail, SaveMenu, SetMenuSort, DeleteMenu} from '@/api/system'
  import {mapActions} from 'vuex'
  import AButton from "ant-design-vue/es/button/button";

  export default {
    name: 'TreeList',
    components: {AButton},
    data() {
      return {
        formLayout: 'horizontal',
        form: this.$form.createForm(this),
        levelList: [{key: 1, name: "1级"}, {key: 2, name: "2级"}, {key: 3, name: "3级"}],
        menuInfo: {
          id: 0,
          pid: 0,
          name: "",
          level: 1,
          sort: 0,
          icon: "",
          key: "",
          path: "",
          status: true
        },
        originMenuInfo: {},
        showMenuModal: false,
        submitLoading: false,
        checkKeyInfo: {
          status: "",
          loadStatus: true,
          msg: "",
          taskId: 0,
        },
      }
    },
    created() {

    },
    methods: {
      ...mapActions(['ChangeMyMenu']),
      handleSubmit(e) {
        e.preventDefault();
        this.form.validateFields((err, values) => {
          if (!err) {
            this.submitLoading = true
            console.log('Received values of form: ', values);
            values.status = values.status ? 0 : 1
            SaveMenu(values).then(res => {
              this.submitLoading = false
              if (res.code == 200) {
                this.$message.success(res.msg);
                this.checkKeyInfo.status = ""
                this.showMenuModal = false
                this.form.resetFields()
                this.ChangeMyMenu()
              } else {
                this.$message.error(res.msg);
              }
            })
          }
        });
      },
      addSubMenuHandle() {
        this.menuInfo = {
          id: 0,
          pid: this.menuInfo.id,
          name: "",
          level: this.menuInfo.level + 1,
          sort: 0,
          icon: "",
          key: "",
          path: "",
          status: true
        }
        this.showMenuModal = true
      },
      editMenuHandle() {
        this.showMenuModal = true
      },
      setSortHandle(type) {
        SetMenuSort({id: this.menuInfo.id, type: type}).then(res => {
          if (res.code == 200) {
            this.$message.success(res.msg);
            this.ChangeMyMenu()
          } else {
            this.$message.warn(res.msg);
          }
        })
      },
      deleteMenuHandle() {
        DeleteMenu({id: this.menuInfo.id}).then(res => {
          if (res.code == 200) {
            this.$message.success(res.msg);
            this.ChangeMyMenu()
            this.showMenuModal = false
            this.menuInfo = {
              id: 0,
              pid: 0,
              name: "",
              level: 1,
              sort: 0,
              icon: "",
              key: "",
              path: "",
              status: true
            }
          } else {
            this.$message.warn(res.msg);
          }
        })
      },
      parentMenuSelectHandle(value) {
        if (value == 0) {
          this.form.setFieldsValue({level: 1})
          return
        }
        getMenuDetail({id: value}).then(res => {
          if (res.code == 200) {
            this.form.setFieldsValue({level: res.data.level + 1})
          }
        })
      },
      iconHelpHandle() {
        window.open("https://vue.ant.design/components/icon-cn/");
      },
      menuSelectHandle(value) {
        this.showMenuModal = false
        getMenuDetail({id: value[0]}).then(res => {
          if (res.code == 200) {
            if (res.data.status == 0) {
              res.data.status = true
            } else {
              res.data.status = false
            }
            this.menuInfo = res.data
            console.log(this.menuInfo)
          }
        })
      },
      getButtonClass() {
        let spanWidth = 5
        if (this.menuInfo.id <= 0) {
          return {width: "20%"}
        }
        if (this.menuInfo.level > 2) {
          spanWidth -= 1;
        }
        if (this.menuInfo.sort == 1 || this.menuInfo.sort == -1) {
          spanWidth -= 1;
        }
        if (this.menuInfo.sort == 0) {
          spanWidth -= 2
        }
        return {width: parseInt(100 / spanWidth) + "%"}
      },
      menuKeyChangeHandle() {
        if (this.menuInfo.id > 0) {
          return
        }
        let that = this
        if (that.checkKeyInfo.taskId > 0) {
          clearTimeout(that.checkKeyInfo.taskId)
        }
        that.checkKeyInfo.status = "validating"
        that.checkKeyInfo.msg = ""
        that.checkKeyInfo.taskId = setTimeout(function () {
          checkMenuKey({key: that.form.getFieldValue("key")}).then(res => {
            if (res.code != 200) {
              that.checkKeyInfo.msg = res.msg
              that.checkKeyInfo.status = "error"
            } else {
              that.checkKeyInfo.status = "success"
            }
          })
        }, 1300)
      },
    }
  }
</script>

<style lang="less">
  .custom-tree {

    /deep/ .ant-menu-item-group-title {
      position: relative;
      &:hover {
        .btn {
          display: block;
        }
      }
    }

    /deep/ .ant-menu-item {
      &:hover {
        .btn {
          display: block;
        }
      }
    }

    /deep/ .btn {
      display: none;
      position: absolute;
      top: 0;
      right: 10px;
      width: 20px;
      height: 40px;
      line-height: 40px;
      z-index: 1050;

      &:hover {
        transform: scale(1.2);
        transition: 0.5s all;
      }
    }
  }
</style>
