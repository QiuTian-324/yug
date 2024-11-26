<template>
  <div class="h-full text-white w-full p-2">
    <a-tree
      v-model:selectedKeys="selectedKeys"
      blockNode
      :tree-data="treeData"
      show-icon
      default-expand-all
      @select="onSelect"
    >
      <template #switcherIcon="{ switcherCls }">
        <down-outlined :class="switcherCls" />
      </template>
      <template #icon="{ key, selected }">
        <template v-if="key === 'group'">
          <div class="h-full flex items-center justify-center">
            <user-outlined />
          </div>
        </template>
        <template v-else-if="key.startsWith('friend-')">
          <div class="h-full flex items-center justify-center">
            <user-outlined />
          </div>
        </template>
        <template v-else>
          <user-outlined v-if="selected" />
          <user-outlined v-else />
        </template>
      </template>
    </a-tree>
  </div>
</template>

<script lang="ts" setup>
import { DownOutlined, UserOutlined } from '@ant-design/icons-vue'
import { computed, ref } from 'vue'
import { UserStore } from '../../stores/user'

const store = UserStore()

const treeData = computed(() => [
  {
    title: '群组',
    key: 'group',
    children: [
      { title: 'leaf', key: 'node4' },
      { title: 'leaf', key: 'node5' }
    ]
  },
  {
    title: '好友',
    key: 'friend',
    children: store.friendList.map((friend) => ({
      title: friend.username,
      key: `friend-${friend.user_id}`
    }))
  }
])

const selectedKeys = ref(['friend-0'])

function onSelect(selectedKeys) {
  const selectedKey = selectedKeys[0]
  if (selectedKey.startsWith('friend-')) {
    const userId = selectedKey.split('-')[1]
    store.selectedSeesionId = userId
  }
}
</script>
