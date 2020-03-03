(window.webpackJsonp=window.webpackJsonp||[]).push([[51],{178:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return i})),r.d(t,"metadata",(function(){return o})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return s}));var n=r(1),c=r(9),a=(r(0),r(181)),i={title:"Predicate [object]"},o={id:"opspec/reference/structure/op-directory/op/call/predicate",title:"Predicate [object]",description:"An object defining a condition which evaluates to true or false.",source:"@site/docs/opspec/reference/structure/op-directory/op/call/predicate.md",permalink:"/docs/opspec/reference/structure/op-directory/op/call/predicate",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/call/predicate.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578700982,sidebar:"docs",previous:{title:"Parallel Loop Call [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/parallel-loop"},next:{title:"Pull Creds [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/pull-creds"}},l=[{value:"Properties",id:"properties",children:[{value:"eq",id:"eq",children:[]},{value:"exists",id:"exists",children:[]},{value:"ne",id:"ne",children:[]},{value:"notExists",id:"notexists",children:[]}]}],p={rightToc:l},u="wrapper";function s(e){var t=e.components,r=Object(c.a)(e,["components"]);return Object(a.b)(u,Object(n.a)({},p,r,{components:t,mdxType:"MDXLayout"}),Object(a.b)("p",null,"An object defining a condition which evaluates to true or false."),Object(a.b)("h2",{id:"properties"},"Properties"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"must have exactly one of",Object(a.b)("ul",{parentName:"li"},Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"#eq"}),"eq")),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"#exists"}),"exists")),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"#ne"}),"ne")),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"#notexists"}),"notExists"))))),Object(a.b)("h3",{id:"eq"},"eq"),Object(a.b)("p",null,"An array defining a predicate, true when all items are equal."),Object(a.b)("p",null,"Items:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"must be one of",Object(a.b)("ul",{parentName:"li"},Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/structure/op-directory/op/reference"}),"reference")," string"),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/structure/op-directory/op/initializer"}),"initializer"))))),Object(a.b)("h3",{id:"exists"},"exists"),Object(a.b)("p",null,"A ",Object(a.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/reference"}),"reference")," string defining a predicate, true when the referenced value exists."),Object(a.b)("h3",{id:"ne"},"ne"),Object(a.b)("p",null,"An array defining a predicate, true when one or more items aren't equal."),Object(a.b)("p",null,"Items:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"must be",Object(a.b)("ul",{parentName:"li"},Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/structure/op-directory/op/reference"}),"reference")," string"),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/structure/op-directory/op/initializer"}),"initializer"))))),Object(a.b)("h3",{id:"notexists"},"notExists"),Object(a.b)("p",null,"A ",Object(a.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/reference"}),"reference")," string defining a predicate, true when the referenced value doesn't exist."))}s.isMDXComponent=!0},181:function(e,t,r){"use strict";r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return O}));var n=r(0),c=r.n(n);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,c=function(e,t){if(null==e)return{};var r,n,c={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(c[r]=e[r]);return c}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(c[r]=e[r])}return c}var p=c.a.createContext({}),u=function(e){var t=c.a.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):o({},t,{},e)),r},s=function(e){var t=u(e.components);return c.a.createElement(p.Provider,{value:t},e.children)},b="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return c.a.createElement(c.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,a=e.originalType,i=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),s=u(r),b=n,f=s["".concat(i,".").concat(b)]||s[b]||d[b]||a;return r?c.a.createElement(f,o({ref:t},p,{components:r})):c.a.createElement(f,o({ref:t},p))}));function O(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var a=r.length,i=new Array(a);i[0]=f;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o[b]="string"==typeof e?e:n,i[1]=o;for(var p=2;p<a;p++)i[p]=r[p];return c.a.createElement.apply(null,i)}return c.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);