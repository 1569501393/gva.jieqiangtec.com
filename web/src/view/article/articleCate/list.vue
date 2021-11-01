<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="alias字段">
          <el-input v-model="searchInfo.alias" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="articleNums字段">
          <el-input v-model="searchInfo.articleNums" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="分类名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="pid字段">
          <el-input v-model="searchInfo.pid" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="seoTitle字段">
          <el-input v-model="searchInfo.seoTitle" placeholder="搜索条件" />
        </el-form-item>
            <el-form-item label="status字段" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择">
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
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
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
        <el-table-column align="left" label="alias字段" prop="alias" width="120" />
        <el-table-column align="left" label="articleNums字段" prop="articleNums" width="120" />
        <el-table-column align="left" label="分类名" prop="name" width="120" />
        <el-table-column align="left" label="pid字段" prop="pid" width="120" />
        <el-table-column align="left" label="seoDesc字段" prop="seoDesc" width="120" />
        <el-table-column align="left" label="seoKeys字段" prop="seoKeys" width="120" />
        <el-table-column align="left" label="seoTitle字段" prop="seoTitle" width="120" />
        <el-table-column align="left" label="sortOrder字段" prop="sortOrder" width="120" />
        <el-table-column align="left" label="status字段" prop="status" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="el-icon-edit" size="small" class="table-button" @click="updateCpsArticleCate(scope.row)">变更</el-button>
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
        <el-form-item label="alias字段:">
          <el-input v-model="formData.alias" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="articleNums字段:">
          <el-input v-model.number="formData.articleNums" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="pid字段:">
          <el-input v-model.number="formData.pid" clearable placeholder="请输入" />
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
        <el-form-item label="sortOrder字段:">
          <el-input v-model.number="formData.sortOrder" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="status字段:">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createCpsArticleCate,
  deleteCpsArticleCate,
  deleteCpsArticleCateByIds,
  updateCpsArticleCate,
  findCpsArticleCate,
  getCpsArticleCateList
} from '@/api/cpsArticleCate' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'CpsArticleCate',
  mixins: [infoList],
  data() {
    return {
      listApi: getCpsArticleCateList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        alias: '',
        articleNums: 0,
        name: '',
        pid: 0,
        seoDesc: '',
        seoKeys: '',
        seoTitle: '',
        sortOrder: 0,
        status: false,
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
        this.deleteCpsArticleCate(row)
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
      const res = await deleteCpsArticleCateByIds({ ids })
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
    async updateCpsArticleCate(row) {
      const res = await findCpsArticleCate({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recpsArticleCate
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        alias: '',
        articleNums: 0,
        name: '',
        pid: 0,
        seoDesc: '',
        seoKeys: '',
        seoTitle: '',
        sortOrder: 0,
        status: false,
      }
    },
    async deleteCpsArticleCate(row) {
      const res = await deleteCpsArticleCate({ ID: row.ID })
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
          res = await createCpsArticleCate(this.formData)
          break
        case 'update':
          res = await updateCpsArticleCate(this.formData)
          break
        default:
          res = await createCpsArticleCate(this.formData)
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
