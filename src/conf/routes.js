// front-end routes

import NotFound from "../root/NotFound.vue"

import Main from "../root/main/a.vue"

// ================================================================================
export default [
  { path: "/", component: Main },

  { path: "*", component: NotFound },
]
