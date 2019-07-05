<template>
  <div>
    <b-row>
      <b-col
        md="2"
        offset-md="10">
        <!-- <a href="#">Create owner</a>-->
        <!-- <router-link :to="{ name: 'DeviceCreate' }">Create owner</router-link> -->
      </b-col>
    </b-row>
    <br>
    <b-row>
      <b-col md="12">
        <div class="table-responsive">
          <table class="table table-striped">
            <thead>
              <tr>
                <th>Device S/N</th>
                <th>Registeration Date</th>
                <th>Firmware</th>
                <th>Status</th>
                <th>Details</th>
              </tr>
            </thead>
            <tbody>
              <device-list-row
                v-for="device in devices"
                :key="device.serialNumber"
                :device="device"
                @details="detailsDevice"
                />
            </tbody>
          </table>
        </div>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import SmartACService from '@/api-services/smartac.service'
import DeviceListRow from '@/components/device/DeviceListRow'

export default {
  name: 'DeviceList',
  components: {
    'device-list-row': DeviceListRow
  },
  data () {
    return {
      devices: [],
      selectedDeviceId: null,
      alertModalTitle: '',
      alertModalContent: ''
    }
  },
  created () {
    this.fetchDevices()
  },
  methods: {
    detailsDevice (serialNumber) {
      this.$router.push({ name: 'DeviceDetails', params: { id: serialNumber } })
    },
    fetchDevices () {
      SmartACService.getAll().then((response) => {
        this.devices = response.data
      })
    }
  }
}
</script>
