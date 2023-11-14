<script setup lang="ts">
  import Nav from '../components/Nav.vue'
  import Novelette from '../components/Novelette.vue'
  const id = window.location.pathname.split("/").pop();
</script>
<template>
  <Nav />
  <suspense>
    <Novelette :id="id"/>
  </suspense>
</template>
  
  