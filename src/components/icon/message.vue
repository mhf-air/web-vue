<template lang="pug">
IconBase(
    iconName="message"

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
  g
    path(
       d="M 51.390625 20.830078 A 36.112953 23.107239 0 0 0 15.279297 43.9375 A 36.112953 23.107239 0 0 0 23.351562 58.480469 L 23.234375 58.457031 L 14.142578 72.220703 L 61.619141 66.287109 L 61.123047 66.185547 A 36.112953 23.107239 0 0 0 87.503906 43.9375 A 36.112953 23.107239 0 0 0 51.390625 20.830078 z "
       )
    ellipse(
       cx="50"
       cy="47"
       rx="4"
       ry="4"
       fill="#f5f5f5"
       )
    ellipse(
       cx="33"
       cy="47"
       rx="4"
       ry="4"
       fill="#f5f5f5"
       )
    ellipse(
       cx="67"
       cy="47"
       rx="4"
       ry="4"
       fill="#f5f5f5"
       )
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-message",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
</style>
