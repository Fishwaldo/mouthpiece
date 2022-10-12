<template>
    <CFormSwitch size="lg" 
        v-bind:valid="model.$dirty ? !model.$error : null"
        v-bind:invalid="model.$dirty ? model.$error : null"
        v-bind:label="cblabel"
        v-model="enabled"
    />
    <CFormFeedback v-for="msg in this.v$.$errors" :key="msg.id" invalid>
        {{ msg.$message }}
    </CFormFeedback>
</template>

<script type="ts">

export default {
    name: 'Validatedswitch',
    props: [
        'model',
        'label',
    ],
    data() {
        return {
            v$: this.model
        }
    },
    computed: {
        cblabel: {
            get() {
                return this.v$.$model
            },
            set() {
            }
        },
        enabled: {
            get() {
                return this.v$.$model == "Enabled" ? true : false
            },
            set(value) {
                this.v$.$model = value ? "Enabled" : "Disabled"
            }
        }
    }
}
</script>


