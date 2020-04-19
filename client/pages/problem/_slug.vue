<template>
  <div class="problem">

    <div class="problem__details">
      <div class="problem__description">
        <h2 class="problem__section-heading">{{ problem.name }}</h2>
        <p>{{ problem.description }}</p>
      </div>
    </div>

    <div class="problem__ideas">
      <h2 class="problem__section-heading">Ideas</h2>
      <a-list
        v-if="problem.ideas.length"
        class="comment-list"
        itemLayout="horizontal"
        :dataSource="problem.ideas"
      >
        <a-list-item slot="renderItem" slot-scope="item, index">
          <a-comment author="Han Solo">
<!--            <template slot="actions">-->
<!--              <span key="comment-basic-like">-->
<!--                <a-tooltip title="Like">-->
<!--                  <a-icon type="like" :theme="action === 'liked' ? 'filled' : 'outlined'" @click="like" />-->
<!--                </a-tooltip>-->
<!--                <span style="padding-left: '8px';cursor: 'auto'">-->
<!--                  {{likes}}-->
<!--                </span>-->
<!--              </span>-->
<!--                      <span key="comment-basic-dislike">-->
<!--                <a-tooltip title="Dislike">-->
<!--                  <a-icon-->
<!--                    type="dislike"-->
<!--                    :theme="action === 'disliked' ? 'filled' : 'outlined'"-->
<!--                    @click="dislike"-->
<!--                  />-->
<!--                </a-tooltip>-->
<!--                <span style="padding-left: '8px';cursor: 'auto'">-->
<!--                  {{dislikes}}-->
<!--                </span>-->
<!--              </span>-->
<!--              <span key="comment-basic-reply-to">Fake news</span>-->
<!--            </template>-->
            <p slot="content">{{item.description}}</p>
            <a-tooltip slot="datetime">
              <span>{{item.datetime}}</span>
            </a-tooltip>
          </a-comment>
        </a-list-item>
      </a-list>
      <div
        v-else
        class="problem__no-ideas"
        message="No ideas"
        description="There are no ideas to solve this problem. Be the first who shares an brilliant idea!"
        type="warning"
        showIcon
      >
        There are no ideas to solve this problem. Be the first who shares an <strong>brilliant</strong> idea!
      </div>

      <TransitionFadeExpand>
        <div v-if="isAddingIdea">
          <a-spin />
        </div>
        <div v-else-if="addedIdea">
          <h2>Succesfully added an idea</h2>
          <p>Thank you. We are really glad you took part in our project! We are going to analyze your idea before publishing it to prevent fake newses.</p>
        </div>
        <a-form :form="form" v-else-if="showAddIdea" class="problem__add-ideas-form">
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
              placeholder="What impact my action would have?"
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

      </TransitionFadeExpand>

      <div class="problem__add-btn-wrapper">
        <div v-if="!isLoggedIn">
          <h3>Do you want to add own idea?</h3>
          <a-button type="primary" class="problem__add-btn" html-type="submit">
            <nuxt-link :to="`/sign-in?back-type=problem&back-slug=${problem.slug}`">
              Sign in at first
            </nuxt-link>
          </a-button>
        </div>
        <a-button type="primary" class="problem__add-btn" @click.native="onclickAddIdea" v-else>
          <template v-if="addedIdea">Add another idea</template>
          <template v-else>Add an idea</template>
        </a-button>
      </div>
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

    import moment from 'moment'
    import TransitionFadeExpand from '~/components/TransitionFadeExpand.vue'

    export default {
        name: "ProblemView",
        components: {
            TransitionFadeExpand
        },
        data () {
            return {
                showAddIdea: false,

                formItemLayout,
                formTailLayout,
                form: this.$form.createForm(this, { name: 'dynamic_rule' }),

                addedIdea: false,
                isAddingIdea: false

            }
        },
        computed: {
            isLoggedIn () {
                return this.$store.getters['auth/isLoggedIn']
            }
        },
        async asyncData ({ store, params, error, $axios }) {
            try {
                let { data } = await $axios.get(`/problems/${params.slug}/ideas`)
                store.dispatch('breadcrumbs/changeTitle', data.problem.name)
                return {
                     problem: {
                         ...data.problem,
                         ideas: data.problem.ideas.map(idea => ({
                             ...idea,
                             datetime: moment().subtract(1, 'days').fromNow()
                         }))
                     }
                }
            } catch (err) {
                error({ statusCode: 404, message: err + 'Problem not found' })
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

<style lang="scss" scoped>
  .problem {
    &__details {

     }

    &__title {
      color: #fff;
      font-weight: bold;
      margin: 0;
    }

    &__ideas {
      padding: 10px 30px;
    }

    &__no-ideas {
      padding: 1em 1em 0;
      margin: 0;
    }

    &__add-btn {
      margin: 20px 0;
    }

    &__add-btn-wrapper {
      text-align: center;
    }

    &__add-ideas-form {
      padding: 18px;
    }

    &__heading {
      display: flex;
      align-items: center;
      margin-bottom: 5px;
      color: #f0f0f0;
      background: dodgerblue;
      padding: 24px;
    }

    &__back-icon {
      margin-right: 15px;
      font-size: 1.5em;
    }

    &__description {
      font-size: 18px;
      line-height: 1.6;
      padding: 10px 30px;
    }

    &__section-heading {
      font-size: 30px;
    }
  }
</style>