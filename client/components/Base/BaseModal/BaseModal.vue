<template>
  <a-modal
    :title="title"
    :visible="modalVisibility"
    :confirm-loading="modalButtonLoader"
    okText="Send"
    :centered="true"
    @ok="handleOk"
    @cancel="handleCancel"
    class="base-modal"
  >
    <slot name="content" />
  </a-modal>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: "BaseModal",

  props: {
    title: {
      type: String,
      require: true
    }
  },

  computed: {
    ...mapState({
      modalVisibility: state => state.ui.modalVisibility,
      modalButtonLoader: state => state.ui.modalButtonLoader
    }),
  },

  methods: {
    handleOk() {
      this.$emit('handleOk')
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

