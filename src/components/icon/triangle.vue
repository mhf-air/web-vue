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
  polygon(
      points="24 30, 50 70, 76 30"
      )
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-triangle",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
</style>
