<script setup lang="ts">
import api, { CODE_STATUS_OK } from '@/api';
import { defineProps, markRaw, proxyRefs, reactive, toRef, toRefs, ref } from 'vue'
import markdownIt from "markdown-it"

// example:
/*
  props:{
    id:{
      type:Number,
      default:0
    }
  },
*/
// 接收父组件传递过来的值
const props = defineProps(['id'])
const id = props.id as number;

const state = reactive({
  articleId: Number(id),
  content: "",
  name: "",
  url: "",
  email: "",
  tel: ""
})

// 文章详情
const articlesDetailApiResult = await api.ArticlesDetail(id)
const articlesDetail = articlesDetailApiResult?.data;
let md = new markdownIt();
var htmlContent = md.render(articlesDetail.article_content);
articlesDetail.article_content = htmlContent;

// 评论列表
const articlesCommentsResult = await api.ArticlesCommentsList(id);
const articlesComments = articlesCommentsResult?.data;

// 评论创建
let showCommentDiv = ref(true)
async function commentCreate() {
  const createResponse = await api.ArticlesCommentsCreate(state);
  if (createResponse.code != CODE_STATUS_OK) {

  }
  alert(createResponse.message)
  showCommentDiv.value = false
  location.reload();
}

</script>

<template>
  <div class="container">
    <div>
      <div class="bg-light p-5 rounded">
        <div class="col-sm-8 mx-auto">
          <h1>{{ articlesDetail.article_title }}</h1>
          <div v-html="articlesDetail.article_content"></div>
        </div>
      </div>

      <div class="my-3 p-3 bg-body rounded shadow-sm">
        <h6 class="border-bottom pb-2 mb-0">最近评论</h6>
        <div class="text-muted pt-3" v-for="comment in articlesComments">
          <p class="pb-3 mb-0 small lh-sm border-bottom">
            <strong class="d-block text-gray-dark"><i><a v-bind:href="comment.comment_url">{{ comment.comment_name
            }}</a></i> &nbsp说道:</strong>
            {{ comment.comment_content }}
            <small class="d-block text-end mt-3">
              <span href="#">{{ comment.comment_create_on }}</span>
            </small>
          </p>
        </div>
      </div>

      <button type="button" class="btn btn-outline-dark" data-bs-toggle="modal" data-bs-target="#exampleModal"
        data-bs-whatever="@mdo">我要评论</button>
      <div v-if="showCommentDiv" class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel"
        aria-hidden="true">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="exampleModalLabel">留下👣 吧：</h5>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <form>
                <div class="mb-3">
                  <label for="comment-name" class="col-form-label">姓名:</label>
                  <input type="text" class="form-control" v-model="state.name" name="name" id="comment-name"
                    placeholder="请输入您的姓名">
                </div>
                <div class="mb-3">
                  <label for="comment-email" class="col-form-label">联系方式:</label>
                  <input type="text" class="form-control" v-model="state.email" name="email" id="comment-email"
                    placeholder="请输入您的邮箱或电话">
                </div>
                <div class="mb-3">
                  <label for="comment-url" class="col-form-label">个人站点:</label>
                  <input type="text" v-model="state.url" class="form-control" name="url" id="comment-url"
                    placeholder="请输入您的个人站地址">
                </div>
                <div class="mb-3">
                  <label for="comment-content" class="col-form-label">留言内容:</label>
                  <textarea class="form-control" id="comment-content" v-model="state.content" name="content"
                    placeholder="请输入留言内容"></textarea>
                </div>
              </form>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
              <!-- <button type="button" class="btn btn-dark" @click=(commentCreate())>Send message</button> -->
              <button type="button" class="btn btn-dark" @click="commentCreate">Send message</button>
            </div>
          </div>
        </div>
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

