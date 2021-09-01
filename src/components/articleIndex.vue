<template>
  <div>
    <!-- 标签页 -->
    <div>
      <!-- 菜单栏 其中mode可以设置水平放置 -->
      <el-menu :default-active="activeIndex"
               mode="horizontal"
               @select="handleselct(items.ID)">
        <el-menu-item index="time">时间</el-menu-item>
        <el-menu-item index="num">浏览量</el-menu-item>
      </el-menu>
    </div>
    <!-- 显示主体内容 -->
    <div class="allcard">
      <el-card class="articleCardcss"
               v-for="items in articles"
               :key="items.ID"
               slot="header">
        <!-- 使用插槽加入文章名-->
        <div slot="header"
             class="titleCss"
             @click="OpenArticle(items.ID, items.Title)">
          <span>{{items.Title}}</span>
        </div>
        <div class="abstractContentcss">
          {{items.Abstract}}......
        </div>
      </el-card>
    </div>
    <!-- 设置分页-->
    <el-pagination background
                   layout="prev, pager, next"
                   :total="total"
                   page-size=6
                   current-page="index"
                   @current-change="clickPageIdex"
                   @prev-click="clickPrev"
                   @next-click="clickNext"
                   class="pagination">
    </el-pagination>
  </div>
</template>

<script>
export default {
  data () {
    return {
      activeIndex: 'time',
      browseSty: 'time',
      articles: [{ ID: 1, Abstract: 'hello', Title: 'world' }],
      total: 1,
      index: 0
    }
  },
  // 创建页面时就请求数据
  created () {
    this.getindex(0)
  },
  methods: {
    // 切换浏览模式
    handleselct (key) {
      this.browseSty = key
      console.log(this.browseSty)
    },
    // 请求后台数据
    async getindex (index) {
      const { data: res } = await this.$http.get('/article/index', { params: { id: index - 1 } })
      this.articles = res.articles
      console.log(res)
      this.total = res.allnum
      console.log(res)
    },

    // 点击页码
    clickPageIdex (index) {
      this.index = index
      this.getindex(this.index)
    },

    // 点击前一页
    clickPrev () {
      this.index -= 1
      this.getindex(this.index)
    },

    // 点击后一页
    clickNext () {
      this.index += 1
      this.getindex(this.index)
    },
    // 获取文章（点击标题）
    OpenArticle (id, title) {
      this.$router.push({ name: 'article', query: { id: id, title: title } })
    }

  }
}
</script>

<style lang="less" scoped>
// .allcard {
//   height: 600px;
// }
.articleCardcss {
  margin: 30px;
  height: 150px;
  border-radius: 10px;
  box-shadow: 1px, 1px, 3 !important;
}
.titleCss {
  font-weight: bold;
  line-height: 90%;
  font-family: "PingFang SC";
  font-size: 100%;
  display: flex;
  justify-content: center;
}
.pagination {
  position: relative;
  bottom: 10%;
  left: 5%;
}
</style>
