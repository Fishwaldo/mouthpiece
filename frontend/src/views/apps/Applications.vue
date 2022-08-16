<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-4">
        <CCardHeader>
          <strong>Application List</strong>
        </CCardHeader>
        <CCardBody>
            <CTable>
              <CTableHead>
                <CTableRow>
                  <CTableHeaderCell scope="col">Name</CTableHeaderCell>
                  <CTableHeaderCell scope="col">Description</CTableHeaderCell>
                  <CTableHeaderCell scope="col">Status</CTableHeaderCell>
                </CTableRow>
              </CTableHead>
              <CTableBody>
                <CTableRow v-for="apps in applications" v-bind:key="apps.AppName">
                  <CTableDataCell>{{apps.AppName}}</CTableDataCell>
                  <CTableDataCell>{{apps.description}}</CTableDataCell>
                  <CTableDataCell>{{apps.Status}}</CTableDataCell>
                </CTableRow>
              </CTableBody>
            </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
import { DefaultService } from '@/generated/'

export default {
  name: 'Applications',
  data () {
    return {
      applications: []
    }
  },
  mounted() {
    console.log("Heloo world")
    DefaultService.getApps().then(response => {
      this.applications = response
      console.log(response)
    })
  },
}
</script>
