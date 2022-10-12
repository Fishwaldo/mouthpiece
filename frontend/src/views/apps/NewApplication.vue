<template>
  <ErrorModal v-model:error="formError" />
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-12">
        <CCardHeader>
          <strong>New Application</strong>
        </CCardHeader>
        <CCardBody>
          <CContainer fluid>
            <CRow>
              <CCol sm="10">
                <AppForm v-model:App="App" @submit="onSubmit"/>
              </CCol>
            </CRow>
          </CContainer>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { appGetResponse, ApiError } from '@/generated/'
import { default as AppForm } from '@/components/AppForm.vue'
import { default as ErrorModal } from '@/components/ErrorModal.vue'
//import {diff, jsonPatchPathConverter} from 'just-diff'




export default defineComponent({
  name: 'AppDetails',
  components: {
    AppForm,
    ErrorModal
},
  props: ['appID'],
  data () {
    return {
      App: {} as appGetResponse,
      formError: {} as ApiError,
    }
  },
  methods: {
    onSubmit (appData: appGetResponse) {
      this.$store.dispatch('POST_APP', {payload: appData}).then((response) => {
        this.$router.push('/apps/'+response.id)
      })
      .catch((error: ApiError) => {
        this.formError = error
      })
    },
  },
})
</script>
