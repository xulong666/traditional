import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import './assets/css/global.css'
// 导入axios 设定基路由 将axios赋值给属性http
import axios from 'axios'

axios.defaults.baseURL = 'http://127.0.0.1:8000'
// 在挂载到原型对象之前， 先配置一下拦截， 增加token验证
// request就是一个请求拦截器， use是配置拦截之后的回调函数
axios.interceptors.request.use(config => {
  config.headers.Authorization = window.sessionStorage.getItem('token')
  // 最后一定return config, 因为这个拦截是在发送请求之前的， 所以需要返回， 进行请求
  return config
})
Vue.prototype.$http = axios

Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
