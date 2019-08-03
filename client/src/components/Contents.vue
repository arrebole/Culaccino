<template>
    <div class="contents" v-html="markdown"></div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props:["md"],
    computed:{
        markdown(){
            return this.md;
        }
    }
})
</script>

<style lang="scss" scoped>

</style>
