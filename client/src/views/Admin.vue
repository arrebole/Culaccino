<template>
  <div class="table">
    <Header />
    <article>
      <section class="table-grid table-header">
        <div class="table-item">ID</div>
        <div class="table-item">时间</div>
        <div class="table-item">标题</div>
        <div class="table-item">浏量</div>
        <div class="contrl">操作</div>
      </section>

      <transition-group name="list" tag="div">
        <section
          class="table-grid table-contents"
          v-for="item in dir"
          :key="item.ID"
        >
          <div class="table-item">{{ item.ID }}</div>
          <div class="table-item">{{ item.CreatedAt }}</div>
          <div class="table-item">{{ item.title }}</div>
          <div class="table-item">{{ item.views }}</div>
          <div class="contrl-contents">
            <router-link :to="{ name: 'Update', params: { id: item.ID } }"
              >修改</router-link
            >/
            <span @click="deleteAt(item.ID)">删除</span>
          </div>
        </section>
      </transition-group>
    </article>
    <Footer />
  </div>
</template>



<script lang="ts">
import Vue from "vue";
import IResp, { IArticle } from "../types/resp";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import api from "../api/index";

interface Data {
  dir: IArticle[];
}

export default Vue.extend({
  data(): Data {
    return {
      dir: []
    };
  },
  created() {
    this.fethData();
  },
  methods: {
    async fethData() {
      let res: IResp = await api.getAllTable();
      this.dir = res.dir.reverse();
    },
    async deleteAt(id: number) {
      let res: IResp = await api.delArticle(id);
      window.location.reload();
    }
  },
  components: {
    Header,
    Footer
  }
});
</script>

<style lang="scss" scoped>
article {
  min-height: 780px;
  padding: 20px;
  background: #fafafa;
}
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
