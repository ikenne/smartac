import Axios from 'axios'

const RESOURCE_NAME = '/device'
const SERVICE = 'service'
const CO2LIMIT = 'co2limit'

export default {
  getAll () {
    return Axios.get(RESOURCE_NAME)
  },

  getDevice (id) {
    return Axios.get(`${RESOURCE_NAME}/${id}`)
  },

  getDeviceForService () {
    return Axios.get(`${RESOURCE_NAME}/${SERVICE}`)
  },

  getDeviceForCO2Limit () {
    return Axios.get(`${RESOURCE_NAME}/${CO2LIMIT}`)
  }

}
