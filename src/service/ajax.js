/**
 * Created by xuwei on 2017/3/16.
 */
import axios from 'axios'
import { Message } from 'iview'
// import storage from '@util/storage'
const timeout = 60000

axios.defaults.baseURL = '/selection/api/'
axios.defaults.timeout = timeout
export default (
  {
    url,
    data = {},
    method = 'get',
    selfHandleMsg = false,
    timeout = 60000,
    config
  }) => {
  method = method.toLowerCase()

  // 向Promise对象中添加该参数用于ajax的取消操作
  const source = axios.CancelToken.source()
  let option = {
    url: url,
    method: method,
    cancelToken: source.token,
    timeout: timeout// 默认超时时间
  }
  // const token = storage.getToken()
  // if (token) {
  //   option.headers = {
  //     token: token
  //   }
  // }
  option[method === 'get' ? 'params' : 'data'] = data

  option = Object.assign({}, option, config)
  let promise = new Promise((resolve, reject) => {
    axios(option).then(({ data }) => {
      if (!selfHandleMsg) {
        if (data.code !== 0) {
          // 登录失效
          if (data.code === 403) {
            location.replace('/login')
          } else {
            Message.warning(data.msg)
          }
          reject(new Error('codeError'))
        } else {
          resolve(data)
        }
      } else {
        resolve(data)
      }
    }).catch(error => {
      // 手动取消请求
      if (axios.isCancel(error)) {
        console.log(error.message)
      } else {
        if (error.message === `timeout of ${timeout}ms exceeded`) {
          Message.warning('请求超时,请检查网络...')
        } else {
          Message.warning('网络异常，请重试...')
        }
      }
      reject(error)
    })
  })
  promise.source = source
  return promise
}
