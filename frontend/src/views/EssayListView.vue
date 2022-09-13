<script setup lang="ts">
import Nav from '../components/Nav.vue'
import EssayList from '../components/EssayList.vue'
</script>

<template>
  <Nav />
  <suspense>
    <EssayList />
  </suspense>  
</template>
