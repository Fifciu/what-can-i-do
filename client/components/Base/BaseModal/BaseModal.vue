<template>
  <a-modal
    :title="title"
    :visible="modalVisibility"
    :confirm-loading="confirmLoading"
    :centered="true"
    @ok="handleOk"
    @cancel="handleCancel"
    class="base-modal"
  >
    <a-textarea placeholder="Why do you want to discard this idea?" allow-clear />
  </a-modal>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: "BaseModal",

  data() {
    return {
      ModalText: 'Content of the modal',
      visible: false,
      confirmLoading: false,
    };
  },

  props: {
    title: {
      type: String,
      require: true
    }
  },

  computed: {
    ...mapState({
      modalVisibility: state => state.ui.modalVisibility
    }),
  },

  methods: {
    handleOk() {
      this.confirmLoading = true;
      setTimeout(() => {
        this.$store.dispatch('ui/changeModalVisibility', false)
        this.confirmLoading = false;
      }, 2000);
    },
    handleCancel() {
      this.$store.dispatch('ui/changeModalVisibility', false)
    }
  }
}
</script>

<style lang="scss">
.base-modal {
  .ant-modal-header,
  .ant-modal-footer {
    border: none;
  }
}
</style>

