<template>
  <container>
    <!-- 商品查询 -->
    <el-form :inline="true" :model="query" ref="query" class="goods_query_form">
      <el-form-item prop="name">
        <el-input v-model.number="query.name" placeholder="类目名称"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" :icon="Search" @click="getCategoryList">查询</el-button>
        <el-button type="primary" @click="addLv1Category">添加一级类目</el-button>
      </el-form-item>
    </el-form>
    <!-- 商品一级类目列表 -->
    <el-table :data="categoryLv1List" height="65vh" style="width: 100%">
      <el-table-column prop="id" label="编号"/>
      <el-table-column prop="name" label="名称"/>
      <el-table-column prop="level" label="级别"/>
      <el-table-column prop="sort" label="排序"/>
      <el-table-column label="操作" min-width="130">
        <template #default="scope">
          <el-button size="small" type="primary" :icon="Edit" @click="editLv1Category(scope.row)"/>
          <el-button size="small" type="primary" :icon="Plus" @click="addLv2Category(scope.row)"/>
          <el-button size="small" type="primary" @click="checkLv2Category(scope.row)">查看二级类目</el-button>
          <el-popconfirm title="此操作将永久删除该信息, 是否继续?"
                         confirmButtonText="确认"
                         cancelButtonText="取消"
                         cancelButtonType="default"
                         :icon="WarningFilled"
                         @confirm="deleteCategory(scope.row)">
            <template #reference>
             <el-button size="small" type="danger" :icon="Delete"/>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
      <template #empty>
        <div style="margin: 50px 0;">
          <el-empty v-if="showEmpty" description="暂时还没有商品哦" />
        </div>
      </template>
    </el-table>
    <div style="padding: 10px 0;">
      <el-pagination layout="total, prev, pager, next"
                   :current-page="pageNum"
                   :page-size="pageSize"
                   :total="total"
                   @current-change="handleCurrentChange"
                   @prev-click="handleCurrentChangePrev"
                   @next-click="handleCurrentChangeNext" background/>
    </div>
    <!-- 添加、编辑商品，通用对话框 -->
    <el-dialog :title="dialogTitle" v-model="categoryDialogVisible" top="30vh" width="35%" @close="cancelCategory">
      <el-form :model="category" label-width="50px" style="padding: 10px;">
        <el-form-item label="名称">
          <el-input v-model="category.name" type="text" maxlength="10" show-word-limit/>
        </el-form-item>
        <el-form-item label="排序">
          <el-input v-model.number="category.sort" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelCategory">取消</el-button>
          <el-button type="primary" @click="confirmCategory">确定</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 查看二级类目列表，对话框 -->
    <el-dialog :title="dialogTitle" v-model="checkLv2CategoryDialogVisible" top="10vh" width="50%">
      <el-table :data="categoryLv2List" height="60vh" style="width: 100%">
        <el-table-column prop="id" label="编号"/>
        <el-table-column prop="name" label="名称"/>
        <el-table-column prop="level" label="级别"/>
        <el-table-column prop="sort" label="排序"/>
        <el-table-column label="操作" min-width="130">
          <template #default="scope">
            <el-button size="small" type="primary" :icon="Edit" @click="editLv2Category(scope.row)"/>
            <el-popconfirm title="此操作将永久删除该信息, 是否继续?"
                         confirmButtonText="确认"
                         cancelButtonText="取消"
                         cancelButtonType="default"
                         :icon="WarningFilled"
                         @confirm="deleteCategory(scope.row)">
            <template #reference>
             <el-button size="small" type="danger" :icon="Delete"/>
            </template>
          </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </container>
</template>

<script>
import container from "../components/container";
import {Delete, Edit, Plus, Search, WarningFilled} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";

