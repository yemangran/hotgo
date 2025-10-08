<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑工单管理 #' + formValue.id : '添加工单管理'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="客户名称" path="customerName">
                  <n-input placeholder="请输入客户名称" v-model:value="formValue.customerName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="客户地址" path="customerAddress">
                  <n-input placeholder="请输入客户地址" v-model:value="formValue.customerAddress" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="客户联系方式" path="customerContact">
                  <n-input placeholder="请输入客户联系方式" v-model:value="formValue.customerContact" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="客户对接人" path="customerPerson">
                  <n-input placeholder="请输入客户对接人" v-model:value="formValue.customerPerson" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="问题描述" path="problemDescription">
                  <n-input type="textarea" placeholder="问题描述" v-model:value="formValue.problemDescription" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="检测描述" path="defectDescription">
                  <n-input type="textarea" placeholder="检测描述" v-model:value="formValue.defectDescription" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="处理人" path="dispatchId">
                  <n-select v-model:value="formValue.dispatchId" options="" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="发起类型" path="sendType">
                  <n-select v-model:value="formValue.sendType" :options="dict.getOptionUnRef('send_type')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="收取方式" path="acceptType">
                  <n-select v-model:value="formValue.acceptType" :options="dict.getOptionUnRef('accept_type')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="产品类型" path="productType">
                  <n-select v-model:value="formValue.productType" :options="dict.getOptionUnRef('product_type')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="工单状态" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('work_status')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="建议处理方案" path="suggestSolution">
                  <n-select v-model:value="formValue.suggestSolution" :options="dict.getOptionUnRef('biz_solution')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="实际处理方案" path="actualSolution">
                  <n-select v-model:value="formValue.actualSolution" :options="dict.getOptionUnRef('biz_solution')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="总金额" path="totalMoney">
                  <n-input-number placeholder="请输入总金额" v-model:value="formValue.totalMoney" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="开票类型" path="invoiceType">
                  <n-select v-model:value="formValue.invoiceType" :options="dict.getOptionUnRef('invoice_type')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="备注" path="remark">
                  <n-input type="textarea" placeholder="备注" v-model:value="formValue.remark" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
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

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { useDictStore } from '@/store/modules/dict';
  import { Edit, View } from '@/api/bsWorkOrder';
  import { State, newState, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

      return;
    }

    // 编辑
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
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

<style lang="less"></style>