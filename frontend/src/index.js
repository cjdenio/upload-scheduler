import Vue from "vue";

// Vue plugins
import Buefy from "buefy";
import "buefy/dist/buefy.css";

Vue.use(Buefy);

import Index from "./views/Index.vue";

new Vue({
  render: (e) => e(Index),
}).$mount("#app");
