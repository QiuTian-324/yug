<template>
  <div
    class="h-screen w-screen flex justify-center items-center bg-gradient-to-r from-common-white via-sky-100 to-sky-50 overflow-hidden"
  >
    <div class="shadow-lg rounded-lg p-6 w-full max-w-md">
      <h2 class="text-2xl font-bold text-center pb-4 text-gray-800">{{ $t('login.title') }}</h2>
      <a-form :model="store.loginFormState" layout="vertical">
        <!-- 登录方式选择 -->
        <div class="mb-4">
          <a-tabs v-model:activeKey="loginType" type="line" class="w-full">
            <a-tab-pane key="username" :tab="$t('login.account')">
              <!-- 用户名登录 -->
              <a-form-item name="username">
                <a-input
                  size="large"
                  placeholder="请输入用户名"
                  v-model:value="store.loginFormState.username"
                >
                  <template #prefix>
                    <user-outlined />
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item name="password">
                <a-input-password
                  size="large"
                  placeholder="请输入密码"
                  v-model:value="store.loginFormState.password"
                >
                  <template #prefix>
                    <lock-outlined />
                  </template>
                </a-input-password>
              </a-form-item>
            </a-tab-pane>
            <a-tab-pane key="email" :tab="$t('login.email')">
              <!-- 邮箱登录 -->
              <a-form-item name="email">
                <a-input
                  size="large"
                  placeholder="请输入邮箱"
                  v-model:value="store.loginFormState.email"
                >
                  <template #prefix>
                    <mail-outlined />
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item name="captcha">
                <a-row :gutter="[8, 8]">
                  <a-col :span="16">
                    <a-input
                      size="large"
                      placeholder="请输入验证码"
                      v-model:value="store.loginFormState.captcha"
                    >
                      <template #prefix>
                        <security-scan-outlined />
                      </template>
                    </a-input>
                  </a-col>
                  <a-col :span="8">
                    <a-button
                      size="large"
                      type="primary"
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
            </a-tab-pane>
            <a-tab-pane key="phone" :tab="$t('login.phone')">
              <!-- 手机号登录 -->
              <a-form-item name="phone">
                <a-input
                  size="large"
                  placeholder="请输入手机号"
                  v-model:value="store.loginFormState.phone"
                >
                  <template #prefix>
                    <phone-outlined />
                  </template>
                </a-input>
              </a-form-item>
              <a-form-item name="phoneCaptcha">
                <a-row :gutter="[8, 8]">
                  <a-col :span="16">
                    <a-input
                      size="large"
                      placeholder="请输入验证码"
                      v-model:value="store.loginFormState.phoneCaptcha"
                    >
                      <template #prefix>
                        <security-scan-outlined />
                      </template>
                    </a-input>
                  </a-col>
                  <a-col :span="8">
                    <a-button
                      size="large"
                      type="primary"
                      block
                      :disabled="sendingCaptcha"
                      @click="sendPhoneCaptcha"
                    >
                      <span class="text-sm">
                        {{ sendingCaptcha ? $t('login.sendingCaptcha') : $t('login.sendCaptcha') }}
                      </span>
                    </a-button>
                  </a-col>
                </a-row>
              </a-form-item>
            </a-tab-pane>
          </a-tabs>
        </div>

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
          <GithubOutlined
            @click="thirdPartyLogin('github')"
            class="text-2xl cursor-pointer hover:text-gray-700 transition duration-300"
          />
          <WechatOutlined
            @click="thirdPartyLogin('wechat')"
            class="text-2xl cursor-pointer hover:text-gray-700 transition duration-300"
          />
          <QqOutlined
            @click="thirdPartyLogin('qq')"
            class="text-2xl cursor-pointer hover:text-gray-700 transition duration-300"
          />
        </div>
      </a-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { LockOutlined, MailOutlined, UserOutlined, PhoneOutlined } from '@ant-design/icons-vue'
import { UserStore } from '@renderer/stores/user/user'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const store = UserStore()

const loginType = ref('username')

const sendingCaptcha = ref(false)

onMounted(() => {
  const params = new URLSearchParams(window.location.search)
  if (params.has('code')) {
    const code = params.get('code')
    store.LoginByGithubCode(code)
  }
})

const handleLogin = async () => {
  if (loginType.value === 'username') {
    await store.handleLoginIn()
  } else if (loginType.value === 'email') {
    await store.LoginByEmail()
  } else if (loginType.value === 'phone') {
    await store.LoginByPhone()
  }
}

const sendPhoneCaptcha = async () => {
  sendingCaptcha.value = true
  // 假设有一个发送手机验证码的方法
  await store.sendPhoneCaptcha()
  sendingCaptcha.value = false
}

const thirdPartyLogin = (provider: string) => {
  switch (provider) {
    case 'github':
      store.GithubLogin()
      break
    case 'wechat':
      newWarningMessage(t('common.thisFeature'))
      break
    case 'qq':
      newWarningMessage(t('common.thisFeature'))
      break
    default:
      console.log('Unsupported provider')
  }
}
</script>

<style scoped>
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
