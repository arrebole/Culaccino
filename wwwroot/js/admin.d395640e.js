(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["admin"],{"0b6d":function(t,e,a){},"6b5e":function(t,e,a){"use strict";var n=a("74cc"),r=a.n(n);r.a},"74cc":function(t,e,a){},"85da":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"table"},[a("Header"),a("article",[a("Table",{attrs:{data:t.dir,deleteAt:t.deleteAt}})],1),a("Footer")],1)},r=[],i=(a("96cf"),a("3b8d")),s=a("2b0e"),c=a("0418"),o=a("fd2d"),l=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("section",[t._m(0),a("transition-group",{attrs:{name:"list",tag:"div"}},t._l(t.data,function(e){return a("div",{key:e.title,staticClass:"table-grid table-contents"},[a("div",{staticClass:"table-item"},[t._v(t._s(e.title))]),a("div",{staticClass:"table-item"},[t._v(t._s(new Date(e.CreatedAt).toLocaleString()))]),a("div",{staticClass:"table-item"},[t._v(t._s(e.views))]),a("div",{staticClass:"contrl-contents"},[a("router-link",{attrs:{to:{name:"Commit",params:{domain:e.author,repo:e.title}}}},[t._v("修改")]),t._v("/\n        "),a("span",{on:{click:function(a){return t.deleteAt({domain:e.author,repoName:e.title})}}},[t._v("删除")])],1)])}),0)],1)},d=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"table-grid table-header"},[a("div",{staticClass:"table-item"},[t._v("仓库")]),a("div",{staticClass:"table-item"},[t._v("时间")]),a("div",{staticClass:"table-item"},[t._v("浏量")]),a("div",{staticClass:"contrl"},[t._v("操作")])])}],u=s["default"].extend({props:["data","deleteAt"],methods:{}}),v=u,f=(a("6b5e"),a("2877")),m=Object(f["a"])(v,l,d,!1,null,"6902fe81",null),b=m.exports,p=a("79f6"),h=s["default"].extend({data:function(){return{dir:[]}},created:function(){this.fethData()},methods:{fethData:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,p["a"].getReposOwner(this.$route.params.domain);case 2:e=t.sent,e.data.dir&&(this.dir=e.data.dir.reverse());case 4:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}(),deleteAt:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(e){return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,p["a"].delRepo(e);case 2:t.sent,window.location.reload();case 4:case"end":return t.stop()}},t)}));function e(e){return t.apply(this,arguments)}return e}()},components:{Header:c["a"],Footer:o["a"],Table:b}}),_=h,w=(a("b560"),Object(f["a"])(_,n,r,!1,null,"3bdf20f2",null));e["default"]=w.exports},b560:function(t,e,a){"use strict";var n=a("0b6d"),r=a.n(n);r.a}}]);
//# sourceMappingURL=admin.d395640e.js.map