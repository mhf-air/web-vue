<template lang="pug">
IconBase(
    iconName="cross"

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
  circle(cx="50" cy="50" r="40")
  line(x1="30" y1="30" x2="70" y2="70" stroke-width="13" :stroke="crossColor")
  line(x1="30" y1="70" x2="70" y2="30" stroke-width="13" :stroke="crossColor")
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-cross",
  props: {
    ...IconBase.props,
    crossColor: {
      type: String,
      default: "white",
    },
  },
  data() {
    return {
    }
  },
}
</script>

<style lang="stylus" scoped>
</style>
