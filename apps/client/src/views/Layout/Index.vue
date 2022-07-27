<script setup lang="ts">
import * as S from './styles'

const showDrawer = ref(false)

const route = useRoute()

const routes = [{ label: 'Home', to: '/' }]

function isActive(currentPath: string) {
  return route.fullPath === currentPath
}
</script>

<template>
  <router-view v-slot="{ Component }">
    <S.Header>
      <S.HeaderBtn @click="showDrawer = !showDrawer" />
    </S.Header>
    <S.NavDrawer v-show="showDrawer">
      <S.NavDrawerItem
        v-for="({ label, to }, idx) in routes"
        :key="idx"
        :to="to"
        :isActive="isActive(to)"
      >
        {{ label }}
      </S.NavDrawerItem>
    </S.NavDrawer>

    <S.MainWrapper>
      <component :is="Component" />
    </S.MainWrapper>
  </router-view>
</template>
