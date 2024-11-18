import './assets/main.css';

import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import { createPinia } from 'pinia';
import piniaPluginPersist from 'pinia-plugin-persist';
import { createApp } from 'vue';
// import MessageBox from "../src/pkg/messages";
import App from './App.vue';
import './assets/css/tailwind.css'; // 引入css文件
import router from './router';

const pinia = createPinia()
const app = createApp(App);

pinia.use(piniaPluginPersist);

// 挂载到vue原型上
// app.config.globalProperties.$message = MessageBox

app.
  use(ArcoVue).
  use(router).
  use(ArcoVueIcon).
  use(pinia);

app.mount('#app');
