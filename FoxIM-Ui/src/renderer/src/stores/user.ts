import { Post } from '@renderer/api'
import { newSuccessMessage } from '@renderer/pkg/messages'
import { defineStore } from 'pinia'
import { ref } from 'vue'
export const UserStore = defineStore('user', () => {
  const userInfo = ref({})
  const token = ref('')
  const addFriendModalVisible = ref(false)
  const addFriendData = ref<{ username: string }>({
    username: ''
  })
  // 修改数据的方法（action）
  const LoginIn = async () => {
    try {
      const res = await Post('/user/login', { username: 'akita1', password: '123' })
      token.value = res.data.token
      newSuccessMessage('登录成功')
    } catch (error) {
      console.log(error)
    }
  }
  // 添加好友
  const addFriend = async () => {
    addFriendModalVisible.value = true
    try {
      const res = await Post('/user/add_friend', addFriendData.value)
      console.log(res)
      newSuccessMessage('添加成功')
    } catch (error) {
      console.log(error)
    }
  }
  return {
    userInfo,
    token,
    addFriendModalVisible,
    addFriendData,
    LoginIn,
    addFriend
  }
}, {
  persist: {
    enabled: true,
    strategies: [
      {
        key: 'token',
        storage: localStorage,
        paths: ['token']
      }
    ]
  }
})
