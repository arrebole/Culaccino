(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["create"],{"1d12":function(e,t,a){},"5f53":function(e,t,a){},b2ae:function(e,t,a){"use strict";var r=a("5f53"),c=a.n(r);c.a},be6c:function(e,t,a){"use strict";var r=a("1d12"),c=a.n(r);c.a},ceb0:function(e,t,a){"use strict";var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{attrs:{id:"editor"}},[a("div",{staticClass:"archive-info"},[a("div",[a("label",[e._v("Author")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.author,expression:"cache.author"}],attrs:{type:"text",placeholder:"作者"},domProps:{value:e.cache.author},on:{input:function(t){t.target.composing||e.$set(e.cache,"author",t.target.value)}}})]),a("div",[a("label",[e._v("Title")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.title,expression:"cache.title"}],attrs:{type:"text",placeholder:"标题"},domProps:{value:e.cache.title},on:{input:function(t){t.target.composing||e.$set(e.cache,"title",t.target.value)}}})]),a("div",[a("label",[e._v("target")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.target,expression:"cache.target"}],attrs:{type:"text",placeholder:"标签"},domProps:{value:e.cache.target},on:{input:function(t){t.target.composing||e.$set(e.cache,"target",t.target.value)}}})]),a("div",[a("label",[e._v("cover")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.cover,expression:"cache.cover"}],attrs:{type:"text",placeholder:"封面"},domProps:{value:e.cache.cover},on:{input:function(t){t.target.composing||e.$set(e.cache,"cover",t.target.value)}}})]),a("div",[a("label",[e._v("Summer")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.summary,expression:"cache.summary"}],attrs:{type:"text",placeholder:"摘要"},domProps:{value:e.cache.summary},on:{input:function(t){t.target.composing||e.$set(e.cache,"summary",t.target.value)}}})]),a("button",{staticClass:"btn",attrs:{type:"submit"},on:{click:e.submit}},[a("span",[e._v("提交")])])]),a("mavon-editor",{attrs:{id:"markdown-editor"},model:{value:e.cache.contents,callback:function(t){e.$set(e.cache,"contents",t)},expression:"cache.contents"}})],1)},c=[],n=(a("96cf"),a("3b8d")),i=a("2b0e"),o=a("79f6"),s=a("4dd8"),u=i["default"].extend({props:["handle","archive"],data:function(){return{cache:new s["b"]}},created:function(){},watch:{archive:function(){Object.assign(this.cache,this.archive)}},methods:{submit:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(){return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:"create"==this.handle?this.apiCreate():"update"==this.handle&&this.apiUpdate();case 1:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),apiCreate:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,o["a"].createArchive(this.cache);case 2:t=e.sent,"success"==t.code&&alert("成功");case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),apiUpdate:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,o["a"].updateArchive(this.$route.params.id,this.cache);case 2:t=e.sent,"success"==t.code&&alert("成功");case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()}}),l=u,d=(a("be6c"),a("2877")),h=Object(d["a"])(l,r,c,!1,null,"ff48d892",null);t["a"]=h.exports},d879:function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("Header"),a("article",[a("section",[a("Editor",{attrs:{handle:"create"}})],1)]),a("Footer")],1)},c=[],n=a("2b0e"),i=a("0418"),o=a("ceb0"),s=a("fd2d"),u=n["default"].extend({data:function(){return{}},methods:{},components:{Header:i["a"],Footer:s["a"],Editor:o["a"]}}),l=u,d=(a("b2ae"),a("2877")),h=Object(d["a"])(l,r,c,!1,null,"7bcdf6b3",null);t["default"]=h.exports}}]);
//# sourceMappingURL=create.e8a6aabf.js.map