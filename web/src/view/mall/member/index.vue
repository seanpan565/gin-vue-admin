<!-- 商城会员管理 -->
<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="手机号">
          <el-input v-model="searchInfo.mobile" placeholder="搜索手机号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="全部" clearable>
            <el-option label="正常" :value="1" />
            <el-option label="冻结" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" style="width: 100%" row-key="ID">
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="手机号" prop="mobile" width="140" />
        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <el-table-column align="left" label="邮箱" prop="email" min-width="160" show-overflow-tooltip />
        <el-table-column align="left" label="状态" width="90">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '正常' : '冻结' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="注册来源" prop="registerSource" width="110" />
        <el-table-column align="left" label="最后登录" width="180">
          <template #default="scope">
            {{ scope.row.lastLoginAt ? formatDate(scope.row.lastLoginAt) : '-' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="注册时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="100" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="showDetail(scope.row)">详情</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer v-model="detailVisible" title="会员详情" size="480px">
      <el-descriptions v-if="detail" :column="1" border>
        <el-descriptions-item label="ID">{{ detail.ID }}</el-descriptions-item>
        <el-descriptions-item label="UUID">{{ detail.uuid }}</el-descriptions-item>
        <el-descriptions-item label="手机号">{{ detail.mobile }}</el-descriptions-item>
        <el-descriptions-item label="昵称">{{ detail.nickname || '-' }}</el-descriptions-item>
        <el-descriptions-item label="邮箱">{{ detail.email || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="detail.status === 1 ? 'success' : 'danger'">
            {{ detail.status === 1 ? '正常' : '冻结' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="注册来源">{{ detail.registerSource }}</el-descriptions-item>
        <el-descriptions-item label="注册 IP">{{ detail.registerIp || '-' }}</el-descriptions-item>
        <el-descriptions-item label="最后登录 IP">{{ detail.lastLoginIp || '-' }}</el-descriptions-item>
        <el-descriptions-item label="最后登录">
          {{ detail.lastLoginAt ? formatDate(detail.lastLoginAt) : '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="注册时间">{{ formatDate(detail.CreatedAt) }}</el-descriptions-item>
        <template v-if="detail.profile">
          <el-descriptions-item label="性别">
            {{ genderText(detail.profile.gender) }}
          </el-descriptions-item>
          <el-descriptions-item label="生日">
            {{ detail.profile.birthday ? formatDate(detail.profile.birthday) : '-' }}
          </el-descriptions-item>
          <el-descriptions-item label="积分">{{ detail.profile.points ?? 0 }}</el-descriptions-item>
          <el-descriptions-item label="等级 ID">{{ detail.profile.levelId ?? '-' }}</el-descriptions-item>
        </template>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
  import { getMemberList, getMemberDetail } from '@/api/mall/member'
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { formatDate } from '@/utils/format'

  defineOptions({
    name: 'MallMember'
  })

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({ mobile: '', status: null })
  const detailVisible = ref(false)
  const detail = ref(null)

  const genderText = (g) => {
    if (g === 1) return '男'
    if (g === 2) return '女'
    return '未知'
  }

  const getTableData = async () => {
    const res = await getMemberList({
      page: page.value,
      pageSize: pageSize.value,
      mobile: searchInfo.value.mobile,
      status: searchInfo.value.status
    })
    if (res.code === 0) {
      tableData.value = res.data.list || []
      total.value = res.data.total
      page.value = res.data.page
      pageSize.value = res.data.pageSize
    }
  }

  const onSubmit = () => {
    page.value = 1
    getTableData()
  }

  const onReset = () => {
    searchInfo.value = { mobile: '', status: null }
    page.value = 1
    getTableData()
  }

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const showDetail = async (row) => {
    const res = await getMemberDetail({ id: row.ID })
    if (res.code === 0) {
      detail.value = res.data
      detailVisible.value = true
    } else {
      ElMessage.error(res.msg || '获取详情失败')
    }
  }

  getTableData()
</script>
