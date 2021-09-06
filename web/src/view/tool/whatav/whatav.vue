<template>
<div>
    <div class="head">
    <el-row>
        <el-col>
        <h1>杀毒软件进程识别</h1>
        </el-col>
    </el-row>
    </div>
    
    <el-row :gutter="20">
        <el-col :span="20" :offset="2">
            <div class="bg-purple ">
        <el-input

          type="textarea"
          :rows=8
          placeholder="输入进程名或将tasklist复制粘贴于此"
          v-model="textarea">
        </el-input>
        </div>
        </el-col>
    </el-row>
    
    <el-row :gutter="20">
        <el-col :span="20" :offset="2">
        <el-button 
            type="primary"
            style="width: 100%;" 
            @click="test">
            查询
        </el-button>
        </el-col>
    </el-row>

    <el-row>
    <el-col :span="22" :offset="1">
        <el-table 
            :data="tableData" 
            border stripe 
            :header-cell-style="{background:'#F2F2F2',color:'#606266',borderColor:'#E6E6E6','text-align':'center'}"
        >
            <el-table-column label="Antivirus" min-width="100" prop="Antivirus" />
            <el-table-column label="processes" min-width="250" prop="processes" />
            <el-table-column label="url" min-width="200" prop="url" />
        </el-table>
    </el-col>
    </el-row>

</div>
</template>
<script>
import infoList from '@/mixins/infoList'
import { whatav } from '@/api/whatav'

export default {
  mixins: [infoList],
  data() {
    return {
      textarea: '',
      listApi: whatav,
      form: {
        Antivirus: '',
        processes: '',
        url: '',
      }
    }
  },
  methods: {
      async test(){
        this.getTableData(1,999,this.textarea)
      },



  }
}
</script>

<style lang="scss" scoped>
.el-row {
    margin: 20px;
}
.el-col {
    border-radius: 4px;
}
.bg-purple {
    background: #bbbec2;
  }



.head {
    h1 {
        text-align:center;
        font-size: 300%
        }
    width: 100%;
    margin-top: 20px;}
</style>