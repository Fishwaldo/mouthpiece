<template>
    <CFormInput
        v-bind:type=type
        v-bind:placeholder=placeholder
        v-bind:autocomplete=autocomplete
        v-bind:valid="model.$dirty ? !model.$error : null"
        v-bind:invalid="model.$dirty ? model.$error : null"
        v-model="v$.$model"
    />
    <CFormFeedback v-for="msg in this.v$.$errors" :key="msg.id" invalid>
        {{ msg.$message }}
    </CFormFeedback>
</template>

<script>
export default {
    name: 'Validatedinput',
    props: [
        'model',
        'placeholder',
        'type',
        'autocomplete'
    ],
    data() {
        return {
            v$: this.model
        }
    }
}
</script>


