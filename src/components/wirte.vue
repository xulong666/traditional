<template>
  <div class="box">
    <!-- 工具头部 -->

    <div class="header">
      <el-button round
                 type="warning"
                 @click="UploadImg"
                 icon="el-icon-picture-outline"></el-button>
      <el-button round
                 type="warning"
                 @click="Delete"
                 icon="el-icon-delete"></el-button>
      <el-button round
                 type="warning"
                 @click="Upload"
                 icon="el-icon-check"></el-button>

    </div>

    <!-- 标题 -->
    <div class="title">
      <input v-model="title"
             class="input"
             placeholder="请输入标题">
    </div>

    <!-- 文本区 -->
    <div class="content">
      <textarea class="text"
                name="writer"
                v-model="text"
                placeholder="请在此处编辑你的文章"></textarea>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      title: '',
      text: ''
    }
  },
  methods: {
    Delete () {
      this.text = ''
    },
    async Upload () {
      if (this.title === '') return this.$message.error('标题不能为空')
      if (this.text === '') return this.$message.error('文章内容不能为空')
      const { data: res } = await this.$http.post('/article/upload', {
        title: this.title, content: this.text
      })
      console.log(res.code)
      if (res.code !== 200) return this.$message.error('上传失败')
      this.$message.success('上传成功')
      return this.$router.push('/index')
    }
  }

}
</script>

<style lang="less" scoped>
.box {
  margin: 10px;
}
.header {
  height: 70px;
  background-color: rgb(243, 234, 217);
  margin-left: 0;
  display: flex;
  justify-content: flex-end;
  padding: 10px;
  button {
    margin-right: 20px;
  }
}
.title {
  margin: 20px 0;
  height: 60px;

  border-bottom: 2px solid rgb(217, 218, 211);
  .input {
    font-size: 30px;
    font-family: "黑体";
    font-weight: bolder;
    height: 55px;
    border: 0px;
    width: 100%;
    line-height: 50px;
    box-shadow: none;
    outline: none; // 不定义， 选中该区域时有边框
    margin: 0 35px;
  }
}
.content {
  width: 95%;
  .text {
    margin: 10px 35px;
    border: 0px;
    outline: none;
    font-size: 25px;
    font-family: "宋体";
    height: 600px;
    width: 100%;
    overflow:auto;
  }
}
</style>
