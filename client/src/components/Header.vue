<template>
  <header>
    <nav>
      <span class="iconfont icon-menu button"></span>
      <div class="breadcrumb">
        <template v-for="(pathName, index) in breadcrumbData">
          <span :key="pathName + 'span'" v-if="index != 0" class="header-font">/</span>
          <router-link class="breadcrumb-item button header-font" :to="{ name: pathName }" :key="pathName">{{ pathName }}</router-link>
        </template>
      </div>
    </nav>
  </header>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      path: window.location.pathname
    };
  },
  methods: {
    // string 数组首字母大写
    firstUpperCase(str: string): string {
      return str.toLowerCase().replace(/( |^)[a-z]/g, L => L.toUpperCase());
    }
  },
  computed: {
    breadcrumbData(): string[] {
      let result: string[] = this.path.split("/").filter(s => s != "");
      if (result.length == 0) result.push("home");
      return result.map(e => this.firstUpperCase(e));
    }
  }
});
</script>

<style lang="scss" scoped>
header {
  height: 48px;
  background-color: #fff;
  display: flex;
  align-items: center;
  box-shadow: 0px 0px 3px rgb(144, 199, 157);
  margin-bottom: 3px;
}

header > nav {
  padding-left: 16px;
}
.button {
  cursor: pointer;
}
.breadcrumb {
  display: inline;
  padding-left: 10px;
}
.breadcrumb-item {
  text-decoration: none;
  padding: 0 5px;
  &:hover {
    color: rgb(0, 0, 0);
  }
}
.header-font{
   font-weight: 500;
    color: rgba(73, 73, 73, 0.795);
}
</style>
