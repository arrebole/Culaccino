<template>
  <div>
    <Header />
    <article>
      <section>
        <Editor handle="update" :archive="repo"></Editor>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../../../components/Header.vue";
import Editor from "../../../components/Editor.vue";
import Footer from "../../../components/Footer.vue";
import api from "../../../api/index";
import IResp, { Repo } from "../../../types/resp";

interface Idata {
  repo: Repo | null;
}

export default Vue.extend({
  name: "Update",
  data(): Idata {
    return {
      repo: null
    };
  },
  created() {
    this.getData();
  },
  methods: {
    async getData() {
      let res: IResp = await api.getRepo({
        domain: this.$route.params.domain,
        repoName: this.$route.params.repo,
      });
      this.repo = res.data.repo;
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
