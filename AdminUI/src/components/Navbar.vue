<template>
  <b-navbar
    toggleable="md"
    type="dark"
    variant="info">
    <b-navbar-brand :to="{ name: 'Home' }">Smart AC</b-navbar-brand>
    <b-navbar-nav>
      <b-nav-item :to="{ name: 'DeviceList'}">Devices</b-nav-item>
      <b-nav-item :to="{ name: 'DeviceListService'}">Require Service</b-nav-item>
      <b-nav-item :to="{ name: 'DeviceListCO2Limit'}">CO2 Limit</b-nav-item>
      <!-- <b-nav-item href="#">Owner Actions</b-nav-item> -->
    </b-navbar-nav>
  </b-navbar>
</template>
<script>
export default {
  name: 'Navbar'
}
</script>
