<template lang="pug">
section.g-h(:class="styleClass")
  slot
</template>

<script>
export default {
  name: "g-h",
  props: {
    wrap: {
      type: Boolean,
      default: false,
    },
    reverse: {
      type: Boolean,
      default: false,
    },
    jC: { // justify-content
      type: String,
      default: "",
      validator: v => ["", "center", "flex-start", "flex-end", "space-around", "space-between", "space-evenly"].indexOf(v) > -1
    },
    aI: { // align items
      type: String,
      default: "",
      validator: v => ["", "center", "flex-start", "flex-end", "baseline", "stretch"].indexOf(v) > -1
    },
    aC: { // align content
      type: String,
      default: "",
      validator: v => ["", "center", "flex-start", "flex-end", "space-around", "space-between", "stretch"].indexOf(v) > -1
    },
  },
  computed: {
    styleClass() {
      let a = {
        "g-h-reverse": this.reverse,
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
.g-h
  display: flex

.g-h-reverse
  flex-direction: row-reverse

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
