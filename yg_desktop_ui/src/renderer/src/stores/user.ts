import { defineStore } from 'pinia';
import { ref } from 'vue';
import { ApiResponse, Get, Post } from '../api';
import i18n from "../i18n/index";
import { AddFriendRequest, FriendListResponse, LoginRequest, LoginResponse, UserInfoResponse } from '../interface/user';
import { newErrorMessage, newSuccessMessage } from '../pkg/messages';
import router from '../router';
import { useChatStore } from './chat';

export const UserStore = defineStore('user', () => {
  const { t } = i18n.global
  const chatStore = useChatStore()
  // 用户信息
  const userInfo = ref<UserInfoResponse | null>(null)
  // token
  const token = ref(localStorage.getItem('token'))
  // 添加好友对话框是否显示
  const addFriendModalVisible = ref(false)
  // 添加好友数据
  const addFriendData = ref<AddFriendRequest>({
    username: "12312"
  })
  // 添加好友表单
  const loginFormState = ref<LoginRequest>({
    username: 'akita1',
    password: '123'
  })

  // 是否开启黑暗模式
  const isDarkMode = ref(false)

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
    chatStore.connectWebSocket(token.value)
  }

  // 切换黑暗模式
  const toggleDarkMode = () => {
    isDarkMode.value = !isDarkMode.value
    document.documentElement.classList.toggle('dark', isDarkMode.value)
  }

  // github登录
  const GithubLogin = async () => {
    const returnUrl = Math.random().toString(36).substring(2) + new Date().getTime().toString();

    const res: ApiResponse<any> = await Post<any>('/user/github/authURL', { state: returnUrl })

    // 重定向到授权 URL
    window.location.href = res.data.url;
  }

  // 通过github code登录
  const LoginByGithubCode = async (code: string | null) => {
    try {
      const res: ApiResponse<any> = await Get<any>('/user/github/callback', { code })
      token.value = res.data.token
      localStorage.setItem('token', res.data.token)
      router.push('/')
      newSuccessMessage(t('success.login'))
    } catch (error: any) {
      newErrorMessage(t('fail.login'))
      return error
    }
  }


  // 登录
  const LoginIn = async () => {
    try {
      const res: ApiResponse<LoginResponse> = await Post<LoginResponse>('/user/login', loginFormState.value);
      token.value = res.extra.token
      localStorage.setItem('token', res.extra.token)
      router.push('/')
      newSuccessMessage(t('success.login'))
      getUserInfo(res.data.user_id)
    } catch (error: any) {
      newErrorMessage(t('fail.login'))
      return error
    }
  }

  // 获取用户信息
  const getUserInfo = async (user_id: string) => {
    const res: ApiResponse<UserInfoResponse> = await Get<UserInfoResponse>('/user/info', { user_id })
    userInfo.value = res.data
  }

  // 添加好友
  const addFriend = async () => {
    addFriendModalVisible.value = true
    try {
      await Post<ApiResponse<any>>('/user/add_friend', { username: addFriendData.value.username })
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
    isDarkMode,
    Logout,
    LoginIn,
    addFriend,
    getFriendList,
    Init,
    GithubLogin,
    LoginByGithubCode,
    toggleDarkMode,
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
        },
        {
          key: 'userInfo',
          storage: localStorage,
          paths: ['userInfo'],
        }
      ]
    }
  }
)

