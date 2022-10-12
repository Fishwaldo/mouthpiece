<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-12">
        <CCardHeader>
          <strong>Messages List</strong>
        </CCardHeader>
        <CCardBody>
          <div class="table table-bordered table-sm" id="example-table"></div>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
import {TabulatorFull as Tabulator} from 'tabulator-tables';
import { DateTime } from 'luxon';


export default {
  name: 'Messages',
  data () {
    return {
      tabulator: null,
    }
  },
  setup() {
    window.DateTime = DateTime
  },
  methods: {
  // eslint-disable-next-line no-unused-vars
  lookupApp(cell, formaterparams, onRendered) {
    let ce = cell
      this.$store.cache.dispatch('FETCH_APP', cell.getValue())
      .then((response) => {
         ce.getElement().innerHTML = response.name
        })
      .catch((error) => {
        console.log("error", error)
        return  "Unknown"
      })
    return ce.getValue()
    },   
    openMessage(e, row) {
      this.$router.push("/messages/" + row.getData().ID)
    },
  },

  mounted() {
    let vm = this
    const token = this.$auth.token().split(';')[0]
    
    this.tabulator = new Tabulator("#example-table", {
      height: "1024px",
      layout: "fitDataStretch",
      placeholder: "No Data Set",
      pagination:true,
      paginationMode:"remote",
      sortMode:"remote",
      ajaxURL:"/api/messages",
      ajaxConfig: {
        headers: {
          'Authorization': 'Bearer ' + token
        }
      },
      ajaxURLGenerator: function (url, config, params) {
        // Get ajaxURL
        let myUrl = url;
        
        // If sorting, then get the field name and direction
        if (params['sort'].length > 0) {
          let field = params['sort'][0]['field'];
          let dir = params['sort'][0]['dir'];
          params['orderBy'] = field;
          params['orderDir'] = dir;
        }
        delete params['sort'];
        // Return request URL
        return myUrl + '?' + new URLSearchParams(params).toString();
      },
      ajaxResponse: function(url, params, response){
        //url - the URL of the request
        //params - the parameters passed with the request
        //response - the JSON object returned in the body of the response.
        // response.data.forEach(function (item) {
        //   vm.$store.cache.dispatch('FETCH_MSG', item.ID)
        // })
        response.last_page = Math.ceil(response.count/params.size)
        return response; //return the tableData property of a response json object
      },
      paginationSize:20,
      paginationSizeSelector:true,
      paginationCounter:"rows",
      columns: [
        { title: "Time", field: "TimeStamp",  formatter:"datetime", formatterParams:{ inputFormat:"iso", invalidPlaceholder:"(invalid date)"}},
        { title: "App", field: "AppID", width:150, formatter:this.lookupApp},
        { title: "Subject", field: "ShortMsg"},
        { title: "Severity", field: "Severity", formatter:"star", formatterParams: {stars: 5}},
        { title: "Full Message", field: "Message"},
      ],
    });
    this.tabulator.on("rowClick", vm.openMessage)
  },
}
</script>
