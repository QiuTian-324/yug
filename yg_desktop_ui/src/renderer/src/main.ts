import './assets/main.css';

import '@ant-design/cssinjs';
import * as Icons from "@ant-design/icons-vue";
import Antd from 'ant-design-vue';
import { createPinia } from 'pinia';
import piniaPluginPersist from 'pinia-plugin-persist';
import { createApp } from 'vue';
import App from './App.vue';
import './assets/css/tailwind.css'; // 引入css文件
import i18n from './i18n';
import iconfont from './plugins/iconfont';
import router from './router';

const pinia = createPinia()
const app = createApp(App);

const icons: any = Icons;


pinia.use(piniaPluginPersist);



app.
  use(Antd).
  use(router).
  use(pinia).
  use(i18n).
  use(iconfont)

app.mount('#app');

for (const i in icons) {
  app.component(i, icons[i]);
}
