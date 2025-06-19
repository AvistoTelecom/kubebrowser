import { createRouter, createWebHistory } from 'vue-router'

import ViewHome from '@/views/ViewHome.vue'

export enum RouteName {
  Home = '/',
  NotFound = '/:pathMatch(.*)*',
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: RouteName.Home,
      name: RouteName.Home,
      component: ViewHome,
    },

    // {
    //   path: RouteName.NotFound,
    //   name: RouteName.NotFound,
    //   component: () => import('@/views/ViewNotFound.vue'),
    // },
  ],
})

export default router
