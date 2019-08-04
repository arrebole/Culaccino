<template>
  <div class="home">
    <Header />
    <article>
      <Reveal
        v-for="(item, index) in dir"
        :key="index"
        :ID="item.ID"
        :date="tranTime(item.CreatedAt)"
        :views="item.views"
        :author="item.author"
        :title="item.title"
        :summary="item.summary"
      ></Reveal>
      <section class="pagination">
        <div @click="rePage(-1)">上一页</div>
        <div>{{ page + 1 }}</div>
        <div @click="rePage(1)">下一页</div>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import Reveal from "../components/Reveal.vue";
import api from "../api/index";
import IResp, { IArticle } from "../types/resp";

interface Idata {
  page: number;
  final: boolean;
  dir: IArticle[] | null;
}

export default Vue.extend({
  data(): Idata {
    return {
      page: 0,
      final: false,
      dir: null
    };
  },
  created() {
    this.getDir();
  },
  methods: {
    async getDir() {
      let res: IResp = await api.getTable(this.page);
      if (res.dir.length == 0) this.final = true;
      else this.dir = res.dir;
    },
    tranTime(str: string) {
      return new Date(str).toLocaleDateString().replace(/\//g, "-");
    },
    isRepage(n: number) :boolean{
      return this.page + n < 0 || (this.final && n > 0);
    },
    rePage(n: number) {
      if (this.isRepage(n)) return;
      this.page += n;
      this.getDir();
    }
  },
  components: {
    Header,
    Footer,
    Reveal
  }
});
</script>>

<style lang="scss" scoped>
article {
  display: flex;
  align-items: center;
  flex-direction: column;
  min-height: 80px;
  background: rgb(248, 248, 248);
  min-height: 820px;
  padding-top: 20px;
}
.pagination {
  display: flex;
  max-width: 900px;
  width: 80%;
  height: 40px;
  justify-content: space-between;
  color: rgb(115, 48, 192);
  div {
    cursor: pointer;
    &:hover {
      color: rgb(255, 115, 0);
    }
  }
}
</style>
