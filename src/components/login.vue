<template>
  <div class="login_container">
    <!--登录区域-->
    <div class="login_box">
      <!--头像区域-->
      <div class="header_pic">
        <img src="../assets/logo.png"
             alt="">
      </div>
      <!--登录表单-->

      <el-form ref="loginFormRef"
               :model="loginForm"
               :rules="rules"
               label-width="100px"
               class="login_form">
        <el-form-item label="账号"
                      prop="name">
          <el-input v-model="loginForm.name"
                    prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item label="密码"
                      prop="password">
          <el-input type="password"
                    v-model="loginForm.password"
                    prefix-icon="el-icon-lock"></el-input>
        </el-form-item>
        <!--登录和注册按钮-->
        <div class="button">
          <el-button type="primary"
                     @click="login">登录</el-button>
          <el-button>注册</el-button>
        </div>
      </el-form>

    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      loginForm: {
        name: '张三',
        password: '123'
      },

      // 定义校验规则
      rules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' },
          { min: 3, message: '长度需大于3个字符', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' },
          { min: 5, max: 15, message: '长度在5 到 15 之间', trigger: 'blur' }]
      }
    }
  },
  methods: {
    login () {
      console.log(this.$refs.loginFormRef)
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return
        // 这个{data：res}需要设置为const
        const { data: res } = await this.$http.post('/user/login', {
          username: this.loginForm.name,
          password: this.loginForm.password
        })
        console.log(res)
        if (res.code !== 200) return this.$message.error('登录失败！')
        this.$message.success('登录成功')
        console.log(res.token)
        // 登录成功后存储tokrn， 转向home
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('/home')
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login_container {
  background-color: rgb(211, 211, 187);
  height: 100%;
}
.login_box {
  background-color: rgb(206, 204, 84);
  height: 40%;
  width: 30%;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  border-radius: 10px;
  box-shadow: 0.5px 0.5px 3px #2e2d27;
  .header_pic {
    background-color: #69633e;
    height: 100px;
    width: 100px;
    border-radius: 50%;
    padding: 10px;
    position: absolute;
    left: 50%;
    transform: translate(-50%, -50%);
    box-shadow: 1px, 1px, 3px, #69633e;
    border: solid #2e2d27;
    img {
      height: 100%;
      width: 100%;
      border-radius: 50%;
      background-color: #eee;
    }
  }
  .login_form {
    position: absolute;
    bottom: 5%;
    width: 100%;
    padding: 20px;
    right: 7%;
    box-sizing: border-box;
    .button {
      display: flex;
      justify-content: flex-end;
    }
  }
}
</style>
