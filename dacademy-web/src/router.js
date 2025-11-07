import { createRouter, createWebHistory } from 'vue-router'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
      { path: '/', component: ()=>import("./views/HomeView.vue") },
      { path: '/about', component: ()=>import("./views/AboutView.vue") },
      {
        path: '/c/:slug(.*)',
        component: () => import('./views/CourseView.vue'),
      },
      {
        path: '/c/:courseSlug(.*)/chapter/:chapterID(.*)',
        component: () => import('./views/course/ChapterView.vue'),
      },
  ],
})

export default router
