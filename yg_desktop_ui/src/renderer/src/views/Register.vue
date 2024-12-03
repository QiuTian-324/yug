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
    <!-- 右侧注册表单 -->
    <div class="bg-white flex-1 flex flex-col justify-center items-center p-4 text-black">
      <h1 class="text-4xl mb-4 font-bold">{{ $t('register.title') }}</h1>
      <p class="text-gray-700 mb-4">{{ $t('register.registerPrompt') }}</p>
      <p class="text-gray-700 mb-4">
        {{ $t('register.hasAccount') }}
        <router-link to="/login" class="underline text-blue-500 hover:text-blue-700">
          {{ $t('register.login') }}
        </router-link>
      </p>
      <a-form :model="registerFormState" layout="vertical" class="w-[300px]">
        <!-- 注册方式选择 -->
        <a-button-group class="mb-4 flex justify-center">
          <a-button
            type="default"
            class="w-[80px]"
            :class="{
              'active-register-button': registerType === 'username',
              'inactive-register-button': registerType !== 'username'
            }"
            @click="registerType = 'username'"
          >
            {{ $t('register.account') }}
          </a-button>
          <a-button
            class="w-[80px]"
            type="default"
            :class="{
              'active-register-button': registerType === 'email',
              'inactive-register-button': registerType !== 'email'
            }"
            @click="registerType = 'email'"
          >
            {{ $t('register.email') }}
          </a-button>
        </a-button-group>

        <!-- 动态显示不同的输入框 -->
        <a-form-item
          v-if="registerType === 'username'"
          name="username"
          :label="$t('register.account')"
        >
          <a-input size="large" v-model:value="registerFormState.username">
            <template #prefix>
              <user-outlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item v-if="registerType === 'email'" name="email" :label="$t('register.email')">
          <a-input size="large" v-model:value="registerFormState.email">
            <template #prefix>
              <mail-outlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item
          v-if="registerType === 'username'"
          name="password"
          :label="$t('register.password')"
        >
          <a-input-password size="large" v-model:value="registerFormState.password">
            <template #prefix>
              <lock-outlined />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item v-if="registerType === 'email'" name="captcha" :label="$t('register.captcha')">
          <a-row :gutter="[8, 8]">
            <a-col :span="16">
              <a-input size="large" v-model:value="registerFormState.captcha">
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
                  {{ sendingCaptcha ? $t('register.sendingCaptcha') : $t('register.sendCaptcha') }}
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
            @click="handleRegister"
          >
            {{ $t('register.registerButton') }}
          </a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>
<script setup lang="ts">
import {
  LockOutlined,
  MailOutlined,
  SecurityScanOutlined,
  UserOutlined
} from '@ant-design/icons-vue'
import { ref } from 'vue'

// 直接导入图片
import image3 from '../assets/icon/工作.svg'
import image2 from '../assets/icon/开心.svg'
import image4 from '../assets/icon/弹吉他.svg'
import image5 from '../assets/icon/钓鱼.svg'
import image1 from '../assets/icon/骑电车.svg'

const registerFormState = ref({
  username: '',
  email: '',
  password: '',
  captcha: ''
})

const registerType = ref('username')
const sendingCaptcha = ref(false)

const slides = [
  { image: image1, text: 'Slide 1' },
  { image: image2, text: 'Slide 2' },
  { image: image3, text: 'Slide 3' },
  { image: image4, text: 'Slide 4' },
  { image: image5, text: 'Slide 5' }
]

const handleRegister = async () => {
  // 处理注册逻辑
  console.log('Registering user:', registerFormState.value)
}

const sendCaptcha = async () => {
  sendingCaptcha.value = true
  // 模拟发送验证码的逻辑
  setTimeout(() => {
    sendingCaptcha.value = false
    console.log('Captcha sent to', registerFormState.value.email)
  }, 2000)
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

.active-register-button {
  background-color: #1890ff !important;
  border-color: #1890ff !important;
  color: white !important;
}

.inactive-register-button {
  background-color: white !important;
  border-color: #d9d9d9 !important;
  color: black !important;
}
</style>
