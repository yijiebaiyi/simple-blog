<script setup lang="ts">
  import api from '@/api';
  const articlesListApiResult = await api.ArticlesList()
  const status = articlesListApiResult?.code;
  const articlesList  = articlesListApiResult?.data;

  function getLocationHref(id:number) {
    return location.href=`/poem/${id}`;
  }
</script>

<template>
  <div class="container">
    <div>
      <div class="bg-light p-5 rounded">
        <div class="col-sm-8 mx-auto">
          <h1>小诗</h1>
          <p>
            <div class="row row-cols-1 row-cols-md-3 g-4">
              <div class="col" v-for="article in articlesList">
                <div class="card" @click=(getLocationHref(article.id))>
                  <img src="https://w.wallhaven.cc/full/j3/wallhaven-j3m8y5.png" class="card-img-top" alt="...">
                  <div class="card-body">
                    <h5 class="card-title">{{article.article_title}}</h5>
                    <p class="card-text">
                      {{article.article_desc}}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </p>
        </div>

      </div>
    </div>
  </div>
</template>