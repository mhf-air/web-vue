<template lang="pug">
transition(:name="transitionName")
  g-page-cache(:initCache="pageCache" @new-key="newKey")
    router-view.root
</template>

<script>
export default {
  data() {
    return {
      transitionName: "",

      pageCache: Object.create(null),
      pageKeyList: new Array(2),
    }
  },
  methods: {
    newKey(key) {
      let keyList = this.pageKeyList
      keyList[0] = keyList[1]
      keyList[1] = key

      if (this.$root.IsBack) {
        this.removePageCache(keyList[0])
      }
    },

    removePageCache(key) {
      const v = this.pageCache[key]
      if (v) {
        v.componentInstance.$destroy()
      }
      this.pageCache[key] = null
    },
  },
  watch: {
    "$route"(to, from) {
      if (this.$root.JustClickBack) {
        this.transitionName = "slide-right"
        this.$root.JustClickBack = false
      } else {
        this.transitionName = "slide-left"
        this.$root.IsBack = false
      }
    },
  },
}
</script>

<style lang="stylus" scoped>
@import "~common.styl"

.slide-left-leave
.slide-left-enter-to
  transform: translateX(0%)

.slide-left-leave-to
  transform: translateX(-100%)

.slide-left-enter
  transform: translateX(100%)

.slide-right-leave
.slide-right-enter-to
  transform: translateX(0%)

.slide-right-leave-to
  transform: translateX(100%)

.slide-right-enter
  transform: translateX(-100%)

.root
  transition: transform 300ms

  width: 100%
  height: 100%
  position: fixed
  background: grey200
  overflow-x: hidden
  overflow-y: scroll
  &::-webkit-scrollbar
    display: none

</style>
