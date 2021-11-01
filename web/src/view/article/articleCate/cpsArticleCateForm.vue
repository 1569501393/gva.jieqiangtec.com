<template>
  <div>
    <div class="gva-form-box">
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createCpsArticleCate,
  updateCpsArticleCate,
  findCpsArticleCate
} from '@/api/cpsArticleCate' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'CpsArticleCate',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findCpsArticleCate({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.recpsArticleCate
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
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
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
