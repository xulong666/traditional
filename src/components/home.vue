<template>
  <el-container class="home_container">
    <!-- 头部 -->
    <el-header>
      <div>
        <img src="../assets/img/headerLogo.jpg"
             alt=""
             @click="clickImg">
        <span class="title">传统文化论坛</span>
      </div>
      <el-button type="danger"
                 round
                 @click="ClickArticle">写文章</el-button>
    </el-header>
    <!-- 页面非头部 -->
    <el-container>
      <!-- 页面主体 -->
      <el-main class="main">
        <router-view></router-view>
      </el-main>
      <!--侧边栏-->
      <el-aside>
        <!-- 排行榜 -->
        <el-card class="asideCard">
          <div class="asideBoxcss">
            <div slot="header">
              <span>排行榜</span>
              <el-button style="float: right; padding: 3px 0"
                         type="text">刷新</el-button>
            </div>
          </div>
        </el-card>
      </el-aside>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data () {
    return {
      activeIndex: 'time',
      browseSty: 'time'
    }
  },
  methods: {
    // 切换浏览模式
    handleselect (key) {
      this.browseSty = key
      console.log(this.browseSty)
    },
    clickImg () {
      this.$router.push('/home')
    },
    flasharticle () {
    },
    ClickArticle () {
      this.$router.push('/article/write')
    }
  }
}
</script>
<style lang="less" scoped>
header {
  height: 10%;
  background-image: url("../assets/img/homeHeaderBack.jpeg");
  display: flex;
  justify-content: space-between;
  > div {
    display: flex;
    align-items: center;
    img {
      width: 15%;
      height: 90%;
      border-radius: 50%;
    }
    .title {
      margin-left: 15px;
      font-family: Hiragino Sans GB;
      font-size: 20px;
    }
  }
}
.home_container {
  height: 100%;
}
aside {
  width: 25% !important;
  background-color: rgb(243, 234, 217);
}
.asideCard {
  margin: 7%;
  background-image: url("../assets/img/1629442995(1).jpg");
}
.asideBoxcss {
  height: 400px;
  opacity: 0.9;
}
.main{padding: 0 !important;}
</style>
