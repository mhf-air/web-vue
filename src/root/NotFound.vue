<template lang="pug">
g-h(j-c="space-around" a-i="center")
  span.path {{ routePath }}
  span Page Not Found
</template>

<script>
export default {
  data(){
    return {
      routePath: this.$route.path,
    }
  },
}
</script>

<style lang="stylus" scoped>

.path
  font-weight: bold
  font-size: 20px

</style>
