<template>
  <div id="app">
    <header>
      <nav>
        <div>
          <span class="iconfont icon-menu"></span>
        </div>

        <div class="nav-bread">
          <span>
            
          </span>
        </div>
      </nav>
    </header>

    <article>
      <router-view />
    </article>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      path: window.location.pathname
    };
  },

  computed: {
    // 监听路由变化
    routeList(): string[] {
      return this.path
    }
  }
});
</script>

<style lang="scss">
#app > header {
  display: flex;
  align-items: center;
  height: 60px;
  background-color: rgb(103, 175, 209);
}
header > nav > div {
  font-size: 26px;
  line-height: 0;
  display: inline-block;
  padding: 0 10px;
  color: rgb(241, 250, 250);
}
</style>
