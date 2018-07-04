<template lang="pug">
section.g-v(:class="styleClass")
  slot
</template>

<script>
export default {
  name: "g-v",
  props: {
    wrap: {
      type: Boolean,
      default: false,
    },
    reverse: {
      type: Boolean,
      default: false,
    },
    jC: {
      type: String,
      default: "",
      validator: v => ["", "center", "flex-start", "flex-end", "space-around", "space-between", "space-evenly"].indexOf(v) > -1
    },
    aI: {
      type: String,
      default: "",
      validator: v => ["", "center", "flex-start", "flex-end", "baseline", "stretch"].indexOf(v) > -1
    },
    aC: {
      type: String,
      default: "",
      validator: v => ["", "center", "flex-start", "flex-end", "space-around", "space-between", "stretch"].indexOf(v) > -1
    },
  },
  computed: {
    styleClass() {
      let a = {
        "g-v-reverse": this.reverse,
        "g-flex-wrap": this.wrap,
      }
      if (this.aI !== "") {
        a["g-flex-align-items-" + this.aI] = true
      }
      if (this.jC !== "") {
        a["g-flex-justify-content-" + this.jC] = true
      }
      return a
    },
  }
}
</script>

<style lang="stylus" scoped>
.g-v
  display: flex
  flex-direction: column

.g-v-reverse
  flex-direction: column-reverse

.g-flex-wrap
  flex-wrap: wrap

/* flex align-items */
.g-flex-align-items-center
  align-items: center

.g-flex-align-items-flex-start
  align-items: flex-start

.g-flex-align-items-flex-end
  align-items: flex-end

/* flex justify-content */
.g-flex-justify-content-center
  justify-content: center

.g-flex-justify-content-flex-end
  justify-content: flex-end

.g-flex-justify-content-space-around
  justify-content: space-around

.g-flex-justify-content-space-between
  justify-content: space-between
</style>
