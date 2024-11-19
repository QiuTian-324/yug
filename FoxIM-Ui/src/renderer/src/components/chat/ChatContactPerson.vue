<template>
  <div class="w-full h-screen p-2">
    <a-tree
      :default-expand-all="false"
      :default-expanded-keys="defaultExpandedKeys"
      blockNode
      size="large"
      :data="treeData"
      @select="handleSelect"
    >
    </a-tree>
  </div>
</template>

<script setup lang="ts">
import { IconUser, IconUserGroup } from '@arco-design/web-vue/es/icon'
import { UserStore } from '@renderer/stores/user'
import { computed, h } from 'vue'
import { FriendListResponse } from '../../interface/user'

const store = UserStore()

// 生成好友节点
const generateFriendNodes = (friends: FriendListResponse[]) => {
  return friends.map((friend) => ({
    title: friend.username,
    key: `friend-${friend.user_id}`
  }))
}

// 生成群组节点
const generateGroupNodes = () => {
  return [
    // {
    //   title: 'Leaf',
    //   key: 'node4'
    // },
    // {
    //   title: 'Leaf',
    //   key: 'node5'
    // }
  ]
}

// 生成 treeData
const treeData = computed(() => {
  return [
    {
      title: '群组',
      key: 'group',
      icon: () => h(IconUserGroup),
      children: generateGroupNodes()
    },
    {
      title: '好友',
      key: 'friend',
      icon: () => h(IconUser),
      children: generateFriendNodes(store.friendList)
    }
  ]
})

// 默认展开的好友列表
const defaultExpandedKeys = computed(() => {
  if (store.friendList.length > 0) {
    return ['friend']
  }
  return []
})

// 处理点击事件
const handleSelect = (selectedKeys: (string | number)[], info: any) => {
  const selectedKey = selectedKeys[0]
  if (typeof selectedKey === 'string' && selectedKey.startsWith('friend-')) {
    const friendId = parseInt(selectedKey.split('-')[1])
    const selectedFriend = store.friendList.find((friend) => friend.user_id === friendId)
    if (selectedFriend) {
      store.selectedSeesionId = selectedFriend.user_id
    }
  }
}
</script>
