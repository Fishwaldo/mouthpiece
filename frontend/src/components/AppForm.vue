<template>
    <CForm class="row g-3 needs-validation" novalidate>
        <CRow class="mb-3">
        <CFormLabel  for="inputName" class="col-sm-2 col-form-label">Name:</CFormLabel>
        <CCol sm="10">
            <Validatedinput v-bind:model="v$.App.name" type="text" placeholder="Name"/> 
        </CCol>
        </CRow>
        <CRow class="mb-3">
        <CFormLabel for="inputDescription" class="col-sm-2 col-form-label">Description:</CFormLabel>
        <CCol sm="10">
            <Validatedinput v-bind:model="v$.App.description" type="text" placeholder="description"/>
        </CCol>
        </CRow>
        <CRow class="mb-3">
        <CFormLabel for="inputStatus" class="col-sm-2 col-form-label">Status:</CFormLabel>
        <CCol sm="10">
            <Validatedswitch label="Enabled" v-bind:model="v$.App.status" />
        </CCol>
        </CRow>
        <CRow class="mb-3">
        <CFormLabel for="inputURL" class="col-sm-2 col-form-label">URL:</CFormLabel>
        <CCol sm="10">
            <Validatedinput v-bind:model="v$.App.URL" type="text" placeholder="URL"/>
        </CCol>
        </CRow>
        <CRow class="mb-3">
        <CFormLabel for="inputICON" class="col-sm-2 col-form-label">Application Icon:</CFormLabel>
        <CCol sm="10">
            <Validatedinput v-bind:model="v$.App.Icon" type="text" placeholder="Icon"/>
        </CCol>
        </CRow>
        <CRow>
        <CCol xs="12" class="align-self-end">
            <CButton color="primary" type="submit" @click.prevent="onSubmit" :disabled="v$.$invalid">{{ButtonText || "Save"}}</CButton>
        </CCol>
        </CRow>
    </CForm>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { $appGetResponse, appGetResponse } from '@/generated/'
import useVuelidate from '@vuelidate/core'
import { or, sameAs, maxLength, minLength, required, url, helpers } from '@vuelidate/validators'
import { default as Validatedinput } from '@/components/Validatedinput.vue'
import { default as Validatedswitch } from '@/components/Validatedswitch.vue'


export default defineComponent({
  name: 'AppDetails',
  components: {
    Validatedinput,
    Validatedswitch,
  },
  props: {
    App: {
        type: Object as () => appGetResponse,
    },
    ButtonText: {
        type: String,
        default: "Submit"
    }
  },
  data () {
    return {
      v$: useVuelidate(),
    }
  },
  validations () {
    return {
      App: {
        name: {
          required,
          pattern: helpers.withMessage('Name must be a single word', helpers.regex(RegExp($appGetResponse.properties.name.pattern))),
          maxLength: maxLength($appGetResponse.properties.name.maxLength),
          minLength: minLength($appGetResponse.properties.name.minLength),
        },
        description: {
          required,
          maxLength: maxLength($appGetResponse.properties.description.maxLength),
          pattern: helpers.withMessage(({
            $model
            }) => `This field has a value of '${$model}' but it cannot contain special characters`, helpers.regex(RegExp($appGetResponse.properties.description.pattern))),
        },
        status: {
          enum: or(sameAs("Enabled"), sameAs("Disabled"))
        },
        URL: {
          url
        },
        Icon: {
          url
        },
      }
    }
  },
  methods: {
    onSubmit () {
      this.v$.$touch()
      this.$emit('submit', this.App)
    },
  },
  mounted() {
    if (this.v$.App.status.$model === undefined) {
        this.v$.App.status.$model = "Enabled"
    }
    console.log(this.v$.$error)
  },
})
</script>
