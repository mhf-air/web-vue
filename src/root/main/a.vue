<template lang="pug">
div.root
  div hello
  router-link(to="/app/login") 登录
</template>

<script>
export default {
  data() {
    return {
    }
  },
}
</script>

<style lang="stylus" scoped>
@import "~common.styl"

</style>
