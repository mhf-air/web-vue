<template lang="pug">
IconBase(
    iconName="search"

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
  circle(cx="35" cy="35" r="30")
  line(x1="57" y1="57" x2="90" y2="90")
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-search",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
</style>
