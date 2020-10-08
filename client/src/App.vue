<template>
  <div id="app">

      <header></header>

      
      <main>
        <AddComponent/>
        <EditComponent/>
      </main>


  </div>
</template>

<script>
import AddComponent from './components/AddComponent'
import EditComponent from './components/EditComponent'

export default {
  name: 'App',
  components: {
      AddComponent,
      EditComponent
  }
}
</script>

<style lang="scss">

  #app {
    height: 100%;
    width: 100%;
  }
  * {
    box-sizing: border-box;
  }

  html, body {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
  }

  header {
    height: 300px;
    background-color: olivedrab;
  }

  main {
    width: 100%;
    height: calc(100% - 300px);
    background-color: #f3e5ab;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-gap: 1em;
    padding: 1em;
  }
</style>
