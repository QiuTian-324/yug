<template>
  <div class="h-screen w-screen flex">
    <!-- 左侧轮播 -->
    <div class="flex-1 overflow-hidden relative h-screen">
      <a-carousel autoplay effect="fade">
        <div v-for="(slide, index) in slides" :key="index" class="h-screen bg-white">
          <div class="flex justify-center items-center h-full">
            <img :src="slide.image" alt="" class="w-[70%] h-full" />
          </div>
        </div>
      </a-carousel>
    </div>
    <!-- 右侧登录表单 -->
    <div class="bg-white flex-1 flex flex-col justify-center items-center p-4 text-black">
      <h1 class="text-4xl mb-4 font-bold">{{ $t('login.title') }}</h1>
      <!-- 新增的文字信息 -->
      <p class="text-gray-700 mb-4">{{ $t('login.loginPrompt') }}</p>
      <p class="text-gray-700 mb-4">
        {{ $t('login.registerPrompt') }}
        <router-link to="/register" class="underline text-blue-500 hover:text-blue-700">{{
          $t('login.register')
        }}</router-link>
      </p>
      <a-form :model="store.loginFormState" layout="vertical" class="w-[300px]">
        <!-- 登录方式选择 -->
        <a-button-group class="mb-4 flex justify-center">
          <a-button
            type="default"
            class="w-[80px]"
            :class="{
              'active-login-button': loginType === 'username',
              'inactive-login-button': loginType !== 'username'
            }"
            @click="loginType = 'username'"
          >
            {{ $t('login.account') }}
          </a-button>
          <a-button
            class="w-[80px]"
            type="default"
            :class="{
              'active-login-button': loginType === 'email',
              'inactive-login-button': loginType !== 'email'
            }"
            @click="loginType = 'email'"
          >
            {{ $t('login.email') }}
          </a-button>
        </a-button-group>

        <!-- 动态显示不同的输入框 -->
        <a-form-item v-if="loginType === 'username'" name="username" :label="$t('login.account')">
          <a-input size="large" v-model:value="store.loginFormState.username">
            <template #prefix>
              <user-outlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item v-if="loginType === 'email'" name="email" :label="$t('login.email')">
          <a-input size="large" v-model:value="store.loginFormState.email">
            <template #prefix>
              <mail-outlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item v-if="loginType === 'username'" name="password" :label="$t('login.password')">
          <a-input-password size="large" v-model:value="store.loginFormState.password">
            <template #prefix>
              <lock-outlined />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item v-if="loginType === 'email'" name="captcha" :label="$t('login.captcha')">
          <a-row :gutter="[8, 8]">
            <a-col :span="16">
              <a-input size="large" v-model:value="store.loginFormState.captcha">
                <template #prefix>
                  <security-scan-outlined />
                </template>
              </a-input>
            </a-col>
            <a-col :span="8">
              <a-button
                type="primary"
                size="large"
                block
                :disabled="sendingCaptcha"
                @click="sendCaptcha"
              >
                <span class="text-sm">
                  {{ sendingCaptcha ? $t('login.sendingCaptcha') : $t('login.sendCaptcha') }}
                </span>
              </a-button>
            </a-col>
          </a-row>
        </a-form-item>
        <a-form-item>
          <a-button
            type="primary"
            html-type="submit"
            size="large"
            block
            class="w-full"
            @click="handleLogin"
          >
            {{ $t('login.loginButton') }}
          </a-button>
        </a-form-item>
        <div class="flex items-center my-4">
          <div class="flex-grow border-t border-gray-300"></div>
          <span class="mx-4 text-gray-500">{{ $t('login.otherLoginMethods') }}</span>
          <div class="flex-grow border-t border-gray-300"></div>
        </div>
        <!-- 第三方登录 -->
        <div class="mt-4 flex justify-center gap-4">
          <WechatOutlined @click="thirdPartyLogin('wechat')" class="text-2xl cursor-pointer" />
          <QqOutlined @click="thirdPartyLogin('qq')" class="text-2xl cursor-pointer" />
        </div>
      </a-form>
    </div>
  </div>
</template>
<script setup lang="ts">
import { LockOutlined, MailOutlined, UserOutlined } from '@ant-design/icons-vue'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { UserStore } from '../stores/user'
const { t } = useI18n()

// 直接导入图片
import { newWarningMessage } from '@renderer/pkg/messages'
import image3 from '../assets/icon/工作.svg'
import image2 from '../assets/icon/开心.svg'
import image4 from '../assets/icon/弹吉他.svg'
import image5 from '../assets/icon/钓鱼.svg'
import image1 from '../assets/icon/骑电车.svg'

const store = UserStore()

const loginType = ref('username')

const sendingCaptcha = ref(false)
const slides = [
  { image: image1, text: 'Slide 1' },
  { image: image2, text: 'Slide 2' },
  { image: image3, text: 'Slide 3' },
  { image: image4, text: 'Slide 4' },
  { image: image5, text: 'Slide 5' }
]

const handleLogin = async () => {
  if (loginType.value === 'username') {
    await store.LoginIn()
  } else if (loginType.value === 'email') {
    // 假设你有一个专门的邮箱登录方法
    await store.LoginByEmail()
  }
}

const thirdPartyLogin = (provider: string) => {
  switch (provider) {
    case 'wechat':
      // 处理微信登录
      newWarningMessage(t('common.thisFeature'))
      break
    case 'qq':
      // 处理QQ登录
      newWarningMessage(t('common.thisFeature'))
      break
    default:
      console.log('Unsupported provider')
  }
}
</script>

<style scoped>
:deep(.slick-slide) {
  text-align: center;
  height: 100vh;
  line-height: 100vh;
  background: #364d79;
  overflow: hidden;
}
:deep(.slick-slide h3) {
  color: #fff;
}

.active-login-button {
  background-color: #1890ff !important;
  border-color: #1890ff !important;
  color: white !important;
}

.inactive-login-button {
  background-color: white !important;
  border-color: #d9d9d9 !important;
  color: black !important;
}
</style>
