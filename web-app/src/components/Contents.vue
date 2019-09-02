<template>
  <div class="Box">
    <div class="Box-header">
      <h3>README.md</h3>
    </div>
    <div class="Box-body">
      <div class="markdown-body" v-html="markdown"></div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

import "highlight.js/styles/tomorrow.css";
import "github-markdown-css";

import marked from "marked";
import hljs from "highlight.js";

export default Vue.extend({
  props: ["contents"],
  computed: {
    markdown() {
      marked.setOptions({
        highlight: function(code: any, lang: any) {
          return hljs.highlight(lang, code).value;
        }
      });
      return marked(this.contents);
    }
  }
});
</script>

<style lang="scss" scoped>
.Box {
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
  border: 1px solid #d1d5da;
  border-radius: 3px;
}
.Box-header {
  background-color: #f6f8fa;
  border: 1px solid #d1d5da;
  padding: 8px 16px;
  border-top-left-radius: 3px;
  border-top-right-radius: 3px;
  margin: -1px -1px 0;

  h3 {
    margin: 0;
    font-size: 14px;
    font-weight: 600;
  }
}
.Box-body {
  padding: 8px 16px;
  border-bottom-left-radius: 2px;
  border-bottom-right-radius: 2px;
  border-bottom: 1px solid #e1e4e8;
}
.markdown-body {
  padding: 32px;
}
</style>