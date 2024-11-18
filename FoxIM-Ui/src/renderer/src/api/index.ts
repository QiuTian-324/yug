import { newErrorMessage, newSuccessMessage } from '@renderer/pkg/messages';
import axios from 'axios';
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
      config.headers.Authorization = token;
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
    newSuccessMessage(resp.data.message)
    return resp.data
  },
  (error) => {
    newErrorMessage(error.response.data.message)
    return Promise.reject(error.response.data)
  }
)

export default request

// GET 请求
export function Get(url, params = {}) {
  return request.get(url, { params });
}

// POST 请求
export function Post(url, data = {}) {
  console.log(url)
  return request.post(url, data);
}