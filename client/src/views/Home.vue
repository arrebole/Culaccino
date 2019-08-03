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
  dir: IArticle[] | null;
}

export default Vue.extend({
  data(): Idata {
    return {
      dir: null
    };
  },
  created() {
    this.getdata();
  },
  methods: {
    async getdata() {
      let res: IResp = await api.getTable(0);
      this.dir = res.dir;
    },
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
  min-height: 80px;
  background: #eee;
  min-height: 820px;
}
</style>
