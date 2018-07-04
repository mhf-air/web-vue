<template lang="pug">
IconBase(
    iconName="check"

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
  polyline(
      points="20 50, 40 80, 90 20"
      )
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-check",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
</style>
