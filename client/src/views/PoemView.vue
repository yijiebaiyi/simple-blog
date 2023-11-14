<script setup lang="ts">
  import Nav from '../components/Nav.vue'
  import Poem from '../components/Poem.vue'
  const id = window.location.pathname.split("/").pop();
</script>
<template>
  <Nav />
  <suspense>
    <Poem :id="id"/>
  </suspense>
</template>