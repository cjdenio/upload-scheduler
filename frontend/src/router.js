import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

// Routes
import Index from "./views/Index.vue";
import Upload from "./views/Upload.vue";
import About from "./views/About.vue";

let router = new VueRouter({
  routes: [
    {
      component: Index,
      path: "/",
      meta: {
        title: "Scheduled",
      },
    },
    {
      path: "/upload",
      component: Upload,
      meta: {
        title: "Upload",
      },
    },
    {
      path: "/about",
      component: About,
      meta: {
        title: "About",
      },
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = `${to.meta.title} | Upload Scheduler`;
  }
  next();
});

export default router;
