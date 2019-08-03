<template>
  <div>
    <Header />
    <article>
      <!-- <section></section> -->
      <section class="cent">
        <Contents :md="article" class="markdown-body"></Contents>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import Contents from "../components/Contents.vue";
import api from "../api/index";
import IResp from "../types/resp";
export default Vue.extend({
  data() {
    return {
      article: ""
    };
  },
  created() {
    this.getData();
  },
  methods: {
    async getData() {
      const id = parseInt(this.$route.params.id);
      const res: IResp = await api.getArticle(id);
      this.article = res.articles.contents;
    }
  },
  components: {
    Header,
    Footer,
    Contents
  }
});
</script>

<style lang="scss" scoped>
article {
  min-height: 820px;
}
.cent {
  padding: 10px;
}
.markdown-body {
  box-sizing: border-box;
  min-width: 200px;
  max-width: 980px;
  margin: 0 auto;
  padding: 45px;
}

@media (max-width: 767px) {
  .markdown-body {
    padding: 15px;
  }
}
</style>
