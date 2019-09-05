<template>
  <section>
    <div class="table-grid table-header">
      <div class="table-item">仓库</div>
      <div class="table-item">时间</div>
      <div class="table-item">浏量</div>
      <div class="contrl">操作</div>
    </div>

    <transition-group name="list" tag="div">
      <div
        class="table-grid table-contents"
        v-for="item in data"
        :key="item.title"
      >
        <div class="table-item">{{ item.title }}</div>
        <div class="table-item">{{ new Date(item.CreatedAt).toLocaleString() }}</div>
        <div class="table-item">{{ item.views }}</div>
        <div class="contrl-contents">
          <router-link :to="{ name: 'Commit', params: { domain: item.author,repo: item.title } }"
            >修改</router-link
          >/
          <span @click="deleteAt({ domain:item.author, repoName:item.title })">删除</span>
        </div>
      </div>
    </transition-group>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props: ["data","deleteAt"],
  methods: {}
});
</script>

<style lang="scss" scoped>
.table-grid {
  width: 90%;
  padding: 5px 0;
  margin: 8px auto;

  font-size: 0.9rem;
  color: #474747;

  border-radius: 8px;
  border: 1px solid rgb(228, 228, 228);

  background-color: #fff;

  display: grid;
  padding: 10px;
  //grid-template-rows: 100px 100px 100px;
  grid-template-columns: 1fr 2fr 2fr 1fr 1fr;
}
.table-header {
  box-shadow: 0px 0px 3px rgb(0, 0, 0);
}
.table-contents {
  box-shadow: 0px 0px 3px rgb(233, 108, 36);
}
.table-item {
  font-family: "Roboto", "Helvetica Neue", Arial, sans-serif;
}
.contrl {
  color: rgb(95, 158, 160);
}

@mixin warn() {
  &:hover {
    color: rgb(247, 2, 2);
  }
}

.contrl-contents {
  a {
    text-decoration: none;
    color: #5c339e;
    @include warn();
  }
  span {
    @include warn();
  }
  cursor: pointer;
  color: #55334c;
}
</style>
