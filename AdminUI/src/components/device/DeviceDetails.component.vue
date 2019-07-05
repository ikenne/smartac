<template>
  <div>
    <div class="well">
      <div class="row">
        <div class="col-md-3">
          <strong>Device serial number:</strong>
        </div>
        <div class="col-md-3">
          {{ device.serialNumber }}
        </div>
      </div>
      <div class="row">
        <div class="col-md-3">
          <strong>Registration Date:</strong>
        </div>
        <div class="col-md-3">
          {{ device.registrationDate }}
        </div>
      </div>
      <div class="row">
        <div class="col-md-3">
          <strong>Firmware:</strong>
        </div>
        <div class="col-md-3">
          {{ device.firmware }}
        </div>
      </div>
      <div class="row">
        <div class="col-md-3">
          <strong>Status:</strong>
        </div>
        <div class="col-md-3">
          {{ device.status }}
        </div>
      </div>
    </div>
    <div
        v-if="device.temperature && device.temperature.length > 0"
        class="row">
      <div class="col-md-3">
        <strong>Temperature:</strong>
      </div>
      <div class="col-md-12">
        <div class="table-responsive">
          <table class="table table-striped">
            <thead>
              <tr>
                <th>Time Acquired</th>
                <th>Sensor value</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="temp in device.temperature"
                :key="temp.date">
                <td>{{ temp.date }}</td>
                <td>{{ temp.value }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
    <div
        v-if="device.co2 && device.co2.length > 0"
        class="row">
      <div class="col-md-3">
        <strong>CO2:</strong>
      </div>
      <div class="col-md-12">
        <div class="table-responsive">
          <table class="table table-striped">
            <thead>
              <tr>
                <th>Time Acquired</th>
                <th>Sensor value</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="co2 in device.co2"
                :key="co2.date">
                <td>{{ co2.date }}</td>
                <td>{{ co2.value }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import DeviceService from '@/api-services/smartac.service'

export default {
  name: 'DeviceDetails',
  data () {
    return {
      device: {}
    }
  },
  created () {
    DeviceService.getDevice(this.$router.currentRoute.params.id).then((response) => {
      this.device = response.data
    })
  }
}
</script>
