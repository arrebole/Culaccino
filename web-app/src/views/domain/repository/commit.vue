<template>
  <div>
    <Header />
    <article>
      <section>
        <Editor handle="update" :archive="archive"></Editor>
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
import IResp, { IArchive, IArchiveBase } from "../../../types/resp";

interface Idata {
  archive: IArchive | null;
}

export default Vue.extend({
  name: "Update",
  data(): Idata {
    return {
      archive: null
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
      this.archive = res.data.archive;
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
