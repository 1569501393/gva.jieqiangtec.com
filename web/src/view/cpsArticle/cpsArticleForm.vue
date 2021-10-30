<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="abst字段:">
          <el-input v-model="formData.abst" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="addTime字段:">
          <el-date-picker v-model="formData.addTime" type="date" placeholder="选择日期" clearable></el-date-picker>
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
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="url字段:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
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
  createCpsArticle,
  updateCpsArticle,
  findCpsArticle
} from '@/api/cpsArticle' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'CpsArticle',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findCpsArticle({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.recpsArticle
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
