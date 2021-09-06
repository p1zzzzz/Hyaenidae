import Vue from 'vue'
import App from './App.vue'
// 引入elementui
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 引入mavonEditor
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
// 引入封装的router
import router from '@/router/index'
import { store } from '@/store'
import '@/permission'
import config from './config'
Vue.prototype.$Hyaenidae = config

import UUID from 'vue-uuid'
Vue.use(UUID)


// 路由守卫
import Bus from '@/utils/bus'
Vue.use(Bus)

// elementui
Vue.use(ElementUI);
// mavonEditor
Vue.use(mavonEditor)

export default new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')

