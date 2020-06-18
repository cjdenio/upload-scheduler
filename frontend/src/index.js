// Core Vue libraries
import Vue from "vue";

// Vue plugins
import vuetify from "./plugins/vuetify";
import router from "./router";

// App
import App from "./App.vue";
import "./style.scss";

new Vue({
  render: (e) => e(App),
  vuetify,
  router,
}).$mount("#app");
