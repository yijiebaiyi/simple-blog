<script setup lang="ts">
import Nav from '../components/Nav.vue'
import ArticleList from '../components/ArticleList.vue'
</script>

<template>
  <Nav />
  <div class="row row-cols-1 row-cols-md-2 g-4" style="width: 1280px;
  margin: 0 auto;
  padding: 2rem;
  font-weight: normal;">
    <suspense>
      <ArticleList />
    </suspense>  
  </div>
</template>
