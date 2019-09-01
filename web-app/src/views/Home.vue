<template>
  <div class="home">
    <Header />
    <article v-loading="loading">
      <Reveal
        v-for="(item, index) in dir"
        :key="index"
        :date="tranTime(item.CreatedAt)"
        :archive="item"
      ></Reveal>
      <section class="pagination">
        <el-pagination
          :page-size="per_page"
          layout="prev, pager, next"
          :total="total"
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
import IResp, { IArchive } from "../types/resp";

interface Idata {
  currPage: number;
  total: number;
  per_page:number;
  dir: IArchive[] | null;
}

export default Vue.extend({
  name: "Home",
  data(): Idata {
    return {
      currPage: 0,
      per_page:5,
      total: 0,
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
      this.setData(await api.getDashboard(this.currPage, this.per_page));
    },
    setData({data}: IResp) {
      this.dir = data.dir;
      this.total = data.count.total;
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
