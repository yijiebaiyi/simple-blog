<script setup lang="ts">
  import Nav from '../components/Nav.vue'
  import PoemList from '../components/PoemList.vue'
</script>

<template>
  <Nav />
  <suspense>
    <PoemList />
  </suspense>
</template>
  