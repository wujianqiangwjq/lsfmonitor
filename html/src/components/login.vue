<template>
<div class="login-back">
    <div class="login-container">
      <el-form :model='user'  ref="user">
        <el-form-item style="text-align:center">
          <h1>用户登录</h1>
        </el-form-item>
          <el-form-item prop="username" style="text-align:center">
              <el-input type="text" autofocus="true" v-model="user.username" auto-complete="off" placeholder="用户名"></el-input>
          </el-form-item>
          <el-form-item >
              <el-input type="password" autofocus="true" v-model="user.password" auto-complete="off"  placeholder="密码"></el-input>
          </el-form-item>
          <el-form-item  style="text-align:center">
              <el-button type="primary" style="width: 150px; margin-top: 10px;"  @click="logind">登录</el-button>
          </el-form-item>
      </el-form>
    </div>
</div>

</template>
<script>

export default {
  data(){
    	return {
	      user: { username: '', password: ''},
	    }
  },
  methods:{
        logind(){      
        	this.$http.post('/api/login',
        		this.user,
        		{
        			'Content-Type':'application/json'
        		}
          	).then((resp) => {
              console.log("go success");
              if (resp.status == 200){
                 window.sessionStorage.setItem("token",resp.data.token);
                 this.$router.push({ path: '/main' });
              }
          	},(err) => {
               this.user.password = "";
               this.$message.error(err.data.error);
            });
           
      }
  }

}
</script>
<style>
.login-container{
    position: absolute;
    transform: translate(-50%,-50%);
    top:50%;
    left:50%;
    min-width:300px;
    max-width: 400px;
    padding: 30px 30px 20px;
    background: rgb(29, 110, 150);
    border-radius: 4px;
    
  }
  .login-back{
   	position: absolute;
    background-image:url("../assets/image/backw.jpg");
    background-size: 100% 100%;
    -moz-background-size: 100% 100%;
   -webkit-background-size: 100% 100%;
    width:100%;
    height:100%;
    opacity: 0.9;
  }
</style>


