<template>
  <a-form :model="form" :style="{ width: '100%' }" @submit="handleSubmit" layout="vertical">
    <!-- 用户名 -->
    <a-form-item field="name" tooltip="请输入用户名" label="用户名">
      <a-input v-model="form.name" placeholder="请输入用户名..." class="input-field" />
    </a-form-item>

    <!-- 密码 -->
    <a-form-item field="password" tooltip="请输入密码" label="密码">
      <a-input
        v-model="form.password"
        type="password"
        placeholder="请输入密码..."
        class="input-field"
      />
    </a-form-item>
  </a-form>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

const form = reactive({
  name: '',
  password: ''
})

// 表单提交处理
const handleSubmit = () => {
  if (!form.name || !form.password) {
    console.log('请输入用户名和密码')
    return
  }
  console.log('表单提交成功:', form)
}
</script>

<style scoped>
/* 输入框样式 */
.input-field {
  height: 40px;
  border-radius: 6px;
}
</style>
