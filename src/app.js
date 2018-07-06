import Vue from "vue"

// vuex
import Vuex from "vuex"
Vue.use(Vuex)

import modules from "./conf/store-modules"
const store = new Vuex.Store({
  modules,
})

// router
import VueRouter from "vue-router"
Vue.use(VueRouter)

import routes from "./conf/routes"
const router = new VueRouter({
  routes,
  mode: "history",
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  },
})

// app specific components
import Components from "./components"
Vue.use(Components)

// filter
import { api, util } from "app"
Vue.filter("Money", (value) => {
  if (typeof value === "string") {
    value = parseFloat(value)
  }
  return value.toFixed(2)
})
Vue.filter("Date", (date, option) => {
  let d = new Date(date.replace(" ", "T"))
  return d.getFullYear() + "-" + (d.getMonth() + 1) + "-" + d.getDate()
})
Vue.filter("Enum", (value, entity, field) => {
  if (!api[entity] ||
    !api[entity][field] ||
    !api[entity][field].value ||
    !api[entity][field].value[value]
  ) {
    console.log(`${value} is an unknown enum`)
    return value
  }
  return api[entity][field].value[value].title
})

// root
import AppRoot from "./root/AppRoot.vue"
new Vue({
  router,
  store,

  el: "#app",
  render: (h) => h(AppRoot),

  data: {
    IsBack: false,
    JustClickBack: false,
  },

  methods: {
    Back() {
      this.IsBack = true
      this.JustClickBack = true
      this.$router.back()
    },
  },

  created() {
    // this.checkLoginStatus()
  },

})
