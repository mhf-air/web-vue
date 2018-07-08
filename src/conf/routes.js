// front-end routes

import NotFound from "../root/NotFound.vue"
import Main from "../root/main/a.vue"

import AppLogin from "../root/app/login/a.vue"

// ================================================================================
export default [
  { path: "/", component: Main },

  { path: "/app/login", component: AppLogin },

  { path: "*", component: NotFound },
]
