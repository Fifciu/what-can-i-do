<template>
  <div class="problem">

    <div class="problem__details">
      <div class="problem__heading">
        <a-icon type="arrow-left" class="problem__back-icon" @click="$router.push('/')"/>
        <h2 class="problem__title">{{ problem.title }}</h2>
      </div>
      <p class="problem__description">{{ problem.description }}</p>
    </div>

    <div class="problem__ideas">
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
        <a-form :form="form" v-if="showAddIdea" class="problem__add-ideas-form">
          <h2>New idea</h2>
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="Description"
          >
            <a-textarea
              v-decorator="[
              'description',
              { rules: [{ required: true, message: 'Please input idea\'s description' }] },
            ]"
              placeholder="Please input idea's description"
              :autoSize="{ minRows: 3, maxRows: 5 }"
            />
          </a-form-item>
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="Price"
          >
            <a-input-number
              :defaultValue="0"
              :formatter="value => `$ ${(+value/100).toFixed(2)}`"
              :parser="value => value.replace('$ ', '').replace('.','')"
            />
          </a-form-item>
        </a-form>
      </TransitionFadeExpand>

      <div class="problem__add-btn-wrapper">
        <a-button type="primary" class="problem__add-btn" @click.native="onclickAddIdea">Add an idea</a-button>
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

                checkNick: false,
                formItemLayout,
                formTailLayout,
                form: this.$form.createForm(this, { name: 'dynamic_rule' })
            }
        },
        async asyncData ({ params, error, $axios }) {
            try {
                let { data } = await $axios.get(`/problems/${params.id}/ideas`)
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
                } else {
                    this.check()
                }
            },
            check() {
                this.form.validateFields(err => {
                    if (!err) {
                        console.info('success');
                    }
                });
            },
            handleChange(e) {
                this.checkNick = e.target.checked;
                this.$nextTick(() => {
                    this.form.validateFields(['nickname'], { force: true });
                });
            },
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
      padding: 0 6px 6px;
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
  }
</style>
