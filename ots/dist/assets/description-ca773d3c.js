import{r as i,az as d,aG as m,aB as f,aE as E,aF as g,aD as h,R as x}from"./index-1a1a81e1.js";let p=i.createContext(null);function c(){let a=i.useContext(p);if(a===null){let t=new Error("You used a <Description /> component, but it is not inside a relevant parent.");throw Error.captureStackTrace&&Error.captureStackTrace(t,c),t}return a}function k(){let[a,t]=i.useState([]);return[a.length>0?a.join(" "):void 0,i.useMemo(()=>function(e){let o=h(r=>(t(s=>[...s,r]),()=>t(s=>{let n=s.slice(),u=n.indexOf(r);return u!==-1&&n.splice(u,1),n}))),l=i.useMemo(()=>({register:o,slot:e.slot,name:e.name,props:e.props}),[o,e.slot,e.name,e.props]);return x.createElement(p.Provider,{value:l},e.children)},[t])]}let v="p",C=d(function(a,t){let e=m(),{id:o=`headlessui-description-${e}`,...l}=a,r=c(),s=f(t);E(()=>r.register(o),[o,r.register]);let n={ref:s,...r.props,id:o};return g({ourProps:n,theirProps:l,slot:r.slot||{},defaultTag:v,name:r.name||"Description"})});export{C as F,k};
