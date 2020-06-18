import Vue from "vue";

// Vue plugins
import vuetify from "./plugins/vuetify";

import App from "./App.vue";

import Index from "./views/Index.vue";

import "./style.scss";

new Vue({
  render: (e) => e(App),
  vuetify,
}).$mount("#app");
