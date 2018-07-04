<template lang="pug">
icon-angle-bracket.icon-root(direction="right" width="20")
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-forward",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
.icon-root
  color: #9e9e9e
</style>
