(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["admin"],{3530:function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"table"},[a("Header"),a("article",[a("Table",{attrs:{data:t.dir,deleteAt:t.deleteAt}})],1),a("Footer")],1)},r=[],i=(a("96cf"),a("3b8d")),s=a("2b0e"),c=a("0418"),l=a("fd2d"),d=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("section",[t._m(0),a("transition-group",{attrs:{name:"list",tag:"div"}},t._l(t.data,function(e){return a("div",{key:e.ID,staticClass:"table-grid table-contents"},[a("div",{staticClass:"table-item"},[t._v(t._s(e.ID))]),a("div",{staticClass:"table-item"},[t._v(t._s(e.CreatedAt))]),a("div",{staticClass:"table-item"},[t._v(t._s(e.title))]),a("div",{staticClass:"table-item"},[t._v(t._s(e.views))]),a("div",{staticClass:"contrl-contents"},[a("router-link",{attrs:{to:{name:"Update",params:{articleID:e.ID}}}},[t._v("修改")]),t._v("/\n        "),a("span",{on:{click:function(a){return t.deleteAt(e.ID)}}},[t._v("删除")])],1)])}),0)],1)},o=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"table-grid table-header"},[a("div",{staticClass:"table-item"},[t._v("ID")]),a("div",{staticClass:"table-item"},[t._v("时间")]),a("div",{staticClass:"table-item"},[t._v("标题")]),a("div",{staticClass:"table-item"},[t._v("浏量")]),a("div",{staticClass:"contrl"},[t._v("操作")])])}],u=s["default"].extend({props:["data","deleteAt"],methods:{}}),v=u,f=(a("dc75"),a("2877")),b=Object(f["a"])(v,d,o,!1,null,"7b72b580",null),m=b.exports,p=a("79f6"),_=s["default"].extend({data:function(){return{dir:[]}},created:function(){this.fethData()},methods:{fethData:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,p["a"].getAllTable();case 2:e=t.sent,this.dir=e.dir.reverse();case 4:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}(),deleteAt:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(e){return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,p["a"].delArticle(e);case 2:t.sent,window.location.reload();case 4:case"end":return t.stop()}},t)}));function e(e){return t.apply(this,arguments)}return e}()},components:{Header:c["a"],Footer:l["a"],Table:m}}),h=_,w=(a("5aa9"),Object(f["a"])(h,n,r,!1,null,"2ef4233a",null));e["default"]=w.exports},"5aa9":function(t,e,a){"use strict";var n=a("9739"),r=a.n(n);r.a},"84c4":function(t,e,a){},9739:function(t,e,a){},dc75:function(t,e,a){"use strict";var n=a("84c4"),r=a.n(n);r.a}}]);
//# sourceMappingURL=admin.d3a95064.js.map