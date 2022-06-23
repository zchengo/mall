<template>
  <Container>
    <!-- 订单查询 -->
    <el-form :inline="true" :model="query" ref="query" class="goods_query_form">
      <el-form-item prop="id">
        <el-input v-model.number="query.id" placeholder="订单ID"/>
      </el-form-item>
      <el-form-item prop="status">
        <el-select v-model="query.status" placeholder="订单状态">
          <el-option label="待付款" value="1"></el-option>
          <el-option label="已取消" value="2"></el-option>
          <el-option label="已付款" value="3"></el-option>
          <el-option label="配送中" value="4"></el-option>
          <el-option label="已完成" value="5"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" :icon="Search" @click="getOrderList">查询</el-button>
        <el-button type="primary" :icon="Brush" @click="resetForm"/>
      </el-form-item>
    </el-form>
    <!-- 订单列表 -->
    <el-table :data="orderList" height="70vh" style="width: 100%">
      <el-table-column prop="id" label="订单号" width="200px"/>
      <el-table-column prop="totalPrice" label="订单金额"/>
      <el-table-column prop="goodsCount" label="商品数量">
        <template #default="scope">x{{scope.row.goodsCount}}</template>
      </el-table-column>
      <el-table-column prop="status" label="订单状态">
        <template #default="scope">
          <el-tag v-if="scope.row.status === 1" size="small" type="info">待付款</el-tag>
          <el-tag v-if="scope.row.status === 2" size="small" type="info">已取消</el-tag>
          <el-tag v-if="scope.row.status === 3" size="small" type="info">已付款</el-tag>
          <el-tag v-if="scope.row.status === 4" size="small" type="primary">配送中</el-tag>
          <el-tag v-if="scope.row.status === 5" size="small" type="success">已完成</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created" label="创建时间" min-width="150" sortable>
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-icon><timer/></el-icon>
            <span style="margin-left: 10px">{{ scope.row.created }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="130">
        <template #default="scope">
          <el-button size="small" @click="checkOrder(scope.$index, scope.row)">详情</el-button>
          <el-button v-if="scope.row.status === 3" size="small" type="primary" @click="modifyOrder(scope.row, 4)">配送</el-button>
          <el-popconfirm title="此操作将永久删除该信息, 是否继续?"
                         confirmButtonText="确认"
                         cancelButtonText="取消"
                         cancelButtonType="default"
                         v-if="scope.row.status === 2 || scope.row.status === 5"
                         :icon="WarningFilled"
                         @confirm="deleteOrder(scope.row)">
            <template #reference>
             <el-button size="small" type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
      <template #empty>
        <div style="margin: 50px 0;">
          <el-empty  v-if="showEmpty" description="暂时还没有订单哦" />
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
    <!-- 查看订单详情，对话框 -->
    <el-dialog :title="dialogTitle" v-model="orderDialogVisible" top="5vh" width="50%">
      <Descriptions label="订单编号">{{orderDetail.id}}</Descriptions>
      <Descriptions label="订单状态">
        <el-tag v-if="orderDetail.status === 1" size="small" type="warning">待付款</el-tag>
        <el-tag v-if="orderDetail.status === 2" size="small" type="info">已取消</el-tag>
        <el-tag v-if="orderDetail.status === 3" size="small" type="success">已付款</el-tag>
        <el-tag v-if="orderDetail.status === 4" size="small" type="primary">配送中</el-tag>
        <el-tag v-if="orderDetail.status === 5" size="small" type="success">已完成</el-tag>
      </Descriptions>
      <Descriptions label="支付金额">{{orderDetail.totalPrice}}</Descriptions>
      <Descriptions label="商品列表">
        <el-table :data="orderDetail.goodsItem" style="width: 100%">
          <el-table-column prop="name" label="主图" width="60">
            <template #default="scope">
              <div style="display: flex;justify-content: center;align-items: center;">
                <el-image class="goods_image" :src="scope.row.imageUrl"/>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="title" label="标题">
            <template #default="scope">
              <div style="font-size: 10px;line-height: 15px;">{{scope.row.title}}</div>
            </template>
          </el-table-column>
          <el-table-column prop="count" label="数量" width="80">
            <template #default="scope">x {{scope.row.count}}</template>
          </el-table-column>
        </el-table>
      </Descriptions>
      <Descriptions label="收货人姓名">{{orderDetail.name}}</Descriptions>
      <Descriptions label="手机号">{{orderDetail.mobile}}</Descriptions>
      <Descriptions label="收货地址">{{orderDetail.province + ' ' + orderDetail.city + ' ' + orderDetail.district +
      ' ' + orderDetail.detailedAddress }}</Descriptions>
      <Descriptions label="创建时间">{{orderDetail.created}}</Descriptions>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="orderDialogVisible = false">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </Container>
</template>

<script>
import Container from "../components/Container";
import {Brush, Delete, Search, Setting, Timer, WarningFilled} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";
import Descriptions from "@/components/Descriptions";

export default {
  name: "Order",
  components: {Descriptions, Container, Timer},
  setup() {
    return {Search, Brush, Setting, Delete, WarningFilled}
  },
  data() {
    return {
      orderList: [],
      order: {
        id: '',
        amount: '',
        goodsCount: '',
        status: '',
        shippingAddress: ''
      },
      query: {
        id: '',
        status: ''
      },
      orderDetail: {
        id: '',
        status: '',
        totalPrice: '',
        goodsItem: [],
        name: '',
        mobile: '',
        province: '',
        city: '',
        district: '',
        detailedAddress: '',
        created: '',
        nickName: ''
      },
      dialogTitle: '',
      operateType: '',
      orderDialogVisible: false,

      // 分页
      total: 0,
      pageNum: 1,
      pageSize: 12,

      // 空状态
      showEmpty: false
    }
  },
  mounted() {
    this.getOrderList()
  },
  methods: {
    // 分页，处理函数
    handleCurrentChangePrev(val) {
      this.pageNum = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.pageNum = val;
      this.getOrderList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.pageNum = val;
      console.log(`下一页: ${val}`);
    },

    // 获取订单列表
    getOrderList(){
      this.$axios.get('/order/list', {
        params: {
          id: this.query.id,
          status: this.query.status,
          pageNum: this.pageNum,
          pageSize: this.pageSize
        }
      }).then((response) => {
        this.total = response.data.data.total
        this.orderList = response.data.data.list
        if (this.orderList.length === 0) {
          this.showEmpty = true
        }
      }).catch((error) => {
        console.log(error)
      })
    },

    // 查看订单详情
    checkOrder(index, row) {
      this.orderDetail.nickName = row.nickName
      console.log(row.id)
      this.dialogTitle = '订单详情'
      this.$axios.get('/order/detail', {
        params: { id: row.id }
      }).then((response) => {
        this.orderDetail.id = response.data.data.id
        this.orderDetail.status = response.data.data.status
        this.orderDetail.totalPrice = response.data.data.totalPrice
        this.orderDetail.goodsItem = response.data.data.goodsItem
        this.orderDetail.name = response.data.data.name
        this.orderDetail.mobile = response.data.data.mobile
        this.orderDetail.province = response.data.data.province
        this.orderDetail.city = response.data.data.city
        this.orderDetail.district = response.data.data.district
        this.orderDetail.detailedAddress = response.data.data.detailedAddress
        this.orderDetail.created = response.data.data.created
      }).catch((error) => {
        console.log(error)
      })
      this.orderDialogVisible = true
    },

    // 修改订单
    modifyOrder(row, status) {
      this.$axios.put('/order/update', {
        id: row.id,
        status: status
      }).then((response) => {
        if (response.data.code === 200 && status === 4){
          ElMessage({message: '订单配送中', type: 'success'})
        }
        this.getOrderList()
      })
    },

    // 重置查询表单
    resetForm() {
      this.$refs['query'].resetFields()
      this.getOrderList()
    },

    // 删除订单
    deleteOrder(row){
      this.$axios.delete('/order/delete', {
          params: {
            id: row.id
          }
        }).then((response) => {
          if (response.data.code === 200) {
            ElMessage({message: response.data.message, type: 'success'})
          }
          this.getOrderList()
        }).catch((error) => {
          console.log(error)
      })
    }
  }
}
</script>

<style scoped>
.goods_image{
  width: 35px;
  height: 35px;
  border-radius: 5px;
}
</style>