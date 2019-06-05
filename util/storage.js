/**
 * Created by xuwei on 2017/7/12.
 * 所有涉及到浏览器本地存储的都统一放在此处管理，防止key的冲突
 */
export default {
  keys: {
    userId: 'userId',
    localCate: 'localCate',
    account: 'account',
    userName: 'userName',
    token: 'loginToken'
  },
  getUserId () {
    return this.getLocal(this.keys.userId)
  },
  setUserId (val) {
    this.setLocal(this.keys.userId, val)
  },
  getLocalCate () {
    return this.getLocal(this.keys.localCate)
  },
  setLocalCate (val) {
    this.setLocal(this.keys.localCate, val)
  },
  getAccount () {
    return this.getLocal(this.keys.account)
  },
  setAccount (val) {
    this.setLocal(this.keys.account, val)
  },
  getUserName () {
    return this.getLocal(this.keys.userName)
  },
  setUserName (val) {
    this.setLocal(this.keys.userName, val)
  },
  getToken () {
    return this.getLocal(this.keys.token)
  },
  setToken (val) {
    this.setLocal(this.keys.token, val)
  },
  getLocal (key) {
    let result = null
    try {
      result = JSON.parse(localStorage.getItem(key))
    } catch (e) {
      result = localStorage.getItem(key)
    }
    return result
  },
  setLocal (key, val) {
    if (!val) {
      localStorage.removeItem(key)
    } else {
      if (typeof val === 'string') {
        localStorage.setItem(key, val)
      } else {
        localStorage.setItem(key, JSON.stringify(val))
      }
    }
  }
}
