<template>
  <div class="w-full h-[90px]">
    <div class="flex justify-center items-center w-full h-full px-4">
      <a-input
        size="large"
        ref="inputRef"
        v-model:value="message"
        class="rounded-2xl h-12"
        :placeholder="$t('chat.input_message')"
      >
        <template #suffix>
          <IconFont class="cursor-pointer text-[20px]" type="icon-image" />
          <IconFont class="cursor-pointer text-[20px]" type="icon-3801wenjian" />
          <a-button
            type="primary"
            shape="round"
            class="flex items-center justify-center"
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

<style scoped></style>
