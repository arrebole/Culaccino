<template>
  <div class="contents" v-html="markdown"></div>
</template>

<script lang="ts">
import Vue from "vue";

import "highlight.js/styles/tomorrow.css";
import "github-markdown-css"

import marked from "marked";
import hljs from "highlight.js";

export default Vue.extend({
  props: ["contents"],
  computed: {
    markdown() {
      marked.setOptions({
        highlight: function(code:any, lang:any) {
          return hljs.highlight(lang, code).value;
        }
      });
      return marked(this.contents);
    }
  }
});
</script>

