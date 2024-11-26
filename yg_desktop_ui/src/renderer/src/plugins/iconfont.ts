import { createFromIconfontCN } from '@ant-design/icons-vue';

export default function install(app) {
  const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/c/font_4759350_824idk90o7.js'
  });

  app.component('IconFont', IconFont);
}