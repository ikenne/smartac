import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import NotFound from '@/components/error-pages/NotFound'
import DeviceList from '@/components/device/DeviceList.component'
import DeviceDetails from '@/components/device/DeviceDetails.component'
import DeviceListService from '@/components/device/DeviceListService.component'
import DeviceListCO2Limit from '@/components/device/DeviceListCO2.component'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/device',
      name: 'DeviceList',
      component: DeviceList
    },
    {
      path: '/device/service',
      name: 'DeviceListService',
      component: DeviceListService
    },
    {
      path: '/device/co2limit',
      name: 'DeviceListCO2Limit',
      component: DeviceListCO2Limit
    },
    {
      path: '/device/:id',
      name: 'DeviceDetails',
      component: DeviceDetails
    },
    {
      path: '*',
      name: 'NotFound',
      component: NotFound
    }
  ]
})
