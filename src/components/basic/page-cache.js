// see also https://github.com/vuejs/vue/blob/dev/src/core/components/keep-alive.js

module.exports = {
  PageCache: {
    name: "g-page-cache",

    // need this for avoiding blank screen when transitioning.
    // also note that abstract is not in public API, so it might change
    abstract: true,

    /* created() {
      this.cache = Object.create(null)
    }, */

    props: {
      initCache: {
        type: Object,
        default: () => Object.create(null),
      },
    },

    data() {
      return {
        cache: this.initCache,
        pageKey: "",
      }
    },

    destroyed() {
      for (const key in this.cache) {
        this.removeCache(key)
      }
    },

    methods: {
      removeCache(key) {
        const v = this.cache[key]
        if (v) {
          v.componentInstance.$destroy()
        }
        this.cache[key] = null
      },
    },

    render() {
      const slot = this.$slots.default
      const vnode = slot[0] // TODO

      let componentOptions = vnode && vnode.componentOptions
      if (componentOptions) {
        let { cache } = this
        let key = (vnode.key == null) ?
          componentOptions.Ctor.cid + (componentOptions.tag ? `::${componentOptions.tag}` : "") :
          vnode.key

        this.$emit("new-key", key)

        if (cache[key]) {
          vnode.componentInstance = cache[key].componentInstance
        } else {
          cache[key] = vnode
        }

        // required for keeping alive vnode
        vnode.data.keepAlive = true
      }

      return vnode
    },

  },
}
