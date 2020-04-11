<template>
  <div class="main-page">
    <h1 class="page-title">What can I do?</h1>
    <p class="page-description">Discover methods which allow one person to has real impact to the World and our future!</p>
    <a-input-search
      placeholder="Find problem and ideas"
      v-model="searchQuery"
      @search="onSearch"
      enterButton
    />
    <br /><br />
      <TransitionFadeExpand>
        <a-list v-if="beenFinding && !!foundProblems.length" itemLayout="horizontal" :dataSource="foundProblems">
          <a-list-item slot="renderItem" slot-scope="item, index">
            <a
              slot="actions"
              @click="$router.push(`/problem/${item.id}`)"
            >Details</a>
            <a-list-item-meta
              :description="item.description.slice(0, 20) + '...'"
            >
              <a slot="title" href="https://www.antdv.com/">{{item.name}}</a>
            </a-list-item-meta>
          </a-list-item>
        </a-list>
        <a-alert
          v-else-if="beenFinding"
          message="No results"
          description="Query returned 0 results"
          type="error"
          showIcon
        />
      </TransitionFadeExpand>
    <div class="problems__more-actions" v-if="beenFinding">
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
        <a-form :form="form" v-else class="problem__add-problem-form" @submit="check">
          <a-divider orientation="left">New problem</a-divider>
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="Name"
          >
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
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="Description"
          >
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
          <a-form-item :wrapper-col="{ span: 12 }" style="margin-top: -20px;">
            <a-button type="primary" html-type="submit">
              Add a problem
            </a-button>
          </a-form-item>
        </a-form>
      </TransitionFadeExpand>
      <a-button v-if="!addingProblemForm" type="primary" class="problem__add-btn" @click.native="onclickAddProblem">
        <template v-if="addedProblem">Add another problem</template>
        <template v-else>Add a problem</template>
      </a-button>
    </div>
    </div>
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

import Vue from 'vue'
import TransitionFadeExpand from '~/components/TransitionFadeExpand.vue'

export default {

    components: {
        TransitionFadeExpand
    },

    data () {
        return {
            searchQuery: '',
            foundProblems: [],
            beenFinding: false,
            formItemLayout,
            formTailLayout,
            form: this.$form.createForm(this, { name: 'dynamic_rule' }),

            addingProblemForm: false,
            isAddingProblem: false,
            addedProblem: false
        }
    },

    methods: {

        onclickAddProblem () {
            this.addingProblemForm = true
            this.addedProblem = false
        },

        check(event) {
            event.preventDefault()
            this.form.validateFields(async err => {
                if (!err) {
                    const { description, name } = this.form.getFieldsValue(['description', 'name'])
                    this.isAddingProblem = true
                    try {
                        await this.$axios.post('/problems', {
                            description,
                            name
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
        },

      async onSearch() {
          this.$router.push({
              query: {
                  ...this.$route.query,
                  searchQuery: this.searchQuery
              }
          })
          if (!this.beenFinding) {
              this.beenFinding = true
          }

          try {
            let { problems } = await this.$axios.$get(`problems?searchQuery=${this.searchQuery.trim()}`)
            this.foundProblems = problems
        } catch (err) {
            this.foundProblems = []
            console.log(err)
        }
      }
    },

    async asyncData ({ query, error, $axios }) {
        try {
            if (!query || !query.searchQuery || !query.searchQuery.length) {
                return
            }
            let { problems } = await $axios.$get(`problems?searchQuery=${query.searchQuery}`)
            return {
                foundProblems: problems,
                beenFinding: true,
                searchQuery: query.searchQuery
            }
        } catch (err) {
            return {
                foundProblems: [],
                beenFinding: true,
                searchQuery: query.searchQuery
            }
        }
    },
}
</script>

<style lang="scss">
  .page-title {
    font-size: 33px;
    font-weight: bold;
    text-align: center;
  }

  .page-description {
    text-align: center;
  }

  .main-page {
    padding: 24px;
  }

  .problems {
    &__more-actions {
      margin-top: 15px;
    }
  }
</style>
