import Vue from "vue";

// Vue plugins
import Buefy from "buefy";
import "buefy/dist/buefy.css";

import vuetify from "./plugins/vuetify";

Vue.use(Buefy);

import App from "./App.vue";

import Index from "./views/Index.vue";

import "./style.scss";

new Vue({
  render: (e) => e(App),
  vuetify,
}).$mount("#app");
