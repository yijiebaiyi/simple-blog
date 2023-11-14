<script setup lang="ts">
  import Nav from '../components/Nav.vue'
  import NoveletteList from '../components/NoveletteList.vue'
  </script>
  
  <template>
    <Nav />
    <suspense>
      <NoveletteList />
    </suspense>
  </template>
  