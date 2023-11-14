import { createRouter, createWebHistory } from "vue-router";
import ArticleListView from "../views/ArticleListView.vue";
import AboutView from "../views/AboutView.vue";
import ArticleView from "../views/ArticleView.vue";
import EssayView from "../views/EssayView.vue";
import NoveletteView from "../views/NoveletteView.vue";
import PoemView from "../views/PoemView.vue";
import EssayListView from "../views/EssayListView.vue";
import NoveletteListView from "../views/NoveletteListView.vue";
import PoemListView from "../views/PoemListView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: ArticleListView,
    },
    {
      path: "/about",
      name: "about",
      component: AboutView,
    },
    {
      path: "/article",
      name: "articlelist",
      component: ArticleListView,
    },
    {
      path: "/article/:id",
      name: "article",
      component: ArticleView,
    },
    {
      path: "/essay",
      name: "essaylist",
      component: EssayListView,
    },
    {
      path: "/novelette",
      name: "novelettelist",
      component: NoveletteListView,
    },
    {
      path: "/poem",
      name: "poemlist",
      component: PoemListView,
    },
    {
      path: "/essay/:id",
      name: "essay",
      component: EssayView,
    },
    {
      path: "/novelette/:id",
      name: "novelette",
      component: NoveletteView,
    },
    {
      path: "/poem/:id",
      name: "poem",
      component: PoemView,
    },
  ],
});

export default router;
