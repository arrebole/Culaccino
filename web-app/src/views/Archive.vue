<template>
  <div>
    <Header />
    <article v-loading="loading">

      <section class="repository-header">
        <RepoHeader :repo="archive"></RepoHeader>
      </section>


      <section class="repository-content">
        <RepoContents :repo="archive"></RepoContents>
      </section>

    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import RepoContents from "../components/RepoContents.vue";
import RepoHeader from "../components/RepoHeader.vue";
import api from "../api/index";
import IResp, { IArchive, Archive } from "../types/resp";

export default Vue.extend({
  data(): { archive: IArchive } {
    return {
      archive: new Archive()
    };
  },
  created() {
    this.getData();
  },
  computed: {
    paramsID() {
      return parseInt(this.$route.params.id);
    },
    loading() {
      if (this.archive.ID < 0) return true;
      return false;
    }
  },
  methods: {
    async getData() {
      const res: IResp = await api.getArchive(this.paramsID);
      this.archive = res.data.archive;
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
.repository-content{
  width: 100%;
  display: flex;
  justify-content: center;
}
</style>
