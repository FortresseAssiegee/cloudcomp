import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import {
  mapState
} from "vuex";
import axios from "axios";
import VueAxios from "vue-axios";
import vuetify from './plugins/vuetify'
import DashboardPlugin from "./plugins/dashboard-plugin";
import 'vuetify/dist/vuetify.min.css'

import Video from 'video.js'
import 'video.js/dist/video-js.css'

Vue.prototype.$video = Video

axios.defaults.baseURL = 'http://47.103.88.114:9999'
// let Config = {
//   TIMEOUT: 1500000,
//   // baseURL: {
//   //   dev: window.BASEURL_01,
//   //   prod: ''
//   // }
// };
// axios.defaults.timeout = Config.TIMEOUT;
// Vue.prototype.axios = axios;

// plugin setup
Vue.use(DashboardPlugin);
Vue.config.productionTip = false
Vue.use(VueAxios, axios);

new Vue({
  router,
  store,
  vuetify,
  mapState,
  render: h => h(App)
}).$mount('#app')