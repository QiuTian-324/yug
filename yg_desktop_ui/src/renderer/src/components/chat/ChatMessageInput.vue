<template>
  <div class="w-full h-[90px]">
    <div class="flex justify-center items-center w-full h-full px-4">
      <a-input
        size="large"
        ref="inputRef"
        v-model:value="message"
        class="rounded-2xl h-12 bg-light-base text-light-text"
      >
        <template #suffix>
          <IconFont
            class="cursor-pointer text-[20px] text-light-text hover:text-icon-hover"
            type="icon-image"
          />
          <IconFont
            class="cursor-pointer text-[20px] text-light-text hover:text-icon-hover"
            type="icon-3801wenjian"
          />
          <a-button
            type="primary"
            shape="round"
            class="flex items-center justify-center bg-button-primary"
            @click="sendMessage"
          >
            <send-outlined size="20" />
          </a-button>
        </template>
      </a-input>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

const message = ref('')

const inputRef = ref(null)

onMounted(() => {
  if (inputRef.value) {
    inputRef.value.focus()
  }
})

// 发送消息的方法
const sendMessage = () => {
  if (message.value.trim()) {
    console.log(message.value)
    message.value = '' // 发送后清空输入框
  }
}
</script>

<style scoped>
:deep(.css-dev-only-do-not-override-1j3pbsq).ant-input-affix-wrapper > input.ant-input {
  @apply bg-light-base;
}
</style>
