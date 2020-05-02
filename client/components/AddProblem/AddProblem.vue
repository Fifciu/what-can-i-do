<template>
  <div class="add-problem">
    <TransitionFadeExpand>
      <div v-if="isAddingProblem">
        <a-spin />
      </div>
      <div v-else-if="addedProblem">
        <h2>Succesfully added a problem</h2>
        <p>Thank you. We are really glad you took part in our project! We are going to analyze your problem before publishing it to prevent fake newses.</p>
      </div>
      <div v-else-if="!addingProblemForm">
        <a-divider orientation="left">Not found your problem?</a-divider>
        <p>Feel free to describe problem that does not exist in our database yet.</p>
      </div>
      <a-form
        v-else
        :form="form"
        class="problem__add-problem-form"
        @submit="check"
      >
        <a-divider orientation="left">New problem</a-divider>
        <a-form-item>
          <a-input
            v-decorator="[
                'name',
                { rules:
                  [
                    { required: true, message: 'Please input problem\'s name' },
                    { validator: (rule, value) => value.length >= 4 || !value.length, message: 'Name should have at least 4 chars' }
                  ]
                },
              ]"
            placeholder="Please input problem's name"
          />
        </a-form-item>
        <a-form-item>
          <a-textarea
            v-decorator="[
                'description',
                { rules:
                  [
                    { required: true, message: 'Please input problem\'s description' },
                    { validator: (rule, value) => value.length >= 15 || !value.length, message: 'Description should have at least 15 chars' }
                  ]
                },
              ]"
            placeholder="Please input problem's description"
            :autoSize="{ minRows: 3, maxRows: 5 }"
          />
        </a-form-item>
        <a-form-item
          :wrapper-col="{ span: 12 }"
        >
          <a-button
            type="primary"
            html-type="submit"
          >
            Add a problem
          </a-button>
        </a-form-item>
      </a-form>
    </TransitionFadeExpand>
    <a-button
      type="primary"
      html-type="submit"
      v-if="!isLoggedIn"
    >
      <nuxt-link to="/sign-in">
        Sign in at first
      </nuxt-link>
    </a-button>
    <a-button
      v-else-if="!addingProblemForm"
      type="primary"
      class="problem__add-btn"
      @click.native="onclickAddProblem"
    >
      <template v-if="addedProblem">Add another problem</template>
      <template v-else>Add a problem</template>
    </a-button>
  </div>
</template>

<script>
    export default {
        name: "AddProblem",

        data() {
            return {
                form: this.$form.createForm(this, {name: 'dynamic_rule'}),

                addingProblemForm: false,
                isAddingProblem: false,
                addedProblem: false
            }
        },

        computed: {
            isLoggedIn() {
                return this.$store.getters['auth/isLoggedIn']
            }
        },

        methods: {
            onclickAddProblem() {
                this.addingProblemForm = true
                this.addedProblem = false
            },

            check(event) {
                event.preventDefault()
                this.form.validateFields(async err => {
                    if (!err) {
                        const {description, name} = this.form.getFieldsValue(['description', 'name'])
                        const token = this.$store.getters['auth/token']
                        if (!token) {
                            return
                        }
                        this.isAddingProblem = true
                        try {
                            await this.$axios.post('/problems', {
                                description,
                                name
                            }, {
                                headers: {
                                    'Authorization': `Bearer ${token}`
                                }
                            })

                            this.addedProblem = true
                            this.addingProblemForm = false

                        } catch (err) {
                            this.$notification.error({
                                message: 'Error',
                                description: err.response.data.message
                            });
                        }
                        this.isAddingProblem = false
                    }
                });
            }

        }
    }
</script>

<style scoped>

</style>
