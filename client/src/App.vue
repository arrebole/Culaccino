<template>
  <div id="app">
    <keep-alive include="Home">
      <router-view />
    </keep-alive>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({});
</script>

<style lang="scss">
.list-enter-active,
.list-leave-active {
  transition: all 1.5s;
}
.list-enter {
  /* .list-leave-active for below version 2.1.8 */
  opacity: 0;
  transform: translateY(30px);
}
.list-leave-to {
  //opacity: 0;
  //transform: translateX(-30px);
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
