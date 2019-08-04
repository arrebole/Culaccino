<template>
  <div class="reveal">
    <div class="reveal-header">
      <router-link :to="{ name: 'Article', params: { id: ID } }">
          {{ title }}</router-link>
    </div>
    <div class="reveal-contents">
      <p>{{ summary }}</p>
    </div>
    <div class="reveal-footer">
      <div>#{{ author }}</div>
      <div class="reveal-footer-info">
        <span class="iconfont icon-liulan"></span><span>{{ views }}</span>
        <span class="iconfont icon-icon-test1"></span>
        <span class="reveal-footer-info-time">{{ date }} </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props: ["title", "summary", "views", "date", "author", "ID"]
});
</script>


<style lang="scss" scoped>
.reveal {
  max-width: 900px;
  width: 80%;
  padding: 20px;
  margin: 10px;
  background-color: #fff;
  box-shadow: 1px 1px 5px rgb(141, 141, 141);
  border: 1px solid rgb(177, 177, 177);
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
}
.reveal-header {
  a {
    font-size: 1.6rem;
    color: rgb(24, 24, 24);
    text-decoration: none;
    &:hover {
      color: rgb(72, 95, 194);
    }
  }
}
.reveal-contents {
  p {
    font-size: 0.9rem;
    color: rgb(114, 114, 114);
  }
  border-bottom: 1px solid rgb(177, 177, 177);
  min-height: 40px;
}
.reveal-footer {
  display: flex;
  font-size: 0.9rem;
  margin-top: 5px;
  justify-content: space-between;
  color: rgb(122, 122, 122);
}
.reveal-footer-info {
  display: flex;
  span {
    padding: 0 2px;
  }
}
.reveal-footer-info-time {
  padding: 0 5px;
  color: rgb(71, 71, 71);
}
</style>
