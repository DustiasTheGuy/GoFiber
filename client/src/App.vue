<template>
  <div id="app">

      <header class="main-header">
        <h1>Employee Database</h1>
      </header>

      
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
  * {
    box-sizing: border-box;
  }

  button {
    border: none;
    padding: 1em;
  }

  html, body {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
    background-color: #c5c4c2;
  }


  header.main-header {
    height: 300px;
    background-color: rgb(16, 82, 168);
    display: flex;
    align-items: center;
    justify-content: center;
    color: #ffffff;
    font-size: 3em;
  }

  main {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-columns: 1fr 2fr;
    grid-gap: 1em;
    padding: 1em;
    background-color: #c5c4c2;
  }

  a {
    cursor: pointer;
  }

  .msg {
      font-size: 18px;
  }
  
  @media only screen and (max-width: 1280px) {
    main {
      grid-template-columns: 1fr;
    }
    header.main-header {
      font-size: 1.93em;
    }
  }

  @media only screen and (max-width: 680px) {
    header.main-header {
      font-size: 1.3em;
    }
  }

</style>
