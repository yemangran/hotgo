<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="工单管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                客户名称
              </template>
              {{ formValue.customerName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                客户地址
              </template>
              {{ formValue.customerAddress }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                客户联系方式
              </template>
              {{ formValue.customerContact }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                客户对接人
              </template>
              {{ formValue.customerPerson }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                问题描述
              </template>
              <span v-html="formValue.problemDescription"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                检测描述
              </template>
              <span v-html="formValue.defectDescription"></span>
            </n-descriptions-item>
            <n-descriptions-item label="处理人">
              <n-tag type="primary" size="small" class="min-left-space">
                {{ formValue.dispatchName }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="发起类型">
              <n-tag :type="dict.getType('send_type', formValue.sendType)" size="small" class="min-left-space">
                {{ dict.getLabel('send_type', formValue.sendType) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="收取方式">
              <n-tag :type="dict.getType('accept_type', formValue.acceptType)" size="small" class="min-left-space">
                {{ dict.getLabel('accept_type', formValue.acceptType) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="产品类型">
              <n-tag :type="dict.getType('product_type', formValue.productType)" size="small" class="min-left-space">
                {{ dict.getLabel('product_type', formValue.productType) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="工单状态">
              <n-tag :type="dict.getType('work_status', formValue.status)" size="small" class="min-left-space">
                {{ dict.getLabel('work_status', formValue.status) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="建议处理方案">
              <template v-for="(item, key) in formValue.suggestSolution" :key="key">
                <n-tag :type="dict.getType('biz_solution', item)" size="small" class="min-left-space">
                  {{ dict.getLabel('biz_solution', item) }}
                </n-tag>
              </template>
            </n-descriptions-item>
            <n-descriptions-item label="实际处理方案">
              <template v-for="(item, key) in formValue.actualSolution" :key="key">
                <n-tag :type="dict.getType('biz_solution', item)" size="small" class="min-left-space">
                  {{ dict.getLabel('biz_solution', item) }}
                </n-tag>
              </template>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                总金额
              </template>
              {{ formValue.totalMoney }}
            </n-descriptions-item>
            <n-descriptions-item label="开票类型">
              <n-tag :type="dict.getType('invoice_type', formValue.invoiceType)" size="small" class="min-left-space">
                {{ dict.getLabel('invoice_type', formValue.invoiceType) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                备注
              </template>
              <span v-html="formValue.remark"></span>
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { View } from '@/api/bsWorkOrder';
import { State, newState } from './model';
import { adaModalWidth } from '@/utils/hotgo';
import { getFileExt } from '@/utils/urlUtils';
import { useDictStore } from '@/store/modules/dict';

const message = useMessage();
const dict = useDictStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref(newState(null));
const dialogWidth = computed(() => {
  return adaModalWidth(580);
});
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

// 下载
function download(url: string) {
  window.open(url);
}

// 打开模态框
function openModal(state: State) {
  showModal.value = true;
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      console.log('工单详情', res);

      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped></style>