<template>
  <div class="w-full h-[90px]">
    <div class="flex justify-center items-center w-full h-full px-4">
      <a-input size="large" ref="inputRef" class="rounded-2xl h-12">
        <template #suffix>
          <icon-file size="25" class="cursor-pointer" />
          <icon-image class="mx-2 cursor-pointer" size="25" />
          <a-button type="primary" shape="round" @click="sendMessage">
            <icon-send size="20" />
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
