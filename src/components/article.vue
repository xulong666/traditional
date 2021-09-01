<template>
  <div>
    <!-- 标题 -->
    <div class="header">
      <span>{{title}}</span>
    </div>
    <!-- 内容 -->
    <div class="content">
      {{content}}
    </div>
    <!-- 评论区 -->
    <div class="pinglun">
      pinglunqu
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      id: 0,
      title: 'sorry',
      content: 'sorry'
    }
  },
  created () {
    this.getcontent()
  },
  methods: {
    async getcontent () {
      this.id = this.$route.query.id
      this.title = this.$route.query.title
      const { data: res } = await this.$http.get('/article/' + this.id.toString())
      this.content = res.content
      console.log(this.content)
    }
  }
}
</script>

<style lang="less" scoped>
.header {
  height: 50px;
  line-height: 50px;
  text-align: center;
  font-size: 30px;
  border-bottom: 2px solid rgb(226, 217, 217);
}
.content {
  margin-top: 20px;
  line-height: 35px;
  font-size: 23px;
  font-family: '宋体';
  margin-left: 15px;
  margin-right: 15px;
}
.pinglun{
  margin-top:70px;
}
</style>
