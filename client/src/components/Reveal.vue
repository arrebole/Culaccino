<template>
  <div class="reveal">
    <div class="reveal-header">
      <router-link
        class="reveal-header-title"
        tag="h1"
        :to="{ name: 'Archives', params: { id: archive.ID } }"
      >
        {{ archive.title }}</router-link
      >
    </div>
    <div class="cover-container cover-container-auto">
      <img class="reveal-cover" :src="archive.cover" />
    </div>

    <div class="reveal-contents">
      <p>{{ archive.summary }}</p>
    </div>
    <div class="reveal-footer">
      <div>{{ archive.target }}</div>
      <div class="reveal-footer-info">
        <span class="iconfont icon-liulan"></span><span>{{ archive.views }}</span>
        <span class="iconfont icon-icon-test1"></span>
        <span class="reveal-footer-info-time">{{ date }} </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props: ["archive","date"]
});
</script>


<style lang="scss" scoped>
.reveal {
  max-width: 820px;
  width: 85%;
  padding: 20px;
  margin: 15px;
  background-color: #fff;
  box-shadow: 1px 1px 5px rgb(141, 141, 141);
  border: 1px solid rgb(177, 177, 177);
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
}
.reveal-header {
  padding: 10px;
  font-family: "Microsoft Yahei", Helvetica, Arial, sans-serif;
  h1 {
    color: rgb(24, 24, 24);
    text-decoration: none;
    text-align: center;
    cursor: pointer;
    &:hover {
      color: rgb(72, 95, 194);
    }
  }
}
.reveal-contents {
  color: #666666;
  background-color: #ecf0f1;
  border: 1px solid #ecf0f1;
  padding: 15px 15px;
  margin: 10px 0;
  border-radius: 4px;
  p {
    margin: 0;
    font-size: 0.9rem;
    color: rgb(114, 114, 114);
  }
}
.reveal-footer {
  display: flex;
  font-size: 0.9rem;
  margin-top: 5px;
  padding: 0 5px;
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

.cover-container {
  width: 100%;
  border-radius: 4px;
  overflow: hidden;
  //box-shadow: 0px 0px 5px rgb(190, 190, 190);
}
.reveal-cover {
  width: 100%;
}

@media screen and (min-width: 996px) {
  .reveal-header-title {
    margin: 10px 0;
  }
}

@media screen and (max-width: 996px) and(min-width: 544px) {
  .reveal-header-title {
    margin: 0;
  }
}
@media screen and (max-width: 544px) {
  .reveal-header-title {
    margin: 0;
  }

}
</style>
