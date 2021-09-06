<template>
  <div class="big">
    <el-row>
      <div class="card">
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                <el-col :xs="4" :md="3" :lg="3">
                  <span class="card-img">
                    <img :src="userInfo.headerImg" alt="">
                  </span>
                </el-col>
                <el-col :xs="20" :lg="12" :md="12">
                  <div class="text">
                    <h4>请开始您一天的工作吧！</h4>
                  </div>
                </el-col>
              </div>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>

    <div class="shadow">
      <el-row :gutter="20">
        <el-col
          v-for="(card, key) in toolCards"
          :key="key"
          :span="4"
          :xs="8"
          @click.native="toTarget(card.name)"
        >
          <el-card shadow="hover" class="grid-content">
            <i :class="card.icon" :style="{ color: card.color }" />
            <p>{{ card.label }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div class="bottom">
      <el-row :gutter="32">
        <el-col :xs="24" :sm="24" :lg="12">
          <div class="chart-player">
            装修中
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :lg="12">
          <div class="chart-player">
            <todo-list />
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import TodoList from './component/todoList'

export default {
  name: 'userDashboard',
  components:{
    TodoList,
  },
  data() {
    return {
      toolCards: [
        {
          label: 'cyberchef',
          icon: 'el-icon el-icon-link',
          name: 'cyberchef',
          color: '#ff9c6e'
        },
        {
          label: 'whatav',
          icon: 'el-icon el-icon-question',
          name: 'whatav',
          color: '#69c0ff'
        },
        {
          label: '菜单管理',
          icon: 'el-icon el-icon-menu',
          name: 'menu',
          color: '#b37feb'
        },
        {
          label: 'markdown',
          icon: 'el-icon el-icon-notebook-1',
          name: 'markdown',
          color: '#ffd666'
        },
        {
          label: '服务器状态',
          icon: 'el-icon el-icon-monitor',
          name: 'state',
          color: '#ff85c0'
        },
        {
          label: '个人信息',
          icon: 'el-icon el-icon-user',
          name: 'person',
          color: '#5cdbd3'
        }
      ]
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name })
    }
  }
}
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;
    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }
  .mid {
    width: 100%;
    height: 380px;
    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }
  .bottom {
    width: 100%;
    height: 300px;
    // margin: 20px 0;
    .el-row {
      margin-right: 4px !important;
    }
    .chart-player {
      width: 100%;
      height: 270px;
      padding: 10px;
      background-color: #fff;
    }
  }
}
</style>
