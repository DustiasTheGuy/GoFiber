<template>
  <div id="app">

      <header class="main-header">
        <h1>Employee Database</h1>
      </header>

      
      <main>
        <div class="grid-item">
            <AddComponent/>
        </div>
        <div class="grid-item">
            <EditComponent/>
        </div>
      </main>


  </div>
</template>

<script>

import AddComponent from './components/AddComponent/AddComponent'
import EditComponent from './components/EditComponent/EditComponent'


export default {
  name: 'App',
  components: {
      AddComponent,
      EditComponent
  }
}
</script>

<style lang="scss" src="./App.scss"></style>
