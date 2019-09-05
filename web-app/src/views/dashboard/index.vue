<template>
  <div class="home">
    <Header />
    <article v-loading="loading">
      
      <Reveal
        v-for="(item, index) in dir"
        :key="index"
        :archive="item"
      />

      <section class="pagination">
        <el-pagination
          layout="prev, pager, next"
          :page-size="per_page"
          :total="total"
          :current-page="currentPage"
          @current-change="currentChange"
          @size-change="sizeChange"
        ></el-pagination>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../../components/Header.vue";
import Footer from "../../components/Footer.vue";
import Reveal from "../../components/Reveal.vue";
import api from "../../api/index";
import IResp, { IArchive } from "../../types/resp";

interface Data {
  total: number;
  per_page: number; //每页显示数量
  dir: IArchive[];
}

export default Vue.extend({
  name: "Home",
  data(): Data {
    return {
      per_page: 5,
      total: 0,
      dir: []
    };
  },
  created() {
    this.getData();
  },
  watch: {
    '$route.query': 'getData'
  },
  computed: {
    loading(): boolean {
      if (!this.dir) return true;
      else return false;
    },
    currentPage(): number {
      if (this.$route.query.page && typeof this.$route.query.page == "string") {
        return parseInt(this.$route.query.page);
      }
      return 1;
    }
  },
  methods: {
    async getData() {
      this.dir = [];
      this.setData(await api.getDashboard(this.currentPage - 1, this.per_page));
    },
    setData({ data }: IResp) {
      this.dir = data.dir;
      this.total = data.count.total;
    },
    currentChange(val: number) {
      this.$router.push({ name: "Home", query: { page: `${val}` } });
    },
    sizeChange(val: number) {},
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
