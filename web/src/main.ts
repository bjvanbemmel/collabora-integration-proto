import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'

axios.defaults.baseURL = "http://192.168.50.216:88/wopi"

const app = createApp(App)

app.use(router)

app.mount('#app')
