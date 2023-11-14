<script setup lang="ts">
  import Nav from '../components/Nav.vue'
  import Essay from '../components/Essay.vue'
  const id = window.location.pathname.split("/").pop();
</script>
<template>
  <Nav />
  <suspense>
    <Essay :id="id"/>
  </suspense>
</template>
  