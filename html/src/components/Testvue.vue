<template>
<div>
<el-button type="primary" @click="getdatas">登录</el-button>
</div>
</template>

<script>
  export default {
    methods: {
      getdatas(){
        var mdata = {
            name: ["lsfadmin"],
            status: []
        }
        this.$http.get(
            "/api/alljob/group/",
             {params: {token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzUyNjcxMzgsImlhdCI6MTUzNTI2MzUzOCwicGFzc3dkIjoibXl0ZXN0Iiwic3ViIjoibXl0ZXN0In0.ve0ar1FeSIT0_K_bxrqveeUIs6aN8FSpoF9NObZTyJA',
             args:JSON.stringify(mdata)}, headers:{
                 "Content-Type": "application/json"
             }}).then(response =>{
               
                    console.log(JSON.parse(response.data)["jobs"])
             },response => {
                  console.log(response)
             });
    }
    },
    data() {
      return {
        tableData2: []
      }
    },
    
  }
</script>