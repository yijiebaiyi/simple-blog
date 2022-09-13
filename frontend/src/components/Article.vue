<script setup lang="ts">
import api from '@/api';
import { defineProps, markRaw } from 'vue'
import {remark} from 'remark'
import remarkGfm from 'remark-gfm'
import remarkToc from 'remark-toc'
import markdownIt from "markdown-it"
  
// 接收父组件传递过来的值
const props = defineProps(['id'])
const id = props.id as number;

const articlesDetailApiResult = await api.ArticlesDetail(id)
const articlesDetail  = articlesDetailApiResult?.data;


// const html = marked.parse('# Marked in Node.js\n\nRendered by **marked**.');

// console.log(html)

// const htmlContent = String(await remark()
//     .use(remarkGfm)
//     .use(remarkToc)
//     .process(articlesDetail.article_content))

// articlesDetail.article_content = htmlContent;

let md = new markdownIt();
var htmlContent = md.render(articlesDetail.article_content);
console.log(md.render(''))
articlesDetail.article_content = htmlContent;

</script>

<template>
  <div class="container">
    <div>
      <div class="bg-light p-5 rounded">
        <div class="col-sm-8 mx-auto">
          <h1>{{ articlesDetail.article_title }}</h1>
          {{articlesDetail.article_content}}
        </div>
      </div>

      <div class="my-3 p-3 bg-body rounded shadow-sm">
    <h6 class="border-bottom pb-2 mb-0">最近评论</h6>
    <div class=" text-muted pt-3">
      <p class="pb-3 mb-0 small lh-sm border-bottom">
        <strong class="d-block text-gray-dark"><i>张三疯</i> &nbsp说道:</strong>
        地维赖以立，天柱赖以尊
        <small class="d-block text-end mt-3">
          <span href="#">2020-01-01 20:20</span>
        </small>
      </p>
    </div>
    <div class="text-muted pt-3">
      <p class="pb-3 mb-0 small lh-sm border-bottom">
        <strong class="d-block text-gray-dark"><i>周星星</i> &nbsp说道:</strong>
        请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？请问博主，如何看待叙利亚的局势？
        <small class="d-block text-end mt-3">
          <span href="#">2020-01-01 20:20</span>
        </small>
      </p>
    </div>
      <div class="text-muted pt-3">
        <p class="pb-3 mb-0 small lh-sm border-bottom">
          <strong class="d-block text-gray-dark"><i>王七七</i> &nbsp说道:</strong>
          请问博主，如何看待最近的中美关系？
          <small class="d-block text-end mt-3">
            <span href="#">2020-01-01 20:20</span>
          </small>
        </p>
      </div>

    </div>

    <button type="button" class="btn btn-outline-dark" data-bs-toggle="modal" data-bs-target="#exampleModal" data-bs-whatever="@mdo">我也来整两句</button>
    <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">大侠，留下👣 吧：</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form>
              <div class="mb-3">
                <label for="recipient-name" class="col-form-label">大侠怎么称呼:</label>
                <input type="text" class="form-control" id="recipient-name"  placeholder="您的名字是？（必填）">
              </div>
              <div class="mb-3">
                <label for="recipient-name" class="col-form-label">大侠怎么联系您:</label>
                <input type="text" class="form-control" id="recipient-name" placeholder="邮箱是？（非必填）">
              </div>
              <div class="mb-3">
                <label for="recipient-name" class="col-form-label">大侠在哪里可以找到您:</label>
                <input type="text" class="form-control" id="recipient-name" placeholder="个人站地址是？（非必填）">
              </div>
              <div class="mb-3">
                <label for="message-text" class="col-form-label">大侠请留言:</label>
                <textarea class="form-control" id="message-text"  placeholder="高低整两句（必填）"></textarea>
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
            <button type="button" class="btn btn-dark">Send message</button>
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
<!-- <link href="../assets/dist/css/bootstrap.min.css" rel="stylesheet" /> -->
<!-- <link rel="canonical" href="https://getbootstrap.com/docs/5.2/examples/navbars/" /> -->
<!-- <link href="navbar.css" rel="stylesheet" /> -->
<!-- <script src="../assets/dist/js/bootstrap.bundle.min.js"></script> -->


