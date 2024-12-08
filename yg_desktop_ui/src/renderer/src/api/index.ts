import { newWarningMessage } from '@renderer/pkg/messages';
import router from '@renderer/router';
import axios from 'axios';
import i18n from '../i18n';
const { t } = i18n.global
/** 创建axios实例 */
const request = axios.create({
  baseURL: "http://127.0.0.1:9000/api",
  timeout: 6000
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 发请求带上token
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = "Bearer " + token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
request.interceptors.response.use(
  (resp: any) => {
    return resp.data
  },
  (error: any) => {
    if (error.response.status === 401) {
      localStorage.removeItem("token");
      newWarningMessage(t('login.loginFirst'));
      router.push({ name: "login" });
    }
    if (error.response.status === 429) {
      newWarningMessage(t('fail.requestTooFrequent'));
    }
    return Promise.reject(error.response.data)
  }
)

export default request

// 定义 API 响应的接口
export interface ApiResponse<T> {
  code: number;
  data: T;
  extra: any;
  message: string;
  success: boolean;
}


// GET 请求
export function Get<T>(url: string, params: Record<string, any> = {}): Promise<ApiResponse<T>> {
  return request.get(url, { params });
}

// POST 请求
export function Post<T>(url: string, data: Record<string, any> = {}): Promise<ApiResponse<T>> {
  console.log(url);
  return request.post(url, data);
}
