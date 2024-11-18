import { ApiResponse, Get, Post } from '@renderer/api';
import { AddFriendRequest, FriendListResponse, LoginRequest, LoginResponse } from '@renderer/interface/user';
import { newErrorMessage, newSuccessMessage } from '@renderer/pkg/messages';
import { defineStore } from 'pinia';
import { ref } from 'vue';
export const UserStore = defineStore('user', () => {
  const userInfo = ref({})
  const token = ref(localStorage.getItem('token'))
  const addFriendModalVisible = ref(false)
  const addFriendData = ref<AddFriendRequest>({
    username: ''
  })

  const init = () => {
    getUserInfo()
  }

  // 修改数据的方法（action）
  const LoginIn = async () => {
    try {
      const data: LoginRequest = { username: 'akita1', password: '123' }
      const res: ApiResponse<LoginResponse> = await Post<LoginResponse>('/user/login', data);
      token.value = res.data.token
      console.log(res)
      localStorage.setItem('token', res.data.token)
      init()
      newSuccessMessage(res.message)
    } catch (error: any) {
      newErrorMessage(error.message)
    }
  }
  // 添加好友
  const addFriend = async () => {
    addFriendModalVisible.value = true
    try {
      const res: ApiResponse<any> = await Post<ApiResponse<any>>('/user/add_friend', addFriendData.value)
      newSuccessMessage(res.message)
    } catch (error: any) {
      newErrorMessage(error.message)
    }
  }

  // 获取好友列表
  const getUserInfo = async () => {
    try {
      const res: ApiResponse<FriendListResponse> = await Get<FriendListResponse>('/user/friends')
      console.log(res)
    } catch (error: any) {
      return error
    }
  }

  return {
    userInfo,
    token,
    addFriendModalVisible,
    addFriendData,
    LoginIn,
    addFriend,
    getUserInfo
  }
}
  // , {
  //   persist: {
  //     enabled: true,
  //     strategies: [
  //       {
  //         key: 'token',
  //         storage: localStorage,
  //         paths: ['token'],

  //       }
  //     ]
  //   }
  // }
)
