import{j as e,r as m,m as c,P as i}from"./index-1a1a81e1.js";import{G as l,f as n}from"./contractMatchParsers-fecf984b.js";import{D as o}from"./DecoratedAddressLink-fc50b0da.js";import{B as d}from"./BlockLink-7a655877.js";import{A as p}from"./AddressHighlighter-79fb858b.js";import{a as x}from"./useUIHooks-890ee51b.js";import"./StandardFrame-33bf0827.js";import"./ContentFrame-677dbbad.js";import"./StandardSelectionBoundary-204c519a.js";import"./StandardTBody-bb7ced43.js";import"./messages-e142d7a5.js";import"./useErigonHooks-c9b7ed80.js";import"./react-content-loader.es-44cea70a.js";const h=()=>e.jsxs(e.Fragment,{children:[e.jsx("th",{className:"w-96",children:"Address"}),e.jsx("th",{className:"w-28",children:"Block"}),e.jsx("th",{className:"w-40",children:"Age"}),e.jsx("th",{children:"Implementation"})]}),j=(s,t)=>({blockNumber:s.blockNumber,timestamp:t.get(s.blockNumber).timestamp,address:s.address,implementation:s.implementation}),g=({blockNumber:s,timestamp:t,address:a,implementation:r})=>e.jsxs(e.Fragment,{children:[e.jsx("td",{children:e.jsx(o,{address:a,eoa:!1,plain:!0})}),e.jsx("td",{children:e.jsx(d,{blockTag:s})}),e.jsx("td",{children:e.jsx(c,{timestamp:t})}),e.jsx("td",{className:"inline-flex",children:e.jsx(p,{address:r,children:e.jsx(o,{address:r,eoa:!1})})})]}),f=m.memo(g),G=()=>{const{pageNumber:s,page:t,total:a}=x("ERC1167",n,j);return document.title="ERC1167 Contracts | Otterscan",e.jsx(l,{title:"ERC1167 contracts",header:e.jsx(h,{}),cols:4,pageNumber:s,pageSize:i,total:a,page:t,Item:r=>e.jsx(f,{...r})})};export{G as default};
