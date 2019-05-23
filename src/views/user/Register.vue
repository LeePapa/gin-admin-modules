<template>
  <div class="main user-layout-register">
    <h3><span>注册</span></h3>
    <a-form ref="formRegister" :autoFormCreate="(form)=>{this.form = form}" id="formRegister">
      <a-form-item
        fieldDecoratorId="mobile"
        :fieldDecoratorOptions="{rules: [{ required: true, message: '请输入正确的手机号', pattern: /^1[3456789]\d{9}$/ }, { validator: this.handlePhoneCheck } ], validateTrigger: ['change', 'blur'] }">
        <a-input size="large" placeholder="11 位手机号">
          <a-select slot="addonBefore" size="large" defaultValue="+86">
            <a-select-option value="+86">+86</a-select-option>
          </a-select>
        </a-input>
      </a-form-item>
      <a-row :gutter="16">
        <a-col class="gutter-row" :span="16">
          <a-form-item
            fieldDecoratorId="code"
            :fieldDecoratorOptions="{rules: [{ required: true, message: '请输入验证码' }], validateTrigger: 'blur'}">
            <a-input size="large" type="text" placeholder="验证码">
              <a-icon slot="prefix" type="mail" :style="{ color: 'rgba(0,0,0,.25)' }"/>
            </a-input>
          </a-form-item>
        </a-col>
        <a-col class="gutter-row" :span="8">
          <a-button
            class="getCaptcha"
            size="large"
            :disabled="state.smsSendBtn"
            @click.stop.prevent="getCaptcha"
            v-text="!state.smsSendBtn && '获取验证码'||(state.time+' s')"></a-button>
        </a-col>
      </a-row>
      <a-popover placement="rightTop" trigger="click" :visible="state.passwordLevelChecked">
        <template slot="content">
          <div :style="{ width: '240px' }">
            <div :class="['user-register', passwordLevelClass]">强度：<span>{{ passwordLevelName }}</span></div>
            <a-progress :percent="state.percent" :showInfo="false" :strokeColor=" passwordLevelColor "/>
            <div style="margin-top: 10px;">
              <span>请至少输入 6 个字符。请不要使用容易被猜到的密码。</span>
            </div>
          </div>
        </template>
        <a-form-item
          fieldDecoratorId="password"
          :fieldDecoratorOptions="{rules: [{ required: true, message: '至少6位密码，区分大小写'}, { validator: this.handlePasswordLevel }], validateTrigger: ['change', 'blur']}">
          <a-input size="large" type="password" @click="handlePasswordInputClick" autocomplete="false"
                   placeholder="至少6位密码，区分大小写"></a-input>
        </a-form-item>
      </a-popover>

      <a-form-item
        fieldDecoratorId="password2"
        :fieldDecoratorOptions="{rules: [{ required: true, message: '至少6位密码，区分大小写' }, { validator: this.handlePasswordCheck }], validateTrigger: ['change', 'blur']}">

        <a-input size="large" type="password" autocomplete="false" placeholder="确认密码"></a-input>
      </a-form-item>

      <a-form-item>
        <a-button
          size="large"
          type="primary"
          htmlType="submit"
          class="register-button"
          :loading="registerBtn"
          @click.stop.prevent="handleSubmit"
          :disabled="registerBtn">注册
        </a-button>
        <router-link class="login" :to="{ name: 'login' }">使用已有账户登录</router-link>
      </a-form-item>
    </a-form>
    <vueTencentCaptcha appid="2071852665" @callback="captchaCallback"/>
  </div>
</template>

