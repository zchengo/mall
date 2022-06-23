<template>
  <Container>
    <!-- 商品查询 -->
    <el-form :inline="true" :model="query" ref="query" class="goods_query_form">
      <el-form-item prop="id">
        <el-input v-model.number="query.id" placeholder="商品ID"/>
      </el-form-item>
      <el-form-item prop="title">
        <el-input v-model="query.title" placeholder="商品标题"/>
      </el-form-item>
      <el-form-item prop="categoryId">
        <el-cascader v-model="query.categoryId"
                     :options="categoryOption" @focus="getCategoryOption" placeholder="请选择" clearable/>
      </el-form-item>
      <el-form-item prop="status">
        <el-select v-model="query.status" placeholder="商品状态">
          <el-option label="已上架" value="1"></el-option>
          <el-option label="未上架" value="2"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" :icon="Search" @click="getGoodsList">查询</el-button>
        <el-button type="primary" :icon="Brush" @click="resetForm"/>
        <el-button type="primary" :icon="Plus" @click="add">商品</el-button>
      </el-form-item>
    </el-form>
    <!-- 商品列表 -->
    <el-table :data="goodsList" height="70vh" style="width: 100%">
      <el-table-column prop="name" label="商品名称" min-width="240">
        <template #default="scope">
          <div style="display: inline-flex;">
            <div class="goods_image_box">
              <el-image class="goods_image" :src="scope.row.imageUrl"/>
            </div>
            <div style="display: block;">
              <div class="goods_title">{{ scope.row.title }}</div>
              <div class="goods_id">ID：{{ scope.row.id }}</div>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="price" label="价格" sortable/>
      <el-table-column prop="quantity" label="库存"/>
      <el-table-column prop="sales" label="销量" sortable/>
      <el-table-column prop="status" label="状态">
        <template #default="scope">
          <el-tag v-if="scope.row.status === 1" size="small">出售中</el-tag>
          <el-tag v-if="scope.row.status === 2" size="small" type="success">仓库中</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created" label="创建时间" min-width="150" sortable>
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-icon>
              <timer/>
            </el-icon>
            <span style="margin-left: 10px">{{ scope.row.created }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="130">
        <template #default="scope">
          <el-button size="small" :icon="Edit" @click="edit(scope.$index, scope.row)"/>
          <el-button v-if="scope.row.status === 2" size="small" type="primary" @click="changeStatus(1,scope.row)">上架
          </el-button>
          <el-button v-if="scope.row.status === 1" size="small" type="primary" @click="changeStatus(2,scope.row)">下架
          </el-button>
          <el-popconfirm title="此操作将永久删除该信息, 是否继续?"
                         confirmButtonText="确认"
                         cancelButtonText="取消"
                         cancelButtonType="default"
                         :icon="WarningFilled"
                         @confirm="deleteGoods(scope.row)">
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
    <el-dialog :title="dialogTitle" v-model="goodsDialogVisible" :lock-scroll="false" top="5vh" width="45%" @close="cancel">
      <el-form :model="goods" label-width="80px">
        <el-form-item label="类目" :required="true">
          <el-cascader v-model="goods.categoryId"
                       :options="categoryOption"
                       @change="changeCategory"
                       change-on-select
                       @focus="getCategoryOption" placeholder="请选择" clearable/>
        </el-form-item>
        <el-form-item label="标题" :required="true">
          <el-input v-model="goods.title" style="width: 90%;" type="text" maxlength="30" show-word-limit/>
        </el-form-item>
        <el-form-item label="名称" :required="true">
          <el-input v-model="goods.name" style="width: 50%;" type="text" maxlength="10" show-word-limit/>
        </el-form-item>
        <el-form-item label="价格" :required="true">
          <el-input v-model.number="goods.price" style="width: 50%;">
            <template #append>元</template>
          </el-input>
        </el-form-item>
        <el-form-item label="数量" :required="true">
          <el-input v-model.number="goods.quantity" style="width: 50%;">
            <template #append>件</template>
          </el-input>
        </el-form-item>
        <el-form-item label="图片" :required="true">
          <el-upload
              action="http://localhost:8000/web/upload"
              :headers="{'token': token}"
              :limit="1"
              name="file"
              :file-list="pictureList"
              list-type="picture-card"
              :on-preview="handleImagePreview"
              :on-remove="handleImageRemove"
              :on-success="handleImageSuccess">
            <div class="goods_image_upload_icon">+</div>
          </el-upload>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="goods.remark"
                    style="width: 90%;"
                    type="textarea"
                    maxlength="100"
                    :autosize="{ minRows: 2}" show-word-limit/>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
          <el-button @click="empty">重置</el-button>
          <el-button type="primary" @click="confirmGoods">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </Container>
</template>

<script>
import {ArrowDown, ArrowUp, Brush, Delete, Edit, Plus, Search, Timer, WarningFilled} from "@element-plus/icons-vue";
import {ElMessage} from 'element-plus'
import Container from "../components/Container";

export default {
  name: "Goods",
  components: {Container, Timer,},
  setup() {
    return {Search, Brush, Plus, Edit, ArrowUp, ArrowDown, Delete, WarningFilled}
  },
  data() {
    return {

      // 查询条件
      query: {
        id: '',
        title: '',
        categoryId: '',
        status: ''
      },

      // 类目选项
      categoryOption: [],

      // 商品
      goodsList: [],
      goods: {
        id: '',
        categoryId: '',
        title: '',
        name: '',
        brandId: '',
        price: '',
        quantity: '',
        imageUrl: '',
        remark: '',
        status: ''
      },

      // 图片列表
      pictureList: [],

      // 分页
      total: 0,
      pageNum: 1,
      pageSize: 10,

      // 请求认证
      token: '',

      // 操作状态
      dialogTitle: '',
      operateType: '',
      goodsDialogVisible: false,

      // 空状态
      showEmpty: false
    }
  },
  mounted() {
    this.getGoodsList()
    this.token = localStorage.getItem('token')
  },
  methods: {

    // 分页，处理函数
    handleCurrentChangePrev(val) {
      this.pageNum = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.pageNum = val;
      this.getGoodsList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.pageNum = val;
      console.log(`下一页: ${val}`);
    },

    // 图片上传，处理函数
    handleImagePreview(file) {
      this.goods.imageUrl = file.url
    },
    handleImageRemove(file, fileList) {
      this.goods.imageUrl = ''
      console.log(file, fileList)
    },
    handleImageSuccess(response) {
      this.goods.imageUrl = response.data
    },

    // 添加、编辑商品
    add() {
      this.empty()
      this.dialogTitle = '添加商品'
      this.operateType = 'add'
      this.goodsDialogVisible = true
    },
    edit(index, row) {
      this.dialogTitle = '编辑商品'
      this.operateType = 'edit'
      this.goods.id = row.id
      this.goods.categoryId = row.categoryId
      this.goods.title = row.title
      this.goods.name = row.name
      this.goods.price = row.price
      this.goods.quantity = row.quantity
      this.goods.imageUrl = row.imageUrl
      this.pictureList.push({url: row.imageUrl})
      this.goods.remark = row.remark
      this.goods.status = row.status
      this.goodsDialogVisible = true
    },

    changeCategory(value) {
      this.goods.categoryId = value[1]
    },

    // 获取商品类目选项
    getCategoryOption() {
      this.$axios.get('/category/option').then((response) => {
        this.categoryOption = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },

    // 获取商品列表
    getGoodsList() {
      this.$axios.get('/goods/list', {
        params: {
          id: this.query.id,
          title: this.query.title,
          categoryId: this.query.categoryId[1],
          status: this.query.status,
          pageNum: this.pageNum,
          pageSize: this.pageSize
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.goodsList = response.data.data.list;
        if (this.goodsList.length === 0) {
          this.showEmpty = true
        }
      }).catch((error) => {
        console.log(error)
      })
    },

    // 商品上架、下架
    changeStatus(status, row) {
      this.$axios.put('/goods/status/update', {
        id: row.id,
        status: parseInt(status)
      }).then((response) => {
        if (response.data.code === 200) {
          ElMessage({message: response.data.message, type: 'success'})
        }
        this.getGoodsList()
      }).catch((error) => {
        console.log(error)
      })
      this.getGoodsList()
    },

    // 确认添加或编辑商品
    confirmGoods() {
      let valid = this.goodsFormValidator()
      if (valid) {
        if (this.operateType === 'add') {
          // 添加商品
          this.$axios.post('/goods/create', {
            categoryId: parseInt(this.goods.categoryId),
            title: this.goods.title,
            name: this.goods.name,
            price: parseInt(this.goods.price),
            quantity: parseInt(this.goods.quantity),
            imageUrl: this.goods.imageUrl,
            remark: this.goods.remark
          }).then((response) => {
            if (response.data.code === 200) {
              ElMessage({message: response.data.message, type: 'success'})
            }
            this.getGoodsList()
          }).catch((error) => {
            console.log(error)
          })
        }
        if (this.operateType === 'edit') {
          // 编辑商品
          this.$axios.put('/goods/update', this.goods).then((response) => {
            if (response.data.code === 200) {
              ElMessage({message: response.data.message, type: 'success'})
            }
            this.getGoodsList()
          }).catch((error) => {
            console.log(error)
          })
        }
        this.goodsDialogVisible = false
      } else {
        console.log('error submit!')
      }
    },

    // 删除商品
    deleteGoods(row) {
      this.$axios.delete('/goods/delete', {
          params: {
            id: row.id
          }
        }).then((response) => {
          if (response.data.code === 200) {
            ElMessage({message: response.data.message, type: 'success'})
          }
          this.getGoodsList()
        }).catch((error) => {
          console.log(error)
      })
    },

    // 表单校验
    goodsFormValidator() {
      if (this.goods.categoryId === ''){
        ElMessage({message: '请选择一个类目', type: 'warning'})
        return false
      }
      if (this.goods.title === ''){
        ElMessage({message: '标题不能为空', type: 'warning'})
        return false
      }
      if (this.goods.name === ''){
        ElMessage({message: '名称不能为空', type: 'warning'})
        return false
      }
      if (this.goods.price === ''){
        ElMessage({message: '价格不能为空', type: 'warning'})
        return false
      }
      if (this.goods.quantity === ''){
        ElMessage({message: '数量不能为空', type: 'warning'})
        return false
      }
      if (this.goods.imageUrl === ''){
        ElMessage({message: '图片不能为空', type: 'warning'})
        return false
      }
      return true
    },

    // 重置查询表单
    resetForm() {
      this.$refs['query'].resetFields()
      this.getGoodsList()
    },

    // 取消
    cancel() {
      this.empty()
      this.goodsDialogVisible = false
    },

    // 清空数据
    empty() {
      this.goods.id = ''
      this.goods.categoryId = ''
      this.goods.title = ''
      this.goods.name = ''
      this.goods.price = ''
      this.goods.quantity = ''
      this.goods.picture = ''
      this.goods.remark = ''
      this.goods.status = ''
      this.operateType = ''
      this.pictureList = []
    }
  }
}
</script>
<style>
@import "../static/element.css";

.goods_query_form {
  margin: 6px 0;
  display: inline-flex;
}

.goods_image_box {
  margin: 4px 5px 5px 0;
}

.goods_image_box, .goods_image {
  width: 62px;
  height: 62px;
  border-radius: 5px;
}

.goods_title {
  font-size: 13px;
  font-weight: 450;
  margin: 2px 3px 1px;
}

.goods_id {
  margin: 0 3px;
  font-size: 12px;
  font-weight: 450;
}

.goods_image_upload_icon {
  color: #9aa9b9;
  font-size: 24px;
  font-weight: 200;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}
</style>