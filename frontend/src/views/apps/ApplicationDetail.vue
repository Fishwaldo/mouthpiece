<template>
  <ErrorModal v-model:error="formError" />
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-12">
        <CCardHeader>
          <strong>Application Details</strong>
        </CCardHeader>
        <CCardBody>
          <CContainer fluid>
            <CRow>
              <CCol>
                <CContainer fluid>
                  <CRow>
                    <CCol>
                      <CImage thumbnail :src="'/api/avatar/app/'+ App.id" width="200" height="100" />
                    </CCol>
                  </CRow>
                  <CRow>
                    <CCol>
                      {{App.name}}
                    </CCol>
                  </CRow>
                </CContainer>
              </CCol>
              <CCol sm="10">
                <AppForm v-model:App="App" v-bind:ButtonText="SubmitButtonText" @submit="onSubmit"/>
              </CCol>
            </CRow>
          </CContainer>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-12">
        <CCardHeader>
          <strong>Filters</strong>
        </CCardHeader>
        <CCardBody>
          <CContainer>
            <CRow v-for="flt in App.filters" :key="flt.id">
              <CCol sm="1"><b>Name:</b></CCol><CCol sm="2"><router-link :to="'/groups/' + flt.id">{{flt.name}}</router-link></CCol><CCol sm="2"><b>Description</b></CCol><CCol sm="7">{{flt.description}}</CCol>
            </CRow>
          </CContainer>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>

  <CRow>
    <CCol :xs="12">
      <CCard class="mb-12">
        <CCardHeader>
          <strong>Groups</strong>
        </CCardHeader>
        <CCardBody>
          <CContainer>
            <CRow v-for="grp in App.groups" :key="grp.id">
              <CCol sm="1"><b>Name:</b></CCol><CCol sm="2"><router-link :to="'/groups/' + grp.ID">{{grp.Name}}</router-link></CCol><CCol sm="2"><b>Description</b></CCol><CCol sm="7">{{grp.description}}</CCol>
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
      Orig: {} as appGetResponse,
      formError: {} as ApiError,
      SubmitButtonText: 'Save' as String,
    }
  },
  methods: {
    onSubmit (appData: appGetResponse) {
      this.SubmitButtonText = 'Saving...'
      this.$store.dispatch('PUT_APP', {id: this.App.id, payload: appData})
      .then((response: appGetResponse) => {
        this.App = response
        this.Orig = Object.assign({}, response)
        this.$store.cache.delete('FETCH_APP', this.App.id)
        this.SubmitButtonText = ''
      })
      .catch((error: ApiError) => {
        this.SubmitButtonText = ''
        this.formError = error
      })
    },
  },
  mounted() {
    this.$store.cache.dispatch('FETCH_APP', this.$route.params.appID)
    .then((response: appGetResponse) => {
      this.App = response
      this.Orig = Object.assign({}, response)
    })
    .catch((error: ApiError) => {
      this.formError = error
    })
  },
})
</script>
