import { defineStore } from 'pinia';
import { ref } from 'vue';
import { ApiResponse, Get, Post } from '../api';
import i18n from "../i18n/index";
import { AddFriendRequest, FriendListResponse, LoginRequest, LoginResponse } from '../interface/user';
import { newErrorMessage, newSuccessMessage } from '../pkg/messages';
import router from '../router';
import { useChatStore } from './chat';

export const UserStore = defineStore('user', () => {
  const { t } = i18n.global
  const chatStore = useChatStore()
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
  // 添加好友表单
  const loginFormState = ref<LoginRequest>({
    username: 'akita1',
    password: '123'
  })

  // 选中目标ID
  const selectedSeesionId = ref(0)

  // 选择当前侧边栏
  const selectedSidebar = ref(0)

  // 好友列表
  const friendList = ref<FriendListResponse[]>([])


  // 初始化数据
  const Init = () => {
    if (!token.value) {
      newErrorMessage(t('login.loginFirst'))
      router.push('/login')
      return
    }
    getFriendList()
    // chatStore.connectWebSocket(token.value)
  }

  // 登录
  const LoginIn = async () => {
    try {
      const res: ApiResponse<LoginResponse> = await Post<LoginResponse>('/user/login', loginFormState.value);
      token.value = res.data.token
      localStorage.setItem('token', res.data.token)
      router.push('/')
      newSuccessMessage(t('success.login'))
    } catch (error: any) {
      newErrorMessage(t('fail.login'))
      return error
    }
  }
  // 添加好友
  const addFriend = async () => {
    addFriendModalVisible.value = true
    try {
      await Post<ApiResponse<any>>('/user/add_friend', addFriendData.value)
      newSuccessMessage(t('success.add_friend'))
    } catch (error: any) {
      newErrorMessage(t('fail.add_friend'))
      return error
    }
  }

  // 退出登录
  const Logout = async () => {
    try {
      await Post<ApiResponse<any>>('/user/logout')
      localStorage.removeItem('token')
      router.push('/login')
      newSuccessMessage(t('success.logout'))
    } catch (error: any) {
      newErrorMessage(t('fail.logout'))
      return error
    }
  }

  // 获取好友列表
  const getFriendList = async () => {
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
    loginFormState,
    addFriendData,
    selectedSidebar,
    friendList,
    selectedSeesionId,
    Logout,
    LoginIn,
    addFriend,
    getFriendList,
    Init
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

