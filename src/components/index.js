function install(Vue) {
  components.map((component) => {
    Vue.component(component.name, component)
  })
}

export default {
  install,
}

export function use(Vue, componentList) {
  componentList.map((c) => {
    Vue.component(c.name, c)
  })
}

// app

// icon
import IconBase from "./icon/aa-icon-base.vue"
import IconAngleBracket from "./icon/angle-bracket.vue"
import IconBack from "./icon/back.vue"
import IconForward from "./icon/forward.vue"
import IconCheck from "./icon/check.vue"
import IconCross from "./icon/cross.vue"
import IconTriangle from "./icon/triangle.vue"
import IconSearch from "./icon/search.vue"
import IconMessage from "./icon/message.vue"
import IconTreeOpen from "./icon/tree-open.vue"
import IconTreeClosed from "./icon/tree-closed.vue"
import IconLoading from "./icon/loading.vue"

// basic
import V from "./basic/v.vue"
import H from "./basic/h.vue"
import { PageCache } from "./basic/page-cache.js"

// data

// form

// navigation

// notice

// other

export {

  IconBase,
  IconAngleBracket,
  IconBack,
  IconForward,
  IconCheck,
  IconCross,
  IconTriangle,
  IconSearch,
  IconMessage,
  IconTreeOpen,
  IconTreeClosed,
  IconLoading,

  V,
  H,
  PageCache,

}

const components = [

  IconBase,
  IconAngleBracket,
  IconBack,
  IconForward,
  IconCheck,
  IconCross,
  IconTriangle,
  IconSearch,
  IconMessage,
  IconTreeOpen,
  IconTreeClosed,
  IconLoading,

  V,
  H,
  PageCache,

]