<script>
  import md5 from 'md5'
  import {mixinDevice} from '@/utils/mixin.js'
  import {getSmsCaptcha, register} from '@/api/login'

  const levelNames = {
    0: '低',
    1: '低',
    2: '中',
    3: '强'
  }
  const levelClass = {
    0: 'error',
    1: 'error',
    2: 'warning',
    3: 'success'
  }
  const levelColor = {
    0: '#ff0000',
    1: '#ff0000',
    2: '#ff7e05',
    3: '#52c41a'
  }
  export default {
    name: 'Register',
    components: {},
    mixins: [mixinDevice],
    data() {
      return {
        form: null,
        captchaCallFunc: null,
        state: {
          time: 60,
          smsSendBtn: false,
          passwordLevel: 0,
          passwordLevelChecked: false,
          percent: 10,
          progressColor: '#FF0000'
        },
        registerBtn: false
      }
    },
    computed: {
      passwordLevelClass() {
        return levelClass[this.state.passwordLevel]
      },
      passwordLevelName() {
        return levelNames[this.state.passwordLevel]
      },
      passwordLevelColor() {
        return levelColor[this.state.passwordLevel]
      }
    },
    methods: {

      handlePasswordLevel(rule, value, callback) {
        let level = 0

        // 判断这个字符串中有没有数字
        if (/[0-9]/.test(value)) {
          level++
        }
        // 判断字符串中有没有字母
        if (/[a-zA-Z]/.test(value)) {
          level++
        }
        // 判断字符串中有没有特殊符号
        if (/[^0-9a-zA-Z_]/.test(value)) {
          level++
        }
        this.state.passwordLevel = level
        this.state.percent = level * 30
        if (level >= 2) {
          if (level >= 3) {
            this.state.percent = 100
          }
          callback()
        } else {
          if (level === 0) {
            this.state.percent = 10
          }
          callback(new Error('密码强度不够'))
        }
      },

      handlePasswordCheck(rule, value, callback) {
        const password = this.form.getFieldValue('password')
        console.log('value', value)
        if (value === undefined) {
          callback(new Error('请输入密码'))
        }
        if (value && password && value.trim() !== password.trim()) {
          callback(new Error('两次密码不一致'))
        }
        callback()
      },

      handlePhoneCheck(rule, value, callback) {
        console.log('handlePhoneCheck, rule:', rule)
        console.log('handlePhoneCheck, value', value)
        console.log('handlePhoneCheck, callback', callback)

        callback()
      },

      handlePasswordInputClick() {
        if (!this.isMobile()) {
          this.state.passwordLevelChecked = true
          return
        }
        this.state.passwordLevelChecked = false
      },

      handleSubmit() {
        const that = this
        this.form.validateFields((err, values) => {
          if (!err) {
            that.state.passwordLevelChecked = false
            register({username: values.mobile, code: values.code, password: md5(values.password)}).then(res => {
              if (res.code !== 200) {
                that.requestFailed(res)
                return
              }
              that.$message.success(res.msg + ", 即将跳转登录", 2, () => {
                that.$router.push({name: 'login'})
              })
            }).catch(err => {
              that.requestFailed(err)
            })
          }
        })
      },

      getCaptcha(e) {
        e.preventDefault()
        const that = this
        this.captchaCallFunc = null
        this.form.validateFields(['mobile'], {force: true},
          (err, values) => {
            if (!err) {
              that.$root.captcha.show();
              that.captchaCallFunc = function (ticket, randstr) {
                this.state.smsSendBtn = true
                const interval = window.setInterval(() => {
                  if (that.state.time-- <= 0) {
                    that.state.time = 60
                    that.state.smsSendBtn = false
                    window.clearInterval(interval)
                  }
                }, 1000)
                const getFail = function (res) {
                  clearInterval(interval)
                  that.state.time = 60
                  that.state.smsSendBtn = false
                  that.requestFailed(res)
                }
                getSmsCaptcha({mobile: values.mobile, type: 1, ticket: ticket, randStr: randstr}).then(res => {
                  if (res.code !== 200) {
                    getFail(res)
                    return
                  }
                  this.$message.success(res.msg)
                }).catch(err => {
                  getFail(err)
                })
              }
            }
          }
        )
      },
      captchaCallback(r) {
        console.log(r)
        if (r.ticket) {
          this.captchaCallFunc(r.ticket, r.randstr)
        }

      },
      requestFailed(err) {
        this.$message.config({
          top: `30px`,
          duration: 3,
          maxCount: 1,
        })
        this.$message.error(err.msg || "请求异常")
        this.registerBtn = false
      }
    },
  }
</script>
<style lang="less">
  .user-register {

    &.error {
      color: #ff0000;
    }

    &.warning {
      color: #ff7e05;
    }

    &.success {
      color: #52c41a;
    }

  }

  .user-layout-register {
    .ant-input-group-addon:first-child {
      background-color: #fff;
    }
  }
</style>
<style lang="less" scoped>
  .user-layout-register {

    & > h3 {
      font-size: 16px;
      margin-bottom: 20px;
    }

    .getCaptcha {
      display: block;
      width: 100%;
      height: 40px;
    }

    .register-button {
      width: 50%;
    }

    .login {
      float: right;
      line-height: 40px;
    }
  }
</style>
