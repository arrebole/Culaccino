<template>
  <footer>
    <div class="footer_text">©2019 - <strong>Councurrent.Exec@gmail.com</strong></div>
  </footer>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({});
</script>


<style lang="scss" scoped>
footer {
  height: 66px;
  //background-color: #f5f5f5;
  background-color: #333;
  color: rgba(0, 0, 0, 0.87);
  box-shadow: 0px -1px 6px rgb(231, 231, 231);
}

.footer_text {
  color: rgb(240, 255, 255);
  font-size: 0.9rem;
  padding: 25px 30px;
}
</style>
