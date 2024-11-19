import { ApiResponse, Get, Post } from '@renderer/api';
import { AddFriendRequest, FriendListResponse, LoginRequest, LoginResponse } from '@renderer/interface/user';
import { newErrorMessage, newSuccessMessage } from '@renderer/pkg/messages';
import { defineStore } from 'pinia';
import { ref } from 'vue';
export const UserStore = defineStore('user', () => {
  // 用户信息
  const userInfo = ref({})
  // token
  const token = ref(localStorage.getItem('token'))
  // 添加好友对话框是否显示
  const addFriendModalVisible = ref(false)
  // 添加好友数据
  const addFriendData = ref<AddFriendRequest>({
    username: ''
  })

  // 选中目标ID
  const selectedSeesionId = ref(0)

  // 选择当前侧边栏
  const selectedSidebar = ref(0)

  // 好友列表
  const friendList = ref<FriendListResponse[]>([])


  // 初始化数据
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
      return error
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
      return error
    }
  }

  // 获取好友列表
  const getUserInfo = async () => {
    try {
      const res: ApiResponse<FriendListResponse[]> = await Get<FriendListResponse[]>('/user/friends')
      friendList.value = res.data
    } catch (error: any) {
      return error
    }
  }

  return {
    userInfo,
    token,
    addFriendModalVisible,
    addFriendData,
    selectedSidebar,
    friendList,
    selectedSeesionId,
    LoginIn,
    addFriend,
    getUserInfo
  }
}
  , {
    persist: {
      enabled: true,
      strategies: [
        {
          key: 'selectedSidebar',
          storage: localStorage,
          paths: ['selectedSidebar'],
        },
        {
          key: 'friendList',
          storage: localStorage,
          paths: ['friendList'],
        }
      ]
    }
  }
)
