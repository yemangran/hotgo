<template>
  <div>
    <n-modal v-model:show="showModal" title="工单处理" :mask-closable="false" :show-icon="false" preset="dialog"
      transform-origin="center" :style="{
        width: dialogWidth,
      }">
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-divider title-placement="left">
            客户信息
          </n-divider>
          <n-descriptions label-placement="left" bordered :column="2">
            <n-descriptions-item label="客户名称">
              {{ formValue.customerName }}
            </n-descriptions-item>
            <n-descriptions-item label="联系方式">
              {{ formValue.customerContact }}
            </n-descriptions-item>
            <n-descriptions-item label="客户地址">
              {{ formValue.customerAddress }}
            </n-descriptions-item>
            <n-descriptions-item label="对接人">
              {{ formValue.customerPerson }}
            </n-descriptions-item>
            <n-descriptions-item colspan="2" label="问题描述">
              {{ formValue.problemDescription }}
            </n-descriptions-item>
          </n-descriptions>
          <n-divider title-placement="left">
            工单处理
          </n-divider>
          <n-form ref="formRef" :model="formValue" :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'" label-width="auto" class="py-4">
            <n-grid cols="2" responsive="screen">
              <n-gi span="2">
                <n-form-item label="检测描述" path="defectDescription">
                  <n-input :autosize="{ minRows: 2 }" type="textarea" placeholder="请输入检测描述"
                    v-model:value="formValue.defectDescription" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="建议处理方案" path="defectDescription">
                  <n-select multiple v-model:value="formValue.suggestSolution"
                    :options="dict.getOptionUnRef('biz_solution')" />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="实际处理方案" path="defectDescription">
                  <n-select multiple v-model:value="formValue.actualSolution"
                    :options="dict.getOptionUnRef('biz_solution')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="收取方式" path="acceptType">
                  <n-select v-model:value="formValue.acceptType" :options="dict.getOptionUnRef('accept_type')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="开票类型" path="invoiceType">
                  <n-select v-model:value="formValue.invoiceType" :options="dict.getOptionUnRef('invoice_type')" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
          <n-divider title-placement="left">
            处理明细
          </n-divider>
          <n-button strong tertiary type="primary" @click="addService">
            <template #icon>
              <NIcon>
                <AddIcon />
              </NIcon>
            </template>
            添加
          </n-button>
          <n-data-table :columns="processColumns" :data="processData" :bordered="false" />
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { adaModalWidth } from '@/utils/hotgo';
import { ref, computed } from 'vue';
import { State, newState, rules, createProcessColumns, RowData } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { View, Process } from '@/api/bsWorkOrder';
import { useDictStore } from '@/store/modules/dict';
import { AddOutline as AddIcon } from '@vicons/ionicons5'
import { NIcon, useDialog } from 'naive-ui'
import { List } from '@/api/bsService';
import { convertListToTree } from '@/utils/hotgo';

const dialog = useDialog();
const dict = useDictStore();
const settingStore = useProjectSettingStore();
const emit = defineEmits(['success']);
const formRef = ref<any>(null);
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});
const processData = ref<any[]>([]);
const serviceOptions = ref<any[]>([]);
const formBtnLoading = ref(false);

async function init() {
  const res = await List({});
  serviceOptions.value = convertListToTree(res.list, 'id', 'pid', true);
}
init()

function closeForm() {
  showModal.value = false;
}

async function confirmForm() {
  // 表单验证
  await formRef.value?.validate();
  
  // 校验明细数据
  if (!processData.value || processData.value.length === 0) {
    window['$message'].warning('请至少添加一条处理明细');
    return;
  }
  
  // 校验明细项目是否完整
  for (let i = 0; i < processData.value.length; i++) {
    const item = processData.value[i];
    if (!item.serviceId) {
      window['$message'].warning(`请选择第${i + 1}条明细的服务项目`);
      return;
    }
  }

  formBtnLoading.value = true;
  try {
    // 构建提交数据
    const submitData = {
      ...formValue.value,
      serviceList: processData.value,
    };

    // 调用后端接口
    await Process(submitData);
    
    window['$message'].success('工单处理成功');
    closeForm();
    // 通知父组件刷新列表
    emit('success');
  } catch (error) {
    console.error('工单处理失败:', error);
  } finally {
    formBtnLoading.value = false;
  }
}

// 删除处理明细（带确认提示）
function deleteService(row: RowData, index: number) {
  dialog.warning({
    title: '确认删除',
    content: '确定要删除这条明细吗',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      processData.value.splice(index, 1);
    },
  });
}

// 创建表格列配置
const processColumns = computed(() => createProcessColumns(deleteService, serviceOptions.value));

// 打开模态框
async function openModal(state: State) {
  // initUserList()
  showModal.value = true;
  //todo)) 这里从后端获取数据
  processData.value = [];

  // 编辑
  loading.value = true;
  const res = await View({ id: state.id })
  formValue.value = res;
  loading.value = false;

  // // 新增
  // if (!state || state.id < 1) {
  //     formValue.value = newState(state);

  //     return;
  // }

  // // 编辑
  // loading.value = true;
  // View({ id: state.id })
  //     .then((res) => {
  //         formValue.value = res;
  //     })
  //     .finally(() => {
  //         loading.value = false;
  //     });
}

function addService() {
  processData.value.push({
    name: '',
    serviceId: null,
    workOrderId: null,
    handleType: '',
    price: 0,
    remark: '',
  });
}

defineExpose({
  openModal,
});
</script>