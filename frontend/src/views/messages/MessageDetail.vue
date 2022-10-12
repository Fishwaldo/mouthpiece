<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-12">
        <CCardHeader>
          <strong>Message {{message.ID}}</strong>
        </CCardHeader>
        <CCardBody>
          <CContainer>
            <CRow>
              <CCol sm="2"><b>Subject:</b></CCol><CCol sm="8">{{message.Subject}}</CCol>
            </CRow>
            <CRow>
              <CCol sm="2"><b>Message:</b></CCol><CCol sm="8">{{message.Message}}</CCol>
            </CRow>
            <CRow>
              <CCol sm="2"><b>Topic:</b></CCol><CCol sm="8">{{message.Topic}}</CCol>
            </CRow>
            <CRow>
              <CCol sm="2"><b>Severity:</b></CCol><CCol sm="8">{{message.Severity}}</CCol>
            </CRow>
            <CRow>
              <CCol sm="2"><b>TimeStamp:</b></CCol><CCol sm="8">{{message.TimeStamp}}</CCol>
            </CRow>
            <CRow>
              <CCol sm="2"><b>Fields:</b></CCol><CCol sm="8">{{message.Fields}}</CCol>
            </CRow>
            <CRow>
              <CCol sm="2"><b>Application:</b></CCol><CCol sm="8" ><router-link :to="'/apps/' + App.id">{{App.name}}</router-link></CCol>
            </CRow>
          </CContainer>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
//import { msgResponse } from '@/generated/'



export default {
  name: 'MessageDetails',
  data () {
    return {
      message: {
        ID: String,
        appid: Number,
        id: String,
        message: String,
        severity: Number,
        shortmsg: String,
        timestamp: String,
        topic: String,
      },
      App: {
        ID: String,
        Name: String,
        Description: String,
      }, 
    }
  },
  setup() {
  },
  methods: {
  },
  mounted() {
    console.log(this.$route.params.messageID)
    this.$store.cache.dispatch('FETCH_MSG', this.$route.params.messageID)
    .then((response) => {
      console.log(response)
      this.message = response
      this.$store.cache.dispatch('FETCH_APP', this.message.AppID)
      .then((response) => {
        console.log(response)
        this.App = response
      })
    })
    .catch((error) => {
      console.log("error", error)
    })
  },
}
</script>
