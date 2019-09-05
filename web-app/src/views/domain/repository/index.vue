<template>
  <div>
    <Header />
    <article v-loading="loading">
      <section class="repository-header">
        <RepoHeader :repo="repo"></RepoHeader>
      </section>
      
      
      <section class="repository-content">
        <RepoContents :repo="repo"></RepoContents>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../../../components/Header.vue";
import Footer from "../../../components/Footer.vue";
import RepoContents from "../../../components/RepoContents.vue";
import RepoHeader from "../../../components/RepoHeader.vue";
import api from "../../../api/index";
import IResp, { IArchive, Archive } from "../../../types/resp";

export default Vue.extend({
  name: "repository",
  data(): { repo: IArchive; loading: boolean } {
    return {
      repo: new Archive(),
      loading: true
    };
  },
  created() {
    this.getData();
  },
  watch: {
    "this.$router": "getData"
  },
  methods: {
    repoSymble() {
      if (typeof this.$route.query.id == "number") {
        return this.$route.query.id
      }
      if(typeof this.$route.query.id == "string"){
        return parseInt(this.$route.query.id)
      }
      return -1;
    },
    async getData() {
      const res: IResp = await api.getRepo(this.repoSymble());
      if (res.code >= 0 && res.data.archive != null) {
        this.repo = res.data.archive;
      }
      this.loading = false;
    }
  },
  components: {
    Header,
    Footer,
    RepoHeader,
    RepoContents
  }
});
</script>

<style lang="scss" scoped>
article {
  min-height: 820px;
  background-color: #fff;
  padding-bottom: 30px;
}
.repository-content {
  width: 90%;
  max-width: 980px;
  margin: auto;
}

</style>
