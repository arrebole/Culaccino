<template>
  <header>
    <nav class="header-font">
      <span class="iconfont icon-menu button" @click="menuClick"></span>
      <div class="breadcrumb">
        <template v-for="(pathName, index) in breadcrumbData">
          <span :key="pathName + 'span'" v-if="index != 0">/</span>
          <router-link
            class="breadcrumb-item button"
            :key="pathName"
            :to="{ name: pathName }"
            >{{ pathName }}</router-link
          >
        </template>
      </div>
    </nav>
    <transition name="fade">
      <div class="menu" v-show="isShowMenu" @click="menuClick">
        <div class="menu-title">
          <h4>Culaccino</h4>
          <span class="iconfont icon-jiantou menu-return"></span>
        </div>
        <ul>
          <router-link tag="li" :to="{ name: 'Home' }">
            <span class="iconfont icon-home menu-icon"></span>首页
          </router-link>
          <router-link tag="li" :to="{ name: 'Create' }">
            <span class="iconfont icon-admin menu-icon"></span>发布
          </router-link>
          <router-link tag="li" :to="{ name: 'Admin' }">
            <span class="iconfont icon-admin menu-icon"></span>管理
          </router-link>
        </ul>
      </div>
    </transition>
  </header>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      isShowMenu: false,
      path: window.location.pathname
    };
  },
  methods: {
    // string 数组首字母大写
    firstUpperCase(str: string): string {
      return str.toLowerCase().replace(/( |^)[a-z]/g, L => L.toUpperCase());
    },
    menuClick() {
      this.isShowMenu = !this.isShowMenu;
    }
  },
  computed: {
    breadcrumbData(): string[] {
      let result: string[] = this.path
        .split("/")
        .filter(s => s != "" && s > "9");
      result.unshift("home");
      return result.map(e => this.firstUpperCase(e));
    }
  }
});
</script>

<style lang="scss" scoped>
header {
  background-color: #24292e;
  box-shadow: 0px 0px 5px rgb(82, 82, 82);
  margin-bottom: 3px;
  padding: 0 16px;
}

header > nav {
  min-height: 48px;
  display: flex;
  align-items: center;
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

  color: #fff;
  &:hover {
    color: rgb(212, 212, 212);
  }
}
.header-font {
  font-weight: 500;
  color: rgb(255, 255, 255);
}
.menu-icon {
  padding: 0 3px;
}

.menu {
  color: rgb(255, 255, 255);
  background-color: #24292e;
  ul {
    list-style: none;
    padding: 0;
    li {
      padding: 15px 0;
      cursor: pointer;
      border-bottom: 1px solid hsla(0, 0%, 100%, 0.15);
      border-top: 1px solid hsla(0, 0%, 100%, 0.15);
      &:hover {
        color: rgb(199, 199, 199);
      }
    }
  }
}

@media screen and (min-width: 544px) {
  .menu {
    position: fixed;
    top: 0px;
    left: 0px;
    width: 200px;
    height: 100%;
    padding: 10px;
    box-shadow: 0 2px 15px rgb(99, 99, 99);
    .menu-title {
      display: flex;
      height: 80px;
      font-size: 1.5rem;
      padding: 0 15px;
      align-items: center;
      .menu-return {
        padding: 0 20px;
        &:hover {
          color: #79a783;
        }
      }
    }
  }
}
@media screen and (max-width: 544px) {
  .menu {
    height: 150px;
    font-size: 0.9rem;
    .menu-title {
      display: none;
    }
  }
}
</style>
