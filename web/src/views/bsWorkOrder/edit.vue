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
                <n-form-item label="用户Id" path="userId">
                  <n-input-number placeholder="请输入用户Id" v-model:value="formValue.userId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="创建时间" path="createTime">
                  <DatePicker v-model:formValue="formValue.createTime" type="datetime" />
                </n-form-item>
              </n-gi>
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
                  <Editor style="height: 450px" id="problemDescription" v-model:value="formValue.problemDescription" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="检测描述" path="defectDescription">
                  <Editor style="height: 450px" id="defectDescription" v-model:value="formValue.defectDescription" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="分配处理人Id" path="dispatchId">
                  <n-input-number placeholder="请输入分配处理人Id" v-model:value="formValue.dispatchId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="发起类型" path="sendType">
                  <n-input-number placeholder="请输入发起类型" v-model:value="formValue.sendType" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="收取方式" path="acceptType">
                  <n-input-number placeholder="请输入收取方式" v-model:value="formValue.acceptType" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="产品类型" path="productType">
                  <n-input-number placeholder="请输入产品类型" v-model:value="formValue.productType" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="工单状态" path="status">
                  <n-select v-model:value="formValue.status" :options="dict.getOptionUnRef('sys_normal_disable')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="建议处理方案(多选)" path="suggestSolution">
                  <n-input placeholder="请输入建议处理方案(多选)" v-model:value="formValue.suggestSolution" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="实际处理方案(多选)" path="actualSolution">
                  <n-input placeholder="请输入实际处理方案(多选)" v-model:value="formValue.actualSolution" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="总金额" path="totalMoney">
                  <n-input-number placeholder="请输入总金额" v-model:value="formValue.totalMoney" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="开票类型" path="invoiceType">
                  <n-input-number placeholder="请输入开票类型" v-model:value="formValue.invoiceType" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="备注" path="remark">
                  <Editor style="height: 450px" id="remark" v-model:value="formValue.remark" />
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
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import Editor from '@/components/Editor/editor.vue';
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