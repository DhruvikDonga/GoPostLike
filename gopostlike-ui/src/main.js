import Vue from 'vue'
import App from './App.vue'
Vue.config.productionTip = false
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import router from './router'
import '@mdi/font/css/materialdesignicons.css'
// Add Axios
import axios from 'axios'

// Add Axios
Vue.prototype.$http = axios;
axios.defaults.headers.common = {
  'Token': `${localStorage.getItem("jwt")}`,

}

Vue.use(Vuetify)
export default new Vuetify({
  icons: {
    iconfont: 'mdi', // 'mdi' || 'mdiSvg' || 'md' || 'fa' || 'fa4' || 'faSvg'
  },
  theme: {
    dark: true
  },
})
new Vue({
  vuetify : new Vuetify(),
  router,
  axios,
  render: h => h(App)
}).$mount('#app')
