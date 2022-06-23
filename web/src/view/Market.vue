<template>
  <Container>
    <!-- 活动查询 -->
    <el-form :inline="true" :model="query" ref="query" class="goods_query_form">
      <el-form-item prop="id">
        <el-input v-model.number="query.id" placeholder="活动ID"/>
      </el-form-item>
      <el-form-item prop="status">
        <el-select v-model="query.type" placeholder="活动类型">
          <el-option label="钻石展位" value="1"></el-option>
          <el-option label="限时特惠" value="2"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="status">
        <el-select v-model="query.status" placeholder="活动状态">
          <el-option label="进行中" value="1"></el-option>
          <el-option label="已结束" value="2"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" :icon="Search" @click="getMarketList">查询</el-button>
        <el-button type="primary" :icon="Brush" @click="resetForm"/>
        <el-button type="primary" :icon="Plus" @click="add">活动</el-button>
      </el-form-item>
    </el-form>
    <!-- 活动列表 -->
    <el-table :data="marketList" height="70vh" style="width: 100%">
     <el-table-column prop="bannerImage" label="活动图片" min-width="50">
        <template #default="scope">
          <el-image :src="scope.row.bannerImage" style="border-radius: 5px;"/>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="活动名称"/>
      <el-table-column prop="type" label="活动类型" min-width="50">
        <template #default="scope">
          <el-tag v-if="scope.row.type === 1" type="primary">新品推荐</el-tag>
          <el-tag v-if="scope.row.type === 2" type="primary">限时特惠</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="beginTime" label="开始时间" min-width="100">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-icon>
              <timer/>
            </el-icon>
            <span style="margin-left: 10px">{{ scope.row.beginTime }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="overTime" label="结束时间" min-width="100">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-icon>
              <timer/>
            </el-icon>
            <span style="margin-left: 10px">{{ scope.row.overTime }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" min-width="50">
        <template #default="scope">
          <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" @click="changeStatus(scope.row)"/>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" @click="edit(scope.row)">编辑</el-button>
          <el-popconfirm title="此操作将永久删除该信息, 是否继续?"
                         confirmButtonText="确认"
                         cancelButtonText="取消"
                         cancelButtonType="default"
                         :icon="WarningFilled"
                         @confirm="deleteMarket(scope.row)">
            <template #reference>
             <el-button size="small" type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
      <template #empty>
        <div style="margin: 50px 0;">
          <el-empty v-if="showEmpty" description="暂时还没有活动哦" />
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
    <!-- 添加、编辑营销，通用对话框 -->
    <el-dialog :title="dialogTitle" v-model="marketDialogVisible" :lock-scroll="false" top="5vh" width="45%"
               @close="cancel">
      <el-form :model="market" label-width="80px">
        <el-form-item label="活动名称" :required="true">
          <el-input v-model="market.name" type="text" maxlength="20" show-word-limit/>
        </el-form-item>
        <el-form-item label="活动类型" :required="true">
          <el-select v-model.number="market.type" placeholder="请选择" clearable>
            <el-option v-for="item in typeOption" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <el-form-item label="活动图片" :required="true">
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
        <el-form-item label="开始时间" :required="true">
          <el-date-picker v-model="market.beginTime" value-format="YYYY-MM-DD HH:mm:ss" type="datetime" placeholder="请选择开始时间"/>
        </el-form-item>
        <el-form-item label="结束时间" :required="true">
          <el-date-picker v-model="market.overTime" value-format="YYYY-MM-DD HH:mm:ss" type="datetime" placeholder="请选择结束时间"/>
        </el-form-item>
        <el-form-item label="关联商品" :required="true">
          <el-button type="primary" :icon="Plus" @click="addGoods"/>
        </el-form-item>
        <el-form-item label="活动商品">
          <el-table :data="selectedGoodsList" style="width: 100%" empty-text="您还没有添加活动商品哦">
            <el-table-column prop="name" label="主图" width="60">
              <template #default="scope">
                <div style="display: flex;justify-content: center;align-items: center;">
                  <el-image class="goods_image" :src="scope.row.imageUrl"/>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="title" label="标题">
              <template #default="scope">
                <div style="font-size: 10px;line-height: 15px;">{{ scope.row.title }}</div>
              </template>
            </el-table-column>
            <el-table-column prop="price" label="价格" width="80">
              <template #default="scope">¥ {{ scope.row.price }}</template>
            </el-table-column>
            <el-table-column label="操作" min-width="50">
              <template #default="scope">
                <el-button size="small" type="danger" @click="removeSelectedGoods(scope.row)">移除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
          <el-button @click="empty">重置</el-button>
          <el-button type="primary" @click="confirmMarket">确定</el-button>
        </span>
      </template>
    </el-dialog>
    <!-- 商品列表 -->
    <el-dialog :title="dialogTitle" v-model="goodsDialogVisible" :lock-scroll="false" top="5vh" width="50%"
               @close="goodsDialogVisible = false">
      <el-table :data="goodsList" height="70vh" style="width: 100%" @select="selectedGoodsIds">
        <el-table-column type="selection" width="55"/>
        <el-table-column prop="name" label="主图" width="60">
          <template #default="scope">
            <div style="display: flex;justify-content: center;align-items: center;">
              <el-image class="goods_image" :src="scope.row.imageUrl"/>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题">
          <template #default="scope">
            <div style="font-size: 10px;line-height: 15px;">{{ scope.row.title }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="价格" width="80">
          <template #default="scope">¥ {{ scope.row.price }}</template>
        </el-table-column>
      </el-table>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="goodsDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmSelectedGoods">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </Container>
</template>

<script>
import Container from "@/components/Container";
import {Brush, Minus, Plus, Search, Timer, WarningFilled} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";

export default {
  name: "Market",
  components: {Container, Timer},
  setup() {
    return {Search, Brush, Plus, Minus, WarningFilled}
  },
  data() {
    return {
      query: {
        id: '',
        type: '',
        status: ''
      },
      marketList: [],
      goodsList: [],
      market: {
        id: '',
        bannerImage: '',
        name: '',
        type: '',
        beginTime: '',
        overTime: '',
        goodsIds: '',
        status: ''
      },
      typeOption: [{
        value: 1,
        label: '钻石展位'
      }, {
        value: 2,
        label: '限时特惠'
      }],

      // 图片列表
      pictureList: [],

      goodsIds: [],
      selectedGoodsList: [],

      marketId: '',

      // 分页
      total: 0,
      pageNum: 1,
      pageSize: 10,

      // 操作状态
      dialogTitle: '',
      operateType: '',
      marketDialogVisible: false,
      goodsDialogVisible: false,

      // 请求认证
      token: '',

      // 空状态
      showEmpty: false
    }
  },
  mounted() {
    this.getMarketList()
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
      this.getMarketList()
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.pageNum = val;
      console.log(`下一页: ${val}`);
    },

    // 图片上传，处理函数
    handleImagePreview(file) {
      this.market.bannerImage = file.url
    },
    handleImageRemove(file, fileList) {
      this.market.bannerImage = ''
      console.log(file, fileList)
    },
    handleImageSuccess(response) {
      this.market.bannerImage = response.data
    },

    // 重置查询表单
    resetForm() {
      this.$refs['query'].resetFields()
      this.getMarketList()
    },

    // 添加活动
    add() {
      this.dialogTitle = '添加活动'
      this.operateType = 'add'
      this.selectedGoodsList = []
      this.marketDialogVisible = true
    },

    // 编辑活动
    edit(row) {
      this.dialogTitle = '编辑活动'
      this.operateType = 'edit'
      this.market.id = row.id
      this.market.name = row.name
      this.market.type = row.type
      this.market.bannerImage = row.bannerImage
      this.pictureList.push({url: row.bannerImage})
      this.market.beginTime = row.beginTime
      this.market.overTime = row.overTime
      this.market.goodsIds = row.goodsIds
      this.market.status = row.status
      this.goodsIds = row.goodsIds.split(',')
      this.getSelectedGoodsList(row.goodsIds)
      this.marketDialogVisible = true
    },

    // 获取营销活动列表
    getMarketList() {
      this.$axios.get('/market/list', {
        params: {
          id: this.query.id,
          type: this.query.type,
          status: this.query.status,
          pageNum: this.pageNum,
          pageSize: this.pageSize
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.marketList = response.data.data.list;
        if (this.marketList.length === 0) {
          this.showEmpty = true
        }
      }).catch((error) => {
        console.log(error)
      })
    },

    // 开启、关闭活动
    changeStatus(row){
      console.log(row.status)
      this.$axios.put('/market/status/update', {
        id: row.id,
        status: parseInt(row.status),
      }).then((response) => {
        if (response.data.code === 200 && row.status === 1){
          ElMessage({message: '活动已开启', type: 'success'})
        }
        if (response.data.code === 200 && row.status === 2){
          ElMessage({message: '活动已关闭', type: 'success'})
        }
      }).catch((error) => {
        console.log(error)
      })
    },

    // 添加关联商品
    addGoods() {
      this.getGoodsList()
      this.dialogTitle = '添加关联商品'
      this.goodsDialogVisible = true
    },

    // 选中时，获取商品ID
    selectedGoodsIds(selection, row) {
      this.goodsIds.push(row.id)
    },

    // 确认选中的商品
    confirmSelectedGoods() {
      this.selectedGoodsList = []
      this.market.goodsIds = this.goodsIds.toString()
      this.goodsList.filter(item => {
        this.goodsIds.some(id => { if (id === item.id){ this.selectedGoodsList.push(item) } })
      })
      this.goodsDialogVisible = false
    },

    // 移除选中的商品
    removeSelectedGoods(row){
      this.selectedGoodsList.splice(this.selectedGoodsList.indexOf(row.id), 1)
      this.goodsIds.splice(this.goodsIds.indexOf(row.id), 1)
      this.market.goodsIds = this.goodsIds.toString()
      console.log(this.goodsIds)
    },

    // 获取已上架商品列表
    getGoodsList() {
      this.$axios.get('/market/goods/list', {
        params: { goodsIds: "" }
      }).then((response) => {
        this.goodsList = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },

    // 获取已关联的商品列表
    getSelectedGoodsList(goodsIds) {
      this.$axios.get('/market/goods/list', {
        params: { goodsIds: goodsIds }
      }).then((response) => {
        this.selectedGoodsList = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },

    // 确认添加或编辑活动
    confirmMarket() {
      let valid = this.marketFormValidator();
      if (valid) {
        if (this.operateType === 'add') {
          if (this.market.goodsIds === "") {
            ElMessage({message: '请至少关联一个活动商品', type: 'warning'})
            return
          }
          // 添加活动
          this.$axios.post('/market/create', {
            name: this.market.name,
            type: this.market.type,
            bannerImage: this.market.bannerImage,
            beginTime: this.market.beginTime,
            overTime: this.market.overTime,
            goodsIds: this.market.goodsIds
          }).then((response) => {
            if (response.data.code === 200) {
              ElMessage({message: response.data.message, type: 'success'})
            }
            this.getMarketList()
          }).catch((error) => {
            console.log(error)
          })
        }
        if (this.operateType === 'edit') {
          // 编辑活动
          this.$axios.put('/market/update', {
            id: this.market.id,
            name: this.market.name,
            type: this.market.type,
            bannerImage: this.market.bannerImage,
            beginTime: this.market.beginTime,
            overTime: this.market.overTime,
            goodsIds: this.market.goodsIds,
            status: this.market.status
          }).then((response) => {
            if (response.data.code === 200) {
              ElMessage({message: response.data.message, type: 'success'})
            }
            this.getMarketList()
          }).catch((error) => {
            console.log(error)
          })
        }
        this.marketDialogVisible = false
      } else {
        console.log("error submit!!");
      }
    },

    // 删除活动
    deleteMarket(row) {
      this.$axios.delete('/market/delete', {
          params: {
            id: row.id
          }
        }).then((response) => {
          if (response.data.code === 200) {
            ElMessage({message: response.data.message, type: 'success'})
          }
          this.getMarketList()
        }).catch((error) => {
          console.log(error)
      })
    },

    // 表单校验
    marketFormValidator() {
      if (this.market.name === '') {
        ElMessage({ message: '名称不能为空', type: 'warning' });
        return false;
      }
      if (this.market.type === '') {
        ElMessage({ message: '类型不能为空', type: 'warning' });
        return false;
      }
      if (this.market.bannerImage === '') {
        ElMessage({ message: '图片不能为空', type: 'warning' });
        return false;
      }
      if (this.market.beginTime === '') {
        ElMessage({ message: "请设置开始时间", type: 'warning' });
        return false;
      }
      if (this.market.overTime === "") {
        ElMessage({ message: "请设置结束时间", type: 'warning' });
        return false;
      }
      if (this.market.goodsIds === "") {
        ElMessage({ message: "至少关联一个商品", type: 'warning' });
        return false;
      }
      return true;
    },

    // 取消
    cancel() {
      this.empty()
      this.marketDialogVisible = false
    },

    // 清空数据
    empty() {
      this.market.bannerImage = '',
      this.market.name = '',
      this.market.type = '',
      this.market.beginTime = '',
      this.market.overTime = '',
      this.market.goodsIds = '',
      this.pictureList = []
    },
  }
}
</script>

<style scoped>
.goods_image {
  width: 35px;
  height: 35px;
  border-radius: 5px;
}
</style>