<template>
  <div>
    <n-modal v-model:show="showModal" :mask-closable="false" :show-icon="false" preset="dialog"
      transform-origin="center" :title="formValue.id > 0 ? '编辑服务项目 #' + formValue.id : '添加服务项目'" :style="{
        width: dialogWidth,
      }">
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form ref="formRef" :model="formValue" :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'" :label-width="100" class="py-4">
            <n-divider title-placement="left">
              <n-icon :size="20" :component="DatabaseOutlined" />
              服务信息
            </n-divider>
            <n-grid cols="2" responsive="screen">
              <n-gi span="2">
                <n-form-item label="上级服务" path="pid">
                  <n-tree-select :options="treeOption" v-model:value="formValue.pid" key-field="id" label-field="name"
                    clearable filterable default-expand-all show-path />
                </n-form-item>
              </n-gi>
              <n-gi span="2">
                <n-form-item label="服务名称" path="name">
                  <n-input placeholder="请输入服务名称" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="价格" path="price">
                  <n-input-number :min="0" placeholder="请输入价格" v-model:value="formValue.price" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="序号" path="orderNum">
                  <n-input-number :min="0" placeholder="请输入序号" v-model:value="formValue.orderNum" />
                </n-form-item>
              </n-gi>
            </n-grid>
            <n-divider title-placement="left">
              <n-icon :size="20" :component="MoneyCollectOutlined" />
              会员等级
            </n-divider>
            <n-button strong tertiary type="primary" @click="addLevel">
              新增
            </n-button>
            <n-data-table :columns="levelColumns" :data="levelData" :bordered="false" />
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
import { ref, computed, h } from 'vue';
import { Edit, View } from '@/api/bsService';
import { State, newState, treeOption, loadTreeOption, rules } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';
import { MoneyCollectOutlined, DatabaseOutlined } from '@vicons/antd'
import { TableAction } from '@/components/Table';
import { useDictStore } from '@/store/modules/dict'

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const dict = useDictStore()

const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});

const levelColumns = [
  {
    title: '会员等级',
    key: 'dictLabel'
  },
  {
    title: '折扣比值',
    key: 'ratio'
  },
  {
    width: 216,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record: State) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '重置',
            onClick: handleReset.bind(null, record),
            auth: ['/bsService/edit'],
          }
        ],
      });
    }
  }
]
const levelData = ref<Recordable[]>([])

function addLevel() {
  const opts = dict.getOptionUnRef('biz_vip_level') || []
  const exist = new Set(levelData.value.map((d: any) => d.dictValue ?? d.dictLabel))
  const rows: any[] = []
  for (const opt of opts) {
    const key = (opt as any).value ?? (opt as any).label
    if (exist.has(key)) continue
    rows.push({ dictValue: key, dictLabel: (opt as any).label, ratio: 100 })
  }
  levelData.value = [...levelData.value, ...rows]
}

//重置折扣为100%
function handleReset(record: Recordable) {
  ;(record as any).ratio = 100
}

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

  // 加载关系树选项
  loadTreeOption();
  dict.loadOptions(['biz_vip_level'])

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