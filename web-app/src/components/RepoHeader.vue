<template>
  <div class="pagehead repohead">
    <div class="repohead-details">
      <h1 class="public">
        <span class="iconfont icon-repo path-divider"></span>
        <span>
          <a>{{ repo.author }}</a>
        </span>
        <span class="path-divider">/</span>
        <span>
          <a>{{ repo.title }}</a>
        </span>
      </h1>
    </div>

    <nav class="repohead-nav">
      <router-link :to="{ name:'Archive', params:$route.params.id }" class="reponav-item selected">
        <span class="iconfont icon-code icon-fontsize"></span>
        <span>Code</span>
      </router-link>
    </nav>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props:["repo"]
});
</script>

<style lang="scss" scoped>
.repohead {
  margin-bottom: 24px;
  padding-top: 24px;
}
.repohead-details {
  height: 28px;
  max-width: 980px;
  width: 90%;
  margin: auto;
  margin-bottom: 20px;
}
.repohead-nav {
  margin: auto;
  max-width: 980px;
  width: 90%;
  height: 40px;
  position: relative;
  top: 1px;
}
.icon-fontsize {
  font-size: 18px;
  padding: 0px 3px;
}
.reponav-item {
  float: left;
  height: 100%;
  color: rgb(0, 0, 0);
  text-decoration: none;
  font-size: 16px;
  padding: 7px 12px 8px;
  border: 1px solid transparent;
  border-top: 3px solid transparent;
  box-sizing: border-box;
}
.repohead-nav > .selected {
  background-color: #fff;
  border-color: #e36209 #e1e4e8 transparent;
  color: #24292e;
}
.public {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 600;
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
  color: #586069;
  .path-divider {
    padding: 0 3px;
  }
  span {
    a {
      color: #0366d6;
    }
  }
}
@media (min-width: 767px) {
  .pagehead {
    background-color: #fafbfc;
    border-bottom: 1px solid #e1e4e8;
  }
}
@media (max-width: 767px) {
  .pagehead {
    background-color: #fafbfc;
    border-bottom: 1px solid #e1e4e8;
  }
  .markdown-body {
    padding: 15px;
  }
}
</style>