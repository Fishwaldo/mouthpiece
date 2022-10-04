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
                  <CTableHeaderCell scope="col">ID</CTableHeaderCell>
                  <CTableHeaderCell scope="col">Name</CTableHeaderCell>
                  <CTableHeaderCell scope="col">Description</CTableHeaderCell>
                </CTableRow>
              </CTableHead>
              <CTableBody>
                <CTableRow v-for="apps in applications" v-bind:key="apps.ID">
                  <CTableDataCell>{{apps.ID}}</CTableDataCell>
                  <CTableDataCell>{{apps.Name}}</CTableDataCell>
                  <CTableDataCell>{{apps.Description}}</CTableDataCell>
                </CTableRow>
              </CTableBody>
            </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
import { AppsService } from '@/generated/'

export default {
  name: 'Applications',
  data () {
    return {
      applications: []
    }
  },
  mounted() {
    AppsService.getApps().then(response => {
      this.applications = response
    })
  },
}
</script>
