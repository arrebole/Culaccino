<template>
  <div class="home">
    <Header />
    <article v-loading="loading">
      <Reveal
        v-for="(item, index) in dir"
        :key="index"
        :date="tranTime(item.CreatedAt)"
        :article="item"
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
  count: number;
  dir: IArticle[] | null;
}

export default Vue.extend({
  name: "Home",
  data(): Idata {
    return {
      currPage: 0,
      count: 0,
      dir: null
    };
  },
  created() {
    this.getData();
  },
  computed: {
    loading() {
      if (!this.dir) return true;
      else return false;
    }
  },
  methods: {
    async getData() {
      this.dir = null;
      const res = await api.getTable(this.currPage);
      this.setData(res.dir, res.count);
    },
    setData(dir: IArticle[], count: number) {
      this.dir = dir;
      this.count = count;
    },
    currentChange(val: number) {
      this.currPage = val - 1;
      this.getData();
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
