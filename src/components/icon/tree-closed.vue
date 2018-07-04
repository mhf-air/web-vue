<template lang="pug">
IconBase(
    iconName="tree-closed"

    :width="width"
    :height="height"

    :fill="fill"
    :fill-opacity="fillOpacity"

    :stroke="stroke"
    :stroke-width="strokeWidth"
    :stroke-linecap="strokeLinecap"
    :stroke-linejoin="strokeLinejoin"
    :stroke-opacity="strokeOpacity"
    )
  polygon(
      points="24 30, 70 50, 24 70"
      )
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-tree-closed",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
</style>
