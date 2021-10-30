<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="cateId字段" prop="cateId">
          <el-select v-model="searchInfo.cateId" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true">
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="isHot字段" prop="isHot">
          <el-select v-model="searchInfo.isHot" clearable placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true">
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="seoTitle字段">
          <el-input v-model="searchInfo.seoTitle" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="searchInfo.title" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="url字段">
          <el-input v-model="searchInfo.url" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询222</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="el-icon-delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="abst字段" prop="abst" width="120" />
        <el-table-column align="left" label="addTime字段" prop="addTime" width="120" />
        <el-table-column align="left" label="附件id" prop="aid" width="120" />
        <el-table-column align="left" label="cateId字段" prop="cateId" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.cateId) }}</template>
        </el-table-column>
        <el-table-column align="left" label="数据状态" prop="dataState" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.dataState) }}</template>
        </el-table-column>
        <el-table-column align="left" label="img字段" prop="img" width="120" />
<!--        <el-table-column align="left" label="信息" prop="info" width="120" />-->
        <el-table-column align="left" label="isBest字段" prop="isBest" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isBest) }}</template>
        </el-table-column>
        <el-table-column align="left" label="isHot字段" prop="isHot" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isHot) }}</template>
        </el-table-column>
        <el-table-column align="left" label="ordid字段" prop="ordid" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.ordid) }}</template>
        </el-table-column>
        <el-table-column align="left" label="orig字段" prop="orig" width="120" />
        <el-table-column align="left" label="分销机构" prop="platformId" width="120" />
        <el-table-column align="left" label="seoDesc字段" prop="seoDesc" width="120" />
        <el-table-column align="left" label="seoKeys字段" prop="seoKeys" width="120" />
        <el-table-column align="left" label="seoTitle字段" prop="seoTitle" width="120" />
        <el-table-column align="left" label="0-待审核 1-已审核" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="title" width="120" />
        <el-table-column align="left" label="创建人" prop="uid" width="120" />
        <el-table-column align="left" label="修改时间" prop="updateTime" width="120" />
        <el-table-column align="left" label="url字段" prop="url" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateCpsArticle(scope.row)">变更</el-button>
            <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="abst字段:">
          <el-input v-model="formData.abst" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="addTime字段:">
          <el-date-picker v-model="formData.addTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="附件id:">
          <el-input v-model="formData.aid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="cateId字段:">
          <el-switch v-model="formData.cateId" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="数据状态：0删除，1正常:">
          <el-switch v-model="formData.dataState" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="img字段:">
          <el-input v-model="formData.img" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="信息:">
          <el-input v-model="formData.info" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="isBest字段:">
          <el-switch v-model="formData.isBest" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="isHot字段:">
          <el-switch v-model="formData.isHot" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="ordid字段:">
          <el-switch v-model="formData.ordid" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="orig字段:">
          <el-input v-model="formData.orig" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分销机构（发布的时候可指定全部或者具体分行、子机构的人员能看到） 0全部:">
          <el-input v-model.number="formData.platformId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="seoDesc字段:">
          <el-input v-model="formData.seoDesc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="seoKeys字段:">
          <el-input v-model="formData.seoKeys" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="seoTitle字段:">
          <el-input v-model="formData.seoTitle" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0-待审核 1-已审核:">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="修改时间:">
          <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="url字段:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createCpsArticle,
  deleteCpsArticle,
  deleteCpsArticleByIds,
  updateCpsArticle,
  findCpsArticle,
  getCpsArticleList
} from '@/api/cpsArticle' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'CpsArticle',
  mixins: [infoList],
  data() {
    return {
      listApi: getCpsArticleList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        abst: '',
        addTime: new Date(),
        aid: '',
        cateId: false,
        dataState: false,
        img: '',
        info: '',
        isBest: false,
        isHot: false,
        ordid: false,
        orig: '',
        platformId: 0,
        seoDesc: '',
        seoKeys: '',
        seoTitle: '',
        status: false,
        title: '',
        uid: 0,
        updateTime: new Date(),
        url: '',
      }
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.cateId === ""){
        this.searchInfo.cateId=null
      }
      if (this.searchInfo.dataState === ""){
        this.searchInfo.dataState=null
      }
      if (this.searchInfo.isBest === ""){
        this.searchInfo.isBest=null
      }
      if (this.searchInfo.isHot === ""){
        this.searchInfo.isHot=null
      }
      if (this.searchInfo.ordid === ""){
        this.searchInfo.ordid=null
      }
      if (this.searchInfo.status === ""){
        this.searchInfo.status=null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteCpsArticle(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteCpsArticleByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateCpsArticle(row) {
      const res = await findCpsArticle({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recpsArticle
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        abst: '',
        addTime: new Date(),
        aid: '',
        cateId: false,
        dataState: false,
        img: '',
        info: '',
        isBest: false,
        isHot: false,
        ordid: false,
        orig: '',
        platformId: 0,
        seoDesc: '',
        seoKeys: '',
        seoTitle: '',
        status: false,
        title: '',
        uid: 0,
        updateTime: new Date(),
        url: '',
      }
    },
    async deleteCpsArticle(row) {
      const res = await deleteCpsArticle({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createCpsArticle(this.formData)
          break
        case 'update':
          res = await updateCpsArticle(this.formData)
          break
        default:
          res = await createCpsArticle(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>