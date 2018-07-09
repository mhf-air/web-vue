<template lang="pug">
div.root
  //- img(:src="bookmark" alt="bookmark")
  img(:src="leaf" alt="bookmark")
  div hello
  router-link(to="/app/login") 登录
</template>

<script>
import bookmark from "../../static/image/bookmark-512.png"
import leaf from "../../static/image/leaf.png"

export default {
  data() {
    return {
      bookmark,
      leaf,
    }
  },
}
</script>

<style lang="stylus" scoped>
@import "~common.styl"

</style>
