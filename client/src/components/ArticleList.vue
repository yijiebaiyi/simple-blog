<script setup lang="ts">
import api from '@/api';

// 文章列表(含标签)
const articlesListApiResult = await api.ArticlesList()

const status = articlesListApiResult?.code;
const articlesList = articlesListApiResult?.data;

function getLocationHref(id: number) {
  return location.href = `/article/${id}`;
}

</script>

<template >
  <div class="col" v-for="article in articlesList">
    <div class="card" @click=(getLocationHref(article.id))>
      <img src="https://w.wallhaven.cc/full/l3/wallhaven-l3xk6q.jpg" class="card-img-top" alt="...">
      <div class="card-body">
        <h5 class="card-title">{{ article.article_title }}</h5>
        <p class="card-text">这里应该是显示博客的部分内容。。。。</p>

        <span v-for="tag in article.tags">
          <span class="badge bg-secondary">{{ tag.name }}</span>&nbsp;
        </span>
      </div>
    </div>
  </div>
</template>
<style>
.bd-placeholder-img {
  font-size: 1.125rem;
  text-anchor: middle;
  -webkit-user-select: none;
  -moz-user-select: none;
  user-select: none;
}

@media (min-width: 768px) {
  .bd-placeholder-img-lg {
    font-size: 3.5rem;
  }
}

.b-example-divider {
  height: 3rem;
  background-color: rgba(0, 0, 0, .1);
  border: solid rgba(0, 0, 0, .15);
  border-width: 1px 0;
  box-shadow: inset 0 .5em 1.5em rgba(0, 0, 0, .1), inset 0 .125em .5em rgba(0, 0, 0, .15);
}

.b-example-vr {
  flex-shrink: 0;
  width: 1.5rem;
  height: 100vh;
}

.bi {
  vertical-align: -.125em;
  fill: currentColor;
}

.nav-scroller {
  position: relative;
  z-index: 2;
  height: 2.75rem;
  overflow-y: hidden;
}

.nav-scroller .nav {
  display: flex;
  flex-wrap: nowrap;
  padding-bottom: 1rem;
  margin-top: -1px;
  overflow-x: auto;
  text-align: center;
  white-space: nowrap;
  -webkit-overflow-scrolling: touch;
}
</style>