export default {
  name: "Category",
  components: {container},
  setup() {
    return {Search, Plus, Edit, Delete, WarningFilled}
  },
  data() {
    return {
      categoryLv1List: [],
      categoryLv2List: [],
      category: {
        id: '',
        name: '',
        parentId: '',
        level: '',
        sort: ''
      },
      query: {
        name: ''
      },
      dialogTitle: '',
      operateType: '',
      categoryDialogVisible: false,
      checkLv2CategoryDialogVisible: false,

      // 分页
      total: 0,
      pageNum: 1,
      pageSize: 10,

      // 空状态
      showEmpty: false
    }
  },
  mounted() {
    this.getCategoryList(1)
  },
  methods: {

    // 分页，处理函数
    handleCurrentChangePrev(val) {
      this.pageNum = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.pageNum = val;
      this.getCategoryList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.pageNum = val;
      console.log(`下一页: ${val}`);
    },

    // 添加一级类目
    addLv1Category(){
      this.operateType = 'add'
      this.dialogTitle = '添加一级类目'
      this.category.level = 1
      this.category.parentId = 1
      this.categoryDialogVisible = true
    },

    // 编辑一级类目
    editLv1Category(row){
      this.operateType = 'edit'
      this.dialogTitle = '编辑一级类目'
      this.category.id = row.id
      this.category.name = row.name
      this.category.sort = row.sort
      this.categoryDialogVisible = true
    },

    // 添加二级类目
    addLv2Category(row){
      this.operateType = 'add'
      this.dialogTitle = '添加二级类目'
      this.category.level = 2
      this.category.parentId = row.id
      this.categoryDialogVisible = true
    },

    // 编辑二级类目
    editLv2Category(row){
      this.operateType = 'edit'
      this.dialogTitle = '编辑二级类目'
      this.category.id = row.id
      this.category.name = row.name
      this.category.sort = row.sort
      this.categoryDialogVisible = true
    },

    // 查看二级类目
    checkLv2Category(row){
      this.dialogTitle = '查看二级类目'
      this.getCategoryList(row.id)
      this.checkLv2CategoryDialogVisible = true
    },

    // 获取类目列表
    getCategoryList(parentId){
      this.$axios.get('/category/list', {
        params: {
          name: this.query.name,
          parentId: parentId,
          sid: localStorage.getItem('sid')
        }
      }).then((response) => {
        if (parentId === 1){
          this.categoryLv1List = response.data.data;
        } else {
          this.categoryLv2List = response.data.data;
        }
        if (this.categoryLv1List.length === 0) {
          this.showEmpty = true
        }
      }).catch((error) => {
        console.log(error)
      })
    },

    // 确定添加、编辑类目
    confirmCategory(){
      if (this.operateType === 'add'){
        this.$axios.post('/category/create', {
          name: this.category.name,
          parentId: this.category.parentId,
          level: this.category.level,
          sort: this.category.sort,
          sid: localStorage.getItem('sid')
        }).then((response) => {
          if (response.data.code === 200) {
            ElMessage({message: response.data.message, type: 'success'})
          }
          this.getCategoryList(1)
        }).catch((error) => {
          console.log(error)
        })
      }
      if (this.operateType === 'edit'){
        this.$axios.put('/category/update', {
          id: this.category.id,
          name: this.category.name,
          sort: this.category.sort
        }).then((response) => {
          if (response.data.code === 200) {
            ElMessage({message: response.data.message, type: 'success'})
          }
          this.getCategoryList(1)
        }).catch((error) => {
          console.log(error)
        })
      }
      this.categoryDialogVisible = false
      this.checkLv2CategoryDialogVisible = false
    },
    deleteCategory(row){
      this.$axios.delete('/category/delete', {
          params: {
            id: row.id
          }
        }).then((response) => {
          if (response.data.code === 200) {
            ElMessage({message: response.data.message, type: 'success'})
          }
          this.getCategoryList(1)
        }).catch((error) => {
          console.log(error)
      })
    },
    cancelCategory(){
      this.categoryDialogVisible = false
      this.emptyCategory()
      this.dialogTitle = ''
      this.operateType = ''
    },

    // 清空类目
    emptyCategory(){
      this.category.id = ''
      this.category.name = ''
      this.category.parentId = ''
      this.category.level = ''
      this.category.sort = ''
    }
  }
}
</script>

<style scoped>
.el-dialog {
  height: 200px !important;
  overflow-y: scroll;
}
</style>