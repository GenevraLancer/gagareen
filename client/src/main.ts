import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import 'bootstrap/scss/bootstrap.scss';
import VueMaterial from 'vue-material'
import OpenApi from 'vue-openapi'
import 'vue-material/dist/vue-material.css'
import VueResource from 'vue-resource'
import jsonApi from '@/swagger/field.swagger.json' //@ is alias to /src

createApp(App)
  .use(store)
  .use(router)
  .use(VueMaterial)
  .use(VueResource)
  .mount("#app");