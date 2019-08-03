<template>
  <div>
    <Header />
    <article>
      <section>
        <Editor handle="update" :article="article"></Editor>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../components/Header.vue";
import Editor from "../components/Editor.vue";
import Footer from "../components/Footer.vue";
import api from "../api/index";
import IResp, { IArticleBase, IArticle } from "../types/resp";

interface Idata {
  article: IArticle | null;
}

export default Vue.extend({
  name: "Update",
  data(): Idata {
    return {
      article: null
    };
  },
  created() {
    this.getData();
  },
  methods: {
    async getData() {
      const id: number = parseInt(this.$route.params.articleID);
      let res: IResp = await api.getArticle(id);
      this.article = res.articles;
    }
  },
  components: {
    Header,
    Footer,
    Editor
  }
});
</script>

<style lang="scss" scoped>
article {
  min-height: 820px;
  display: flex;
}
section {
  flex: 1;
  padding: 30px;
}
</style>
