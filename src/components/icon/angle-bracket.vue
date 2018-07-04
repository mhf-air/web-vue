<template lang="pug">
IconBase(
    iconName="angle-bracket"

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
      v-if="dir === 'left'"
      points="70 10, 30 50, 70 90"
      )
  polyline(
      v-else-if="dir === 'right'"
      points="30 10, 70 50, 30 90"
      )
  polyline(
      v-else-if="dir === 'top'"
      points="10 70, 50 30, 90 70"
      )
  polyline(
      v-else-if="dir === 'bottom'"
      points="10 30, 50 70, 90 30"
      )
</template>

<script>
import IconBase from "./aa-icon-base.vue"

export default {
  name: "icon-angle-bracket",
  props: {
    ...IconBase.props,
    direction: {
      type: String,
      default: "left", // left, right, top, bottom
    },
  },
  data() {
    return {
      dir: this.direction,
    }
  },
}
</script>

<style lang="stylus" scoped>
</style>
