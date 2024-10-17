import { createApp } from 'vue'
import './style.css'
import App from './App.vue';
import pinia from './modules/pinia'
import router from './modules/router'


import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import VueWechatTitle from 'vue-wechat-title'


const app = createApp(App)
app.use(router)
app.use(VueWechatTitle)

app.use(pinia)
app.use(ElementPlus)

app.mount('#app')
