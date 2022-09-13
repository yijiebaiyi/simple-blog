import { createApp, type AppContext } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import api, { type ApiMethodList } from "./api";

// import "./assets/main.css";
import "./assets/dist/js/bootstrap.bundle.min.js";
import "./assets/dist/css/bootstrap.min.css";
import "./assets/dist/css/bootstrap.min.css";
import "./assets/navbar.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.mount("#app");
app._props = {
  api,
};
