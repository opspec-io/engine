(window.webpackJsonp=window.webpackJsonp||[]).push([[40],{133:function(e,r,t){"use strict";t.r(r),t.d(r,"frontMatter",(function(){return a})),t.d(r,"metadata",(function(){return i})),t.d(r,"rightToc",(function(){return p})),t.d(r,"default",(function(){return b}));var n=t(1),c=t(6),o=(t(0),t(147)),a={title:"Variable Reference [string]"},i={id:"opspec/reference/structure/op-directory/op/variable-reference",title:"Variable Reference [string]",description:"A string referencing the location of/for a value in the form of `$(REFERENCE)` where `REFERENCE` MUST start with an [identifier [string]](identifier.md) and MAY end with one or more:",source:"@site/docs/opspec/reference/structure/op-directory/op/variable-reference.md",permalink:"/docs/opspec/reference/structure/op-directory/op/variable-reference",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/variable-reference.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1586986005,sidebar:"docs",previous:{title:"Markdown [string]",permalink:"/docs/opspec/reference/structure/op-directory/op/markdown"},next:{title:"Array",permalink:"/docs/opspec/reference/types/array"}},p=[],l={rightToc:p},s="wrapper";function b(e){var r=e.components,t=Object(c.a)(e,["components"]);return Object(o.b)(s,Object(n.a)({},l,t,{components:r,mdxType:"MDXLayout"}),Object(o.b)("p",null,"A string referencing the location of/for a value in the form of ",Object(o.b)("inlineCode",{parentName:"p"},"$(REFERENCE)")," where ",Object(o.b)("inlineCode",{parentName:"p"},"REFERENCE")," MUST start with an ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/identifier"}),"identifier [string]")," and MAY end with one or more:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/types/array#item-referencing"}),"array item references")),Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/types/object#property-referencing"}),"object property references")),Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/types/dir#entry-referencing"}),"dir entry references"))),Object(o.b)("p",null,"References can be used to either define or access values in the current scope. "),Object(o.b)("p",null,"When an op starts, it's initial scope includes:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},Object(o.b)("inlineCode",{parentName:"li"},"/")," with a value of the current op directory i.e. the current op's ",Object(o.b)("inlineCode",{parentName:"li"},"op.yml")," can be accessed via ",Object(o.b)("inlineCode",{parentName:"li"},"$(/op.yml)"),"."),Object(o.b)("li",{parentName:"ul"},"any defined inputs")),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"note: variable references can be escaped by prefixing the ","[would be]"," variable reference with ",Object(o.b)("inlineCode",{parentName:"p"},"\\")," i.e. ",Object(o.b)("inlineCode",{parentName:"p"},"\\$(wouldBeVariableReference)")," would not be treated as a variable reference. ")))}b.isMDXComponent=!0},147:function(e,r,t){"use strict";t.d(r,"a",(function(){return b})),t.d(r,"b",(function(){return m}));var n=t(0),c=t.n(n);function o(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function a(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function i(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?a(Object(t),!0).forEach((function(r){o(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function p(e,r){if(null==e)return{};var t,n,c=function(e,r){if(null==e)return{};var t,n,c={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(c[t]=e[t]);return c}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(c[t]=e[t])}return c}var l=c.a.createContext({}),s=function(e){var r=c.a.useContext(l),t=r;return e&&(t="function"==typeof e?e(r):i({},r,{},e)),t},b=function(e){var r=s(e.components);return c.a.createElement(l.Provider,{value:r},e.children)},u="mdxType",f={inlineCode:"code",wrapper:function(e){var r=e.children;return c.a.createElement(c.a.Fragment,{},r)}},d=Object(n.forwardRef)((function(e,r){var t=e.components,n=e.mdxType,o=e.originalType,a=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),b=s(t),u=n,d=b["".concat(a,".").concat(u)]||b[u]||f[u]||o;return t?c.a.createElement(d,i({ref:r},l,{components:t})):c.a.createElement(d,i({ref:r},l))}));function m(e,r){var t=arguments,n=r&&r.mdxType;if("string"==typeof e||n){var o=t.length,a=new Array(o);a[0]=d;var i={};for(var p in r)hasOwnProperty.call(r,p)&&(i[p]=r[p]);i.originalType=e,i[u]="string"==typeof e?e:n,a[1]=i;for(var l=2;l<o;l++)a[l]=t[l];return c.a.createElement.apply(null,a)}return c.a.createElement.apply(null,t)}d.displayName="MDXCreateElement"}}]);