import { createApp } from 'vue'
import vuetify from './plugins/vuetify'
import App from './App.vue'
import router from './router'

const app = createApp(App).use(router)
app.use(vuetify)

Vue.config.productionTip = false;
const base = axios.create({
  baseURL: "https://localhost:8000"
})
Vue.prototype.$http = base;
axios.defaults.baseURL = 'https://localhost:8000';

app.mount('#app')
