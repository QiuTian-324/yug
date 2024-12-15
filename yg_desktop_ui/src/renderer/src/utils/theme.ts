import { merge } from 'lodash'; // 确保安装了 lodash
import { genMixColor } from './gen-color';

// 主题配置类型定义
export type Theme = {
  colors: {
    primary: string;
    success: string;
    danger: string;
    info: string;
    warning: string;
  };
};

// 默认主题配置
export const defaultThemeConfig: Theme = {
  colors: {
    primary: '#FF6A00',
    info: '#eeeeee',
    warning: '#fbbd23',
    danger: '#f87272',
    success: '#36d399',
  },
};

// 本地缓存 key
const THEME_KEY = 'theme';

// 获取本地缓存主题
export const getTheme = (): Theme => {
  const theme = localStorage.getItem(THEME_KEY);
  return theme ? JSON.parse(theme) : defaultThemeConfig;
};

// 设置主题
export const setTheme = (data: Theme = defaultThemeConfig) => {
  const oldTheme = getTheme();

  // 将传入配置与旧的主题合并，以填补缺省的值
  data = merge(oldTheme, data || {});

  // 将缓存到浏览器
  localStorage.setItem(THEME_KEY, JSON.stringify(data));

  // 将主题更新到css变量中，使之生效
  updateThemeColorVar(data);
};

// 设置css变量
function setStyleProperty(propName: string, value: string) {
  document.documentElement.style.setProperty(propName, value);
}

// 更新主题色到css变量
function updateThemeColorVar({ colors }: Theme) {
  for (const brand in colors) {
    updateBrandExtendColorsVar(
      colors[brand as keyof Theme['colors']] as string,
      brand
    );
  }

  function updateBrandExtendColorsVar(color: string, name: string) {
    const { DEFAULT, dark, light } = genMixColor(color);
    setStyleProperty(`--${name}-lighter-color`, light[5]);
    setStyleProperty(`--${name}-light-color`, light[3]);
    setStyleProperty(`--${name}-color`, DEFAULT);
    setStyleProperty(`--${name}-deep-color`, dark[2]);
    setStyleProperty(`--${name}-deeper-color`, dark[4]);

    setStyleProperty(`--el-color-${name}`, DEFAULT);
    setStyleProperty(`--el-color-${name}-dark-2`, dark[2]);
    setStyleProperty(`--el-color-${name}-light-3`, light[3]);
    setStyleProperty(`--el-color-${name}-light-5`, light[5]);
    setStyleProperty(`--el-color-${name}-light-7`, light[7]);
    setStyleProperty(`--el-color-${name}-light-8`, light[8]);
    setStyleProperty(`--el-color-${name}-light-9`, light[9]);
  }
}

export const lightTheme: Theme = {
  colors: {
    primary: '#409EFF',
    success: '#67C23A',
    danger: '#F56C6C',
    info: '#909399',
    warning: '#E6A23C',
  },
};

export const darkTheme: Theme = {
  colors: {
    primary: '#1F2D3D',
    success: '#2E8B57',
    danger: '#8B0000',
    info: '#2F4F4F',
    warning: '#DAA520',
  },
};

export function applyTheme(theme: Theme) {
  const root = document.documentElement;
  Object.entries(theme.colors).forEach(([key, value]) => {
    root.style.setProperty(`--el-color-${key}`, value);
  });

  // 更新弹窗背景色
  const dialogBackgroundColor = theme === darkTheme ? '#333333' : '#ffffff'; // 根据主题设置颜色
  root.style.setProperty('--dialog-background-color', dialogBackgroundColor);
}