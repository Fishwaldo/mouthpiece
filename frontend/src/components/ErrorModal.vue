<template>
  <CModal backdrop="static" :visible="visibleStaticBackdropDemo" @close="() => { visibleStaticBackdropDemo = false }">
    <CModalHeader>
      <CModalTitle>ERROR: {{error.body.title || "Error - " + error.status }}</CModalTitle>
    </CModalHeader>
    <CModalBody>
        <CAlert color="danger">
            {{error.body.detail || error.message}}:
            <ul>
                <li v-for="detail in error.body.errors" :key="detail.id">
                    {{detail.message}}
                </li>
            </ul>
        </CAlert>
    </CModalBody>
    <CModalFooter>
      <CButton color="secondary" @click="() => { visibleStaticBackdropDemo = false }">
        Close
      </CButton>
    </CModalFooter>
  </CModal>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { ApiError } from '@/generated/'

export default defineComponent({
  name: 'ErrorModal',
  props: {
    error: ApiError,
  },
  data() {
    return { 
        visibleStaticBackdropDemo: false,
    }
  },
  watch: {
    error: function () {
      console.log("Error Changed", this.error)
      this.visibleStaticBackdropDemo = true
    },
  },
})
</script>
