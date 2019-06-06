<style lang="less" scoped>
  @import "~@base/fn"; // #1c2438
  .bg {
    padding-top: 10px;
    height: calc(~'100vh - 180px');
    background: transparent url("./bg.jpg") no-repeat center center;
    background-size: cover;
  }

  .login {
    float: right;
    margin-right: 50px;
    width: 400px;
    height: 340px;
    border-radius: 6px;
    border: 1px solid @border-color;
    box-shadow: 0 0 6px rgba(0, 0, 0, .2);
    background-color: #fff;
    padding: 40px 30px 0;
    margin-top: 10vh;
    &-title {
      color: #000;
      font-size: 16px;
      margin-bottom: 20px;
    }
  }

  .form-item {
    margin-bottom: 20px;
  }

  .register {
    text-align: center;
    font-size: 12px;
    margin-top: 10px;
  }
</style>
<template>
  <div class="bg clearfix" @keyup.enter="onLogin">
    <div class="login">
      <div class="login-title">
        {{$route.meta.title}}
      </div>
      <div class="form-item">
        <Input v-model="userName" placeholder="请输入用户名" size="large">
        <Button slot="prepend" icon="ios-person"></Button>
        </Input>
      </div>
      <div class="form-item">
        <Input v-model="pwd" type="password" placeholder="请输入密码" size="large">
        <Button slot="prepend" icon="ios-lock"></Button>
        </Input>
      </div>
      <div class="form-item">
        <Row type="flex" justify="space-between" align="middle">
          <Col>
            <Checkbox v-model="rememberMe">近30天自动登录</Checkbox>
          </Col>
          <Col style="font-size: 12px;line-height: 1;">
            <!--<router-link :to="{path: '/forgetPwd'}">忘记密码?</router-link>-->
          </Col>
        </Row>
      </div>
      <Button type="primary" size="large" @click="onLogin" long>登录</Button>
      <div class="register">
        <!--<router-link :to="{path: '/register'}">注册</router-link>-->
      </div>
    </div>
  </div>
</template>
<script>
// import api from '@service/system'
import storage from '@util/storage'

export default {
  data () {
    return {
      userName: '',
      pwd: '',
      rememberMe: true
    }
  },
  mounted () {
    // account localStorage作为存储记住账户密码的
    let account = storage.getAccount()
    if (account && account.rememberMe) {
      this.userName = account.userName
      this.pwd = account.pwd
      this.rememberMe = account.rememberMe
    }
    // username localStorage作为存储当前用户是否登录过了，只要不点退出按钮就一直处于登录状态，哪怕是重新打开浏览器
    // 这块还是考虑注释掉，由服务器端控制
    if (storage.getUserName()) {
      this.$store.commit('updateUserName', { userName: storage.getUserName() })
      this.$router.replace('/')
    }
  },
  methods: {
    onLogin () {
      let userName = this.userName.trim()
      if (!userName || !this.pwd) {
        return this.$Message.warning('请输入账号和密码')
      }
      storage.setUserName(userName)
      this.$store.commit('updateUserName', { userName })
      this.$router.replace('/')
      // api.loginIn({ phone: userName, password: this.pwd, type: 'password' })
      //   .then(({ data }) => {
      //     if (this.rememberMe) {
      //       storage.setAccount(JSON.stringify({
      //         userName: this.userName,
      //         pwd: this.pwd,
      //         rememberMe: this.rememberMe
      //       }))
      //     } else {
      //       storage.setAccount()
      //     }
      //     storage.setUserName(userName)
      //     storage.setToken(data.token)
      //     storage.setUserInfo(data)
      //     this.$router.replace('/')
      //   })
    }
  }
}
</script>
