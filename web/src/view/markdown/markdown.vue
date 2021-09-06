<template>
<div>
  <el-row>
    <el-col :span=3>

      <div class="button-box clearflex">
      <el-button-group>
        <el-button @click="addNode" icon="el-icon-plus" type="primary" size="small">添加</el-button>
        <el-button @click="deleteNode" icon="el-icon-remove" type="primary" size="small">删除</el-button>
      </el-button-group>
      </div>

      <div class="filter-input">
        <el-input placeholder="" v-model="filterText" size="small" prefix-icon="el-icon-search">
           </el-input>
      </div>


      <el-tree
        ref="tree"
        :data="treeData"
        :props="defaultProps"
        node-key="uuid"
        highlight-current
        :filter-node-method="filterNode"
        :expand-on-click-node="false"
        :default-expand-all="true"
        @node-contextmenu="cancelSelect"
        @node-click="handleNodeClick">
        <span class="custom-tree-node" slot-scope="{ node }">
          <span v-show="!node.isEdit">
            {{ node.label }}
          </span>
          <span class="edit-input" v-show="node.isEdit">
        			<el-input size="mini"
                v-model="markdowninfo.title"
                :ref="'inputRef' + node.data.uuid"
                @blur="updatenode(node)"
                @change="updatenode(node)">
              </el-input>
         </span>
          <span v-show="node.store.currentNode && node.key === node.store.currentNode.key">
              <el-button
                icon='el-icon-edit-outline'
                size='mini'
                type='text'
                @click="editnode(node)">
              </el-button> 
          </span>

        </span>
      </el-tree>
    </el-col>

    <el-col :span=21>
      <mavon-editor
            ref="md"
            placeholder=""
            :boxShadow="false"
            style="z-index:1;border: 1px solid #d9d9d9;height:85vh"
            v-model="mkvalue"
            @save="save"
            @imgAdd="handleEditorImgAdd"
        />
    </el-col>
  </el-row>
</div>
</template>
<script>
import{savemd,getmd,deletemd,getmdtree} from '@/api/markdown'
import{uploadFile,deleteFile} from '@/api/fileUploadAndDownload'
import { mapGetters} from 'vuex'
import { uuid } from 'vue-uuid'
let Base64 = require('js-base64').Base64


  export default {
    name: 'Markdown',
    computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
    data() {
      return {
        treeData: [],
        defaultProps: {
          children: "children",
          label: "title",
        },
        markdowninfo:{
          uuid : '',
          editor : '',
          title : '',
          parentId : '',
        },
        tempNodeLabel:'',
        mkvalue : this.mkvalue,
        filterText: '',
      }
    },
    async created() {
      this.getNodeTree()
    },
    watch: {
      filterText(val) {
        this.$refs.tree.filter(val);
      }
    },
    methods: {
      cancelSelect(){
        this.$refs.tree.setCurrentKey(null)
      },

      filterNode(value, data) {
        if (!value) return true;
        return data.title.indexOf(value) !== -1;
      },

      addNode() {
        if (this.$refs.tree.getCurrentKey()){
          let node = this.$refs.tree.getNode(this.$refs.tree.getCurrentKey())
          const data = {           
            uuid: uuid.v4(),
            editor : this.userInfo.userName,
            title : "新建md",
            mkvalue : "",
            parentId: node.data.ID.toString()
            }
          savemd(data).then((res)=> {
            if (res.code === 0) {
              data.ID = res.data.markdown.ID
              this.$message.success('保存成功')
            }
          })
          this.$refs.tree.append(data,this.$refs.tree.getCurrentKey())
        }
        else{
          console.log("addroot")
          this.addRootNode()
        }
      },

      addRootNode(){
        const data = {           
           uuid: uuid.v4(),
           editor : this.userInfo.userName,
           title : "新建分类",
           mkvalue : "",
           parentId: 0}
        savemd(data).then((res)=> {
           if (res.code === 0) {
             data.ID = res.data.markdown.ID
             this.$message.success('保存成功')
            }
         })
        
        this.treeData.push(data)
      },
      


      updatenode(node){
        //console.log(node.data.uuid)
        this.$set(node, 'isEdit', false)

        this.$nextTick(() => {
          this.$set(node.data,'title',this.markdowninfo.title)
          this.save()
      })

      },

      editnode(node){
        if (!node.isEdit) { // 检测isEdit是否存在or是否为false
          this.$set(node, 'isEdit', true)
          this.$nextTick(() => {
            this.$refs["inputRef"+node.data.uuid].$el.children[0].focus();
          })
        }
      },
      


      handleEditorImgAdd(pos, $file){
        let formdata = new FormData()
        formdata.append('file', $file)
        uploadFile(formdata).then((res)=> {
          let url = 'https://'+res.data.file.url
          this.$refs.md.$img2Url(pos, url)
         })

      },

      async deleteNode() {
        if(this.$refs.tree.getCurrentKey()){
          let nodekey = this.$refs.tree.getCurrentKey()
          let node = this.$refs.tree.getNode(nodekey)
          if(node.childNodes && node.childNodes.length){
              this.$message.warning('含有子节点，请删除子节点后再删除')
            }
          else{
            const res = await deletemd({UUID: node.data.uuid})
            if (res.code === 0){
              this.$message.success('删除成功')
            }
            this.$nextTick(() => {
              this.$refs.tree.remove(nodekey)
              this.cancelSelect()
            })
          }
        }
        else{
          this.$message.warning('请选择想删除的节点')
        }
        
        
      },
      async save(){ 
        savemd({
           uuid: this.markdowninfo.uuid,
           editor :this.userInfo.userName,
           title : this.markdowninfo.title,
           mkvalue : Base64.encode(this.mkvalue),
           parentId: this.markdowninfo.parentId
         }).then((res)=> {
           if (res.code === 0) {
             this.$message.success('保存成功')
            }
         })
        
      },
      


      async handleNodeClick(data) {

        const res = await getmd({UUID: data.uuid})
        if (res.code === 0) {
          this.mkvalue = Base64.decode(res.data.markdown.mkvalue)
          this.markdowninfo.uuid = res.data.markdown.uuid
          this.markdowninfo.editor = res.data.markdown.editor
          this.markdowninfo.title = res.data.markdown.title
          this.markdowninfo.parentId = res.data.markdown.parentId
        }
      },

      async getNodeTree(){
        const res = await getmdtree()
        this.treeData = res.data.markdown
      },

    },
  }
</script>

<style lang="scss" >
.button-box {
  padding: 10px 20px;
  .el-button-group {
    float: left;
  }
}


.el-input__inner,.el-input-group__prepend{
       width: calc(100% - 20px);
       margin:0 10px;
       height: 40px;
       border-top:none;
       border-width: 0 0 1px;
       border-radius:0;
     }
.el-input__prefix{
    .el-input__icon{
      margin-right: 15px;
      display: inline-block;
    }
    font-size:18px;
  }
.el-input:focus-within{
        .el-icon-search:before {
           color: #3c6eff;
           font-weight: bold;
        }
      }

.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
    .edit-input{
      .el-input,.el-input--mini{
         .el-input__inner{
           height: 30px;
           border-top:none;
         } 
      }
    }

}



</style>