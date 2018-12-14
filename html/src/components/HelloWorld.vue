<template>
<div>
	<el-row>
	  <el-col :span="12"><div id="jobpie" style="height: 300px;" ></div></el-col>
	  <el-col :span="6"><div id="nodesbar" style="height: 300px;" ></div></el-col>
	</el-row>
	<el-table ref="multipleTable" :data="tableData3" tooltip-effect="dark" style="width: 100%" show-header>
    <el-table-column type="selection" width="55"> </el-table-column>
    <el-table-column prop="jobid" label="Jobid" style="width:15%"></el-table-column>
    <el-table-column prop="submiter" label="用户" style="width:15%"></el-table-column>
    <el-table-column prop="qtime" :formatter="dateformate" label="提交时间" style="width:15%"></el-table-column>
    <el-table-column prop="starttime" :formatter="dateformate" label="开始时间" style="width:15%"></el-table-column>
    <el-table-column prop="status" label="状态" style="width:15%"></el-table-column>
    <el-table-column prop="endtime" :formatter="dateformate" label="结束时间" style="width:15%"></el-table-column>
  </el-table>
</div>
</template>

<script>
import echarts from 'echarts'
import formate  from '../common/formate'
export default {
  name: 'HelloWorld',
  
data() {
      return {
        tableData3: [],
        multipleSelection: [],
        name:[],
        status:[]
      }
},
mounted(){
	this.$nextTick(function() {
      this.drainLine("nodesbar");
      this.drainPie("jobpie");
      this.getdata();
  })
  	
  },
methods:{
    drainLine(id){
        var charts = echarts.init(document.getElementById(id));
        var options = {
            xAxis: {
                type: 'category',
                boundaryGap: true,
                data: ["用户在线","用户离线"]
            },
            yAxis: {
            type: 'value', 
            show: false
            },
            series: [{
              name: "在线",
                data: [23,70],
            type: 'bar',
            itemStyle:{
              normal:{
              color:function(params){
                var colorlist= ['green','red'];
                return colorlist[params.dataIndex]
              },
              label:{
                show: true,
                position:'top',
                formatter:'{c}'
              }
              }
            }
              
            }]
            
      };
      charts.setOption(options);
    },
    drainPie(id){
      var piecharts = echarts.init(document.getElementById(id));
        var options = {
            legend:{
              data:["Runing",'Waiting','Complete']
            },
            series: [{
                name: "数据",
                data: [
                {
                  value:60, name:'Runing'
                }
                ,{
                  value:20, name:'Waiting'
                },{
                  value:20, name:'Complete'
                }],
                type:'pie',
                label:{
              position: 'inside',
              formatter: "{c}%"
                }
                
            }]
      };
      piecharts.setOption(options);
    },
    getdata(){
      console.log(window.token);
      var mdata = {
            name: ["lsfadmin"],
            status: []
        }
        this.$http.get(
            "/api/alljob/group/",
             {params: {token:window.sessionStorage.getItem("token"),
             args:JSON.stringify(mdata)}, headers:{
                 "Content-Type": "application/json"
             }}).then(response =>{   
                console.log(window.sessionStorage.getItem("token"));
                this.tableData3 = JSON.parse(response.data)["jobs"];

             },response => {
                  console.log(response)
             });
    },
    dateformate(row, column, cellValue, index){
      return formate.formatDateTime(cellValue);
    }
}

    
}
</script>
