<script setup lang="ts">
import Nav from '../components/Nav.vue'
import Article from '../components/Article.vue'
const id = window.location.pathname.split("/").pop();
</script>
<template>
  <Nav />
  <suspense>
    <Article :id="id"/>
  </suspense>
</template>
