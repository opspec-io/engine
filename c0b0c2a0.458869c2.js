(window.webpackJsonp=window.webpackJsonp||[]).push([[43],{169:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return a})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return l}));var n=r(1),c=r(9),o=(r(0),r(181)),a={title:"Socket Parameter [object]"},i={id:"opspec/reference/structure/op-directory/op/parameter/socket",title:"Socket Parameter [object]",description:"An object defining a parameter which accepts a [socket typed value](../../../../types/socket.md).",source:"@site/docs/opspec/reference/structure/op-directory/op/parameter/socket.md",permalink:"/docs/opspec/reference/structure/op-directory/op/parameter/socket",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/parameter/socket.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578700982,sidebar:"docs",previous:{title:"Object Parameter [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/parameter/object"},next:{title:"String Parameter [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/parameter/string"}},p=[{value:"Socket Properties:",id:"socket-properties",children:[{value:"description",id:"description",children:[]},{value:"isSecret",id:"issecret",children:[]}]}],s={rightToc:p},u="wrapper";function l(e){var t=e.components,r=Object(c.a)(e,["components"]);return Object(o.b)(u,Object(n.a)({},s,r,{components:t,mdxType:"MDXLayout"}),Object(o.b)("p",null,"An object defining a parameter which accepts a ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/types/socket"}),"socket typed value"),"."),Object(o.b)("h2",{id:"socket-properties"},"Socket Properties:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"must have:",Object(o.b)("ul",{parentName:"li"},Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"#description"}),"description")))),Object(o.b)("li",{parentName:"ul"},"may have:",Object(o.b)("ul",{parentName:"li"},Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"#issecret"}),"isSecret"))))),Object(o.b)("h3",{id:"description"},"description"),Object(o.b)("p",null,"A ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/markdown"}),"markdown")," string defining a human friendly description of the parameter."),Object(o.b)("h3",{id:"issecret"},"isSecret"),Object(o.b)("p",null,"A boolean indicating if the value of the parameter is secret. This will cause it to be hidden in UI's for example. "))}l.isMDXComponent=!0},181:function(e,t,r){"use strict";r.d(t,"a",(function(){return l})),r.d(t,"b",(function(){return m}));var n=r(0),c=r.n(n);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,c=function(e,t){if(null==e)return{};var r,n,c={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(c[r]=e[r]);return c}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(c[r]=e[r])}return c}var s=c.a.createContext({}),u=function(e){var t=c.a.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},l=function(e){var t=u(e.components);return c.a.createElement(s.Provider,{value:t},e.children)},b="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return c.a.createElement(c.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,o=e.originalType,a=e.parentName,s=p(e,["components","mdxType","originalType","parentName"]),l=u(r),b=n,f=l["".concat(a,".").concat(b)]||l[b]||d[b]||o;return r?c.a.createElement(f,i({ref:t},s,{components:r})):c.a.createElement(f,i({ref:t},s))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=r.length,a=new Array(o);a[0]=f;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[b]="string"==typeof e?e:n,a[1]=i;for(var s=2;s<o;s++)a[s]=r[s];return c.a.createElement.apply(null,a)}return c.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);