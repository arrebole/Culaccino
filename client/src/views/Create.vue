<template>
  <div>
    <Header />
    <article>
      <section>
        <Editor handle='create'></Editor>
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


export default Vue.extend({
  data() {
    return{}
  },
  methods: {
 
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
