<template>
  <div v-if="isAddingIdea">
    <a-spin />
  </div>
  <div v-else-if="addedIdea">
    <h2>Succesfully added an idea</h2>
    <p>Thank you. We are really glad you took part in our project! We are going to analyze your idea before publishing it to prevent fake newses.</p>
  </div>
  <a-form :form="form" v-else class="problem__add-ideas-form">
    <h2>New idea</h2>
    <a-form-item
      :label-col="formItemLayout.labelCol"
      :wrapper-col="formItemLayout.wrapperCol"
      label="Action's description"
    >
      <a-textarea
        v-decorator="[
                'action_description',
                { rules:
                  [
                    { required: true, message: 'Please input idea\'s description' },
                    { validator: (rule, value) => value.length >= 15 || !value.length, message: 'Description should have at least 15 chars' }
                  ]
                },
              ]"
        placeholder="What could I do?"
        :autoSize="{ minRows: 3, maxRows: 5 }"
      />
    </a-form-item>
    <a-form-item
      :label-col="formItemLayout.labelCol"
      :wrapper-col="formItemLayout.wrapperCol"
      label="Result's description"
    >
      <a-textarea
        v-decorator="[
                'results_description',
                { rules:
                  [
                    { required: true, message: 'Please input idea\'s description' },
                    { validator: (rule, value) => value.length >= 15 || !value.length, message: 'Description should have at least 15 chars' }
                  ]
                },
              ]"
        placeholder="What impact would my action have?"
        :autoSize="{ minRows: 3, maxRows: 5 }"
      />
    </a-form-item>
    <a-form-item
      :label-col="formItemLayout.labelCol"
      :wrapper-col="formItemLayout.wrapperCol"
      label="Money Price"
    >
      <a-input-number
        v-decorator="[
                'money_price',
                { initialValue: 0 }
              ]"
        :min="0"
        :formatter="value => `$ ${(+value/100).toFixed(2)}`"
        :parser="value => isNaN(value.replace('$ ', '').replace('.','')) ? 0 : value.replace('$ ', '').replace('.','')"
      />
    </a-form-item>
    <a-form-item
      :label-col="formItemLayout.labelCol"
      :wrapper-col="formItemLayout.wrapperCol"
      label="Time price (minutes)"
    >
      <a-input-number
        v-decorator="[
                'time_price',
                { initialValue: 0 }
              ]"
        :min="0"
      />
    </a-form-item>
  </a-form>
</template>

<script>
    const formItemLayout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 8 },
    };
    const formTailLayout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 8, offset: 4 },
    };

    export default {
        name: "AddIdea",
        data () {
            return {
                showAddIdea: true,
                addedIdea: false,
                isAddingIdea: false
            }
        },
        methods: {
            onclickAddIdea () {
                if (!this.showAddIdea) {
                    this.showAddIdea = true
                    this.addedIdea = false
                } else {
                    this.check()
                }
            },
            check() {
                this.form.validateFields(async err => {
                    if (!err) {
                        const {
                            action_description,
                            results_description,
                            money_price,
                            time_price
                        } = this.form.getFieldsValue(['action_description', 'results_description', 'money_price', 'time_price'])
                        const token = this.$store.getters['auth/token']
                        if (!token) {
                            return
                        }
                        this.isAddingIdea = true
                        try {
                            await this.$axios.post('/ideas', {
                                problem_id: Number(this.problem.id),
                                action_description,
                                results_description,
                                money_price: money_price / 100,
                                time_price
                            }, {
                                headers: {
                                    'Authorization': `Bearer ${token}`
                                }
                            })

                            this.addedIdea = true
                            this.showAddIdea = false

                        } catch (err) {
                            this.$notification.error({
                                message: 'Error',
                                description: err.response.data.message
                            });
                        }
                        this.isAddingIdea = false
                    }
                });
            }
        }
    }
</script>

<style scoped>

</style>
