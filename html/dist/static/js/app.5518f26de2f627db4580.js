webpackJsonp([1],{0:function(t,e){},DeKW:function(t,e){},NHnr:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var n=a("7+uW"),o={render:function(){var t=this.$createElement;return(this._self._c||t)("router-view")},staticRenderFns:[]},r=a("VU/8")(null,o,!1,null,null,null).exports,s=a("/ocq"),i={data:function(){return{user:{username:"",password:""}}},methods:{logind:function(){var t=this;this.$http.post("/api/login",this.user,{"Content-Type":"application/json"}).then(function(e){console.log("go success"),200==e.status&&(window.sessionStorage.setItem("token",e.data.token),t.$router.push({path:"/main"}))},function(e){t.user.password="",t.$message.error(e.data.error)})}}},l={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"login-back"},[a("div",{staticClass:"login-container"},[a("el-form",{ref:"user",attrs:{model:t.user}},[a("el-form-item",{staticStyle:{"text-align":"center"}},[a("h1",[t._v("用户登录")])]),t._v(" "),a("el-form-item",{staticStyle:{"text-align":"center"},attrs:{prop:"username"}},[a("el-input",{attrs:{type:"text",autofocus:"true","auto-complete":"off",placeholder:"用户名"},model:{value:t.user.username,callback:function(e){t.$set(t.user,"username",e)},expression:"user.username"}})],1),t._v(" "),a("el-form-item",[a("el-input",{attrs:{type:"password",autofocus:"true","auto-complete":"off",placeholder:"密码"},model:{value:t.user.password,callback:function(e){t.$set(t.user,"password",e)},expression:"user.password"}})],1),t._v(" "),a("el-form-item",{staticStyle:{"text-align":"center"}},[a("el-button",{staticStyle:{width:"150px","margin-top":"10px"},attrs:{type:"primary"},on:{click:t.logind}},[t._v("登录")])],1)],1)],1)])},staticRenderFns:[]};var u=a("VU/8")(i,l,!1,function(t){a("DeKW")},null,null).exports,c=a("mvHQ"),p=a.n(c),d=a("XLwt"),m=a.n(d);function f(t,e){var a={"M+":t.getMonth()+1,"d+":t.getDate(),"h+":t.getHours(),"m+":t.getMinutes(),"s+":t.getSeconds(),"q+":Math.floor((t.getMonth()+3)/3),S:t.getMilliseconds()};for(var n in/(y+)/.test(e)&&(e=e.replace(RegExp.$1,(t.getFullYear()+"").substr(4-RegExp.$1.length))),a)new RegExp("("+n+")").test(e)&&(e=e.replace(RegExp.$1,1==RegExp.$1.length?a[n]:("00"+a[n]).substr((""+a[n]).length)));return e}var h={formatDateTime:function(t,e){return e||(e="yyyy-MM-dd hh:mm"),t?f(new Date(1e3*t),e):"-"},baseFormateDate:f},g={name:"HelloWorld",data:function(){return{tableData3:[],multipleSelection:[],name:[],status:[]}},mounted:function(){this.$nextTick(function(){this.drainLine("nodesbar"),this.drainPie("jobpie"),this.getdata()})},methods:{drainLine:function(t){m.a.init(document.getElementById(t)).setOption({xAxis:{type:"category",boundaryGap:!0,data:["用户在线","用户离线"]},yAxis:{type:"value",show:!1},series:[{name:"在线",data:[23,70],type:"bar",itemStyle:{normal:{color:function(t){return["green","red"][t.dataIndex]},label:{show:!0,position:"top",formatter:"{c}"}}}}]})},drainPie:function(t){m.a.init(document.getElementById(t)).setOption({legend:{data:["Runing","Waiting","Complete"]},series:[{name:"数据",data:[{value:60,name:"Runing"},{value:20,name:"Waiting"},{value:20,name:"Complete"}],type:"pie",label:{position:"inside",formatter:"{c}%"}}]})},getdata:function(){var t=this;console.log(window.token);this.$http.get("/api/alljob/group/",{params:{token:window.sessionStorage.getItem("token"),args:p()({name:["lsfadmin"],status:[]})},headers:{"Content-Type":"application/json"}}).then(function(e){console.log(window.sessionStorage.getItem("token")),t.tableData3=JSON.parse(e.data).jobs},function(t){console.log(t)})},dateformate:function(t,e,a,n){return h.formatDateTime(a)}}},b={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-row",[a("el-col",{attrs:{span:12}},[a("div",{staticStyle:{height:"300px"},attrs:{id:"jobpie"}})]),t._v(" "),a("el-col",{attrs:{span:6}},[a("div",{staticStyle:{height:"300px"},attrs:{id:"nodesbar"}})])],1),t._v(" "),a("el-table",{ref:"multipleTable",staticStyle:{width:"100%"},attrs:{data:t.tableData3,"tooltip-effect":"dark","show-header":""}},[a("el-table-column",{attrs:{type:"selection",width:"55"}}),t._v(" "),a("el-table-column",{staticStyle:{width:"15%"},attrs:{prop:"jobid",label:"Jobid"}}),t._v(" "),a("el-table-column",{staticStyle:{width:"15%"},attrs:{prop:"submiter",label:"用户"}}),t._v(" "),a("el-table-column",{staticStyle:{width:"15%"},attrs:{prop:"qtime",formatter:t.dateformate,label:"提交时间"}}),t._v(" "),a("el-table-column",{staticStyle:{width:"15%"},attrs:{prop:"starttime",formatter:t.dateformate,label:"开始时间"}}),t._v(" "),a("el-table-column",{staticStyle:{width:"15%"},attrs:{prop:"status",label:"状态"}}),t._v(" "),a("el-table-column",{staticStyle:{width:"15%"},attrs:{prop:"endtime",formatter:t.dateformate,label:"结束时间"}})],1)],1)},staticRenderFns:[]},v=a("VU/8")(g,b,!1,null,null,null).exports,y={methods:{getdatas:function(){this.$http.get("/api/alljob/group/",{params:{token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzUyNjcxMzgsImlhdCI6MTUzNTI2MzUzOCwicGFzc3dkIjoibXl0ZXN0Iiwic3ViIjoibXl0ZXN0In0.ve0ar1FeSIT0_K_bxrqveeUIs6aN8FSpoF9NObZTyJA",args:p()({name:["lsfadmin"],status:[]})},headers:{"Content-Type":"application/json"}}).then(function(t){console.log(JSON.parse(t.data).jobs)},function(t){console.log(t)})}},data:function(){return{tableData2:[]}}},w={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("el-button",{attrs:{type:"primary"},on:{click:this.getdatas}},[this._v("登录")])],1)},staticRenderFns:[]};a("VU/8")(y,w,!1,null,null,null).exports;n.default.use(s.a);var x=new s.a({routes:[{path:"/",component:u},{path:"/main",component:v}]}),_=a("zL8q"),S=a.n(_),I=(a("tvR6"),a("8+8L"));n.default.use(I.a),n.default.use(S.a),window.gApp=new n.default({el:"#app",router:x,render:function(t){return t(r)}})},tvR6:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.5518f26de2f627db4580.js.map