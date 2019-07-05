<template>
  <tr>
    <td>{{ device.serialNumber }}</td>
    <td>{{ device.registrationDate }}</td>
    <td>{{ device.firmware }}</td>
    <td>{{ device.status }}</td>
    <td>
      <b-button
        variant="default"
        @click="onDetailsClick">Details</b-button>
    </td>
    <!-- <td>
      <b-button
        variant="success"
        @click="onUpdateClick">Update</b-button>
    </td>
    <td>
      <b-button
        variant="danger"
        @click="onDeleteClick">Delete</b-button>
    </td> -->
  </tr>
</template>
<script>
export default {
  name: 'DeviceListRow',
  props: [ 'device' ],
  methods: {
    onDetailsClick () {
      this.$emit('details', this.device.serialNumber) // emit new event on clicck
    }
    // ,
    // onUpdateClick () {
    //   this.$emit('update', this.device.id)
    // },
    // onDeleteClick () {
    //   this.$emit('delete', this.device.id)
    // }
  }
}
</script>
