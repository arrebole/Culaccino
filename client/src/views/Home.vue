<template>
  <div class="home">
    <Header />
    <article v-loading="loading">

        <Reveal
          v-for="(item, index) in dir"
          :key="index"
          :ID="item.ID"
          :date="tranTime(item.CreatedAt)"
          :views="item.views"
          :author="item.author"
          :title="item.title"
          :summary="item.summary"
          :cover = "item.cover"

        ></Reveal>
        <section class="pagination">
          <el-pagination
            :page-size="5"
            layout="prev, pager, next"
            :total="count"
            @current-change="currentChange"
            @size-change="sizeChange"
          >
          </el-pagination>
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
  currPage: number;
  loading:boolean;
  count: number;
  dir: IArticle[] | null;
}

export default Vue.extend({
  name: "Home",
  data(): Idata {
    return {
      loading:true,
      currPage: 0,
      count: 0,
      dir: null
    };
  },
  created() {
    this.getCount();
    this.getDir(this.currPage);
  },
  computed: {},
  methods: {
    async getCount() {
      let res = await api.count();
      this.count = res.count;
    },
    async getDir(page: number) {
      this.loading = true;
      let res: IResp = await api.getTable(page);
      this.loading = false;
      this.dir = res.dir;
    },
    currentChange(val: number) {
      this.getDir(val - 1);
    },
    sizeChange(val: number) {},
    tranTime(str: string) {
      return new Date(str).toLocaleDateString().replace(/\//g, "-");
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
  background: rgb(255, 255, 255);
  min-height: 800px;
  padding-top: 20px;
}
.pagination {
  margin: 20px;
}
</style>
