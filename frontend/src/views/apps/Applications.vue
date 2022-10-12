<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-4">
        <CCardHeader>
          <CNavbar color-scheme="light" class="bg-light">
            <CContainer fluid>
              <CNavbarBrand href="#">Application List</CNavbarBrand>
              <CForm class="d-flex">
                <CButton type="submit" color="success" variant="outline" @click="newClicked">New</CButton>
              </CForm>
            </CContainer>
          </CNavbar>
          <strong></strong>
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
                  <CTableDataCell><router-link :to="'/apps/' + apps.ID">{{apps.Name}}</router-link></CTableDataCell>
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
  methods: {
    newClicked() {
      this.$router.push('/apps/new')
    }
  }
}
</script>
