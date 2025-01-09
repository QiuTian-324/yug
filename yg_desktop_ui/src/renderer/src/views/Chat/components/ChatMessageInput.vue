<template>
  <div class="w-full">
    <div class="flex justify-center items-center w-full h-full px-4">
      <div class="flex flex-col items-start w-full">
        <div class="flex space-x-2 mb-2">
          <IconFont
            class="cursor-pointer text-[20px] text-light-text hover:text-icon-hover"
            type="icon-image"
          />
          <IconFont
            class="cursor-pointer text-[20px] text-light-text hover:text-icon-hover"
            type="icon-3801wenjian"
          />
        </div>
        <a-input
          size="large"
          ref="inputRef"
          v-model:value="message"
          class="rounded-lg h-12 dark:bg-dark-fill-2 w-full"
          @keyup.enter="sendMessage"
        >
          <template #suffix>
            <a-tooltip :title="$t('chat.sendMessage')">
              <div
                @click="sendMessage"
                class="flex items-center justify-center cursor-pointer bg-light-primary-6 dark:bg-dark-primary-6 w-8 h-8 rounded-full text-common-white"
              >
                <IconFont type="icon-fasong" class="text-lg" />
              </div>
            </a-tooltip>
          </template>
        </a-input>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

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
:deep(.css-dev-only-do-not-override-1p3hq3p).ant-input-affix-wrapper > input.ant-input {
  @apply dark:bg-dark-fill-2 text-common-black dark:text-common-white;
}
</style>
