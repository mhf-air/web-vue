<template lang="pug">
IconBase(
    iconName="loading"

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
  circle(cx="50" cy="50" fill="none" stroke="#28292f" stroke-width="10" r="35"
      stroke-dasharray="164.93361431346415 56.97787143782138" transform="rotate(306 50 50)")
    animateTransform(attributeName="transform" type="rotate" calcMode="linear" values="0 50 50;360 50 50"
        keyTimes="0;1" dur="1s" begin="0s" repeatCount="indefinite")
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-loading",
  props: {
    ...IconBase.props,
  },
}
</script>

<style lang="stylus" scoped>
</style>
