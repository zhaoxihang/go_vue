<template>
  <div>
    <v-app-bar mobileBreakpoint="sm" app dark flat color="white darken-5">
      <v-app-bar-nav-icon dark class="hidden-md-and-up" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title class='hidden-md-and-down'>
        <v-app-bar-nav-icon class="mx-15">
          <v-avatar size="40" color="grey">
            <img :src="deploy.logo" alt />
          </v-avatar>
        </v-app-bar-nav-icon>

        <strong class="green--text text--lighten-1" v-html="deploy.abstract" @mouseenter="enter" @mouseleave="leave"></strong>
        <span v-show="details_show" class="blue--text text--lighten-2">
          &nbsp;&nbsp;({{ deploy.details }})
        </span>

      </v-toolbar-title>

<!--      <v-tabs dark center-active centered class="hidden-sm-and-down">-->
<!--        <v-tab @click="$router.push('/')">首页</v-tab>-->
<!--        <v-tab-->
<!--          v-for="item in cateList"-->
<!--          :key="item.id"-->
<!--          text-->
<!--          @click="gotoCate(item.id)"-->
<!--        >{{ item.name }}</v-tab>-->
<!--      </v-tabs>-->

      <v-spacer></v-spacer>
      <v-hover>
        <template v-slot:default="{ hover }">
          <v-text-field
            dense
            hide-details
            solo-inverted
            rounded
            label="请输入文章标题查找"
            append-icon="mdi-text-search"
            v-model="searchName"
            @change="searchTitle(searchName)"
            class='blue lighten-2'
            :class="`elevation-${hover ? 50 : 6}`"
            style='width: 1%'
          ></v-text-field>
        </template>
      </v-hover>

    </v-app-bar>

<!--    <v-navigation-drawer v-model="drawer" color="indigo" dark app temporary>-->
<!--      <v-list>-->
<!--        <v-list-item-title>-->
<!--          <v-btn href="/" dark text>-->
<!--            <v-icon small>mdi-home</v-icon>首页-->
<!--          </v-btn>-->
<!--        </v-list-item-title>-->

<!--        <v-list-item-->
<!--          v-model="group"-->
<!--          active-class="deep-purple&#45;&#45;text text&#45;&#45;accent-4"-->
<!--          v-for="item in cateList"-->
<!--          :key="item.id"-->
<!--        >-->
<!--          <v-list-item-title>-->
<!--            <v-btn dark text @click="gotoCate(item.id)">{{ item.name }}</v-btn>-->
<!--          </v-list-item-title>-->
<!--        </v-list-item>-->
<!--      </v-list>-->
<!--    </v-navigation-drawer>-->
  </div>
</template>

<script>
export default {
  data() {
    return {
      deploy:{
        logo:false,
        abstract:false,
        details:false,
      },
      details_show:false,
      drawer: false,
      group: null,
      valid: true,
      registerformvalid: true,
      cateList: [],
      searchName: '',
      formdata: {
        username: '',
        password: ''
      },
      checkPassword: '',
      dialog: false,
      headers: {
        Authorization: '',
        username: ''
      },
    }
  },
  watch: {
    group() {
      this.drawer = false
    }
  },
  created() {
    this.GetCateList()
    this.getBlogHeadPortrait()
  },
  mounted() {
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`,
      username: window.sessionStorage.getItem('username')
    }
  },
  methods: {
    // 获取分类
    async GetCateList() {
      const { data: res } = await this.$http.get('category')
      this.cateList = res.data
    },

    // 查找文章标题
    searchTitle(title) {
      if (title.length == 0) return this.$message.error('你还没填入搜索内容哦')
      this.$router.push(`/search/${title}`)
    },

    gotoCate(cid) {
      console.log(cid,"cid")
      this.$router.push(`/category/${cid}`).catch((err) => err)
    },

    async getBlogHeadPortrait(){
      // 获取 博客设置
      const { data: res } = await this.$http.get(
        `deploy`
      )
      this.deploy = res.data
    },
    enter(){
      this.details_show = true;
    },
    leave(){
      this.details_show = false;
    }
  }
}
</script>

<style></style>
