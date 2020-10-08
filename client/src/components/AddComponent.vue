<template>
  <div>

      <form>
        <h1>Add Employee</h1>

        <div class="form-group">
            <label for="">Employee Name</label>
            <input type="text" placeholder="Mark Smith" v-model="name">
        </div>
        <div class="form-group">
            <label for="">Employee Age</label>
            <input type="text" placeholder="23" v-model="age">
        </div>
        <div class="form-group">
            <label for="">Employee Department</label>
            <input type="text" placeholder="Marketing" v-model="department">
        </div>
        <div class="form-group">
            <label for="">Employee Salary</label>
            <input type="text" placeholder="3900" v-model="salary">
        </div>
        
        <div class="button-container">
            <button type="button" @click="submit">Add Employee</button>
        </div>
      </form>


        {{name}}
        {{age}}
        {{department}}
        {{salary}} 
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: 'AddComponent',
  props: {
  },
  data() {
      return {
        name: undefined,
        age: undefined,
        department: undefined,
        salary: undefined
      }
  },
  methods: {
      submit() {
          let data = {
              name: this.name,
              age: parseInt(this.age),
              department: this.department,
              salary: parseInt(this.salary)
          }        
          axios.post("http://localhost:3000/create", data)
          .then(response => {
              console.log(response)
          })
          .catch(err => {
              console.log(err)
          })
        
      }
  }
}
</script>

<style scoped lang="scss">
    form {
        background: #ffffff;
        padding: 1em;
        h1 {
            margin: 0;
            color: olivedrab;
            font-size: 48px;
        }

        .form-group {
            display: flex;
            flex-direction: column;
            padding: .77em 0;

            label {
                margin-bottom: 8px;
                font-weight: 600;
                font-size: 17px;
            }
            input {
                padding: .66em;
            }
        }

        .button-container {
            display: flex;
            justify-content: flex-end;

            button {
                padding: .88em;
                background-color: olivedrab;
                color: #ffffff;
                border: none;
            }
        }
    }
</style>
