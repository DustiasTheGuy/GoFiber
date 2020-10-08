<template>
  <div>


    <div class="container" v-if="!editEmployee">
        <div class="header">
          <span>Employee</span>
          <span>Department</span>
          <span>Age</span>
          <span class="salary">Salary</span>
        </div>
        <div class="body" v-for="emp in data" :key="emp._id">
          <span><b>Name:</b> {{emp.name}}</span>
          <span><b>Department:</b> {{emp.department}}</span>
          <span><b>Age:</b> {{emp.age}}</span>
          <span class="salary"><div><b>Salary:</b> {{emp.salary}}:-</div> <a @click="selectEmployee(emp, true)">Edit</a></span>
        </div>
    </div>
    <div class="container" v-if="editEmployee">
        <div class="edit-form">
          <h1>{{editEmployeeData.name}}</h1>
          <button class="deleteBtn" @click="deleteEmployee()">Delete</button>
          <md-field>
            <label>Employee</label>
            <md-input v-model="editEmployeeData.name"></md-input>
          </md-field>
          <md-field>
            <label>Department</label>
            <md-input v-model="editEmployeeData.department"></md-input>
          </md-field>
          <md-field>
            <label>Age</label>
            <md-input v-model="editEmployeeData.age"></md-input>
          </md-field>
          <md-field>
            <label>Salary</label>
            <md-input v-model="editEmployeeData.salary"></md-input>
          </md-field>

          <div class="button-container">
            <button class="goBackBtn" @click="selectEmployee(null, false)">Go Back</button>
            <button class="submitBtn" @click="submitUpdate()">Update</button>
          </div>
        </div>
    </div>

    <md-snackbar :md-active.sync="showSnackbar">
        <span class="msg" v-bind:style="{ color: msgColor }">{{message}}</span>
    </md-snackbar>

  </div>
</template>

<script>
import axios from "axios";
export default {
  name: 'EditComponent',
  data () {
    return {
      data: [],
      editEmployee: false,
      editEmployeeData: undefined,
      message: undefined,
      msgColor: "green",
      showSnackbar: false 
    }
  },
  created() {
    this.loadData()
  },
  methods: {
    loadData() {
      axios.get("http://localhost:3000/read").then(response => {
        this.data = response.data.data
        this.editEmployee = false;
      })
    },
    selectEmployee(emp, state) {
      this.editEmployee = state;
      this.editEmployeeData = emp;
    },
    deleteEmployee() {
      axios.post("http://localhost:3000/delete", this.editEmployeeData).then(response => {
        this.showSnackbar = true
        this.message = response.data.message
        if(response.data.success) {
            this.loadData();
            this.msgColor = "#006e10"
        } else {
            this.msgColor = "#ff2929"
        }
      })
    },
    submitUpdate() {
      axios.post("http://localhost:3000/update", {
        _id: this.editEmployeeData._id,
        name: this.editEmployeeData.name,
        department: this.editEmployeeData.department,
        age: parseInt(this.editEmployeeData.age),
        salary: parseInt(this.editEmployeeData.salary)
      }).then(response => {
        this.showSnackbar = true
        this.message = response.data.message
        if(response.data.success) {
            this.loadData();
            this.msgColor = "#006e10"
        } else {
            this.msgColor = "#ff2929"
        }
      })
    }
  },
  props: {}
}
</script>

<style scoped lang="scss">

    b {
      display: none;
    }
    @media only screen and (max-width: 680px) {
      h1 {
        font-size: 32px !important;
      }
      .header {
        display: none !important;
      }
      .body {
        display: flex !important;
        flex-direction: column;
        border-bottom: 1px solid lightgray;
        span {
          b {
            display: inline-block;
          }
          border: none !important;
        }
      }
    }
  
  .deleteBtn {
    background-color: crimson;
    color: #ffffff;
    position: absolute;
    top: 0;
    right: 0;
    margin: 2.43em;
    &:hover {
      background-color: rgb(179, 25, 56);
      cursor: pointer;
    }
  }
  .edit-form {
    padding: 2.43em;
    position: relative;
    h1 {
      font-size: 48px;
      color: #1052a8;
    }
    .button-container {
      display: flex;
      justify-content: space-between;

       .goBackBtn {
          background-color: #c5c4c2;
          &:hover {
            cursor: pointer;
            background-color: #a7a5a2;
          }
        }
        .submitBtn {
          background-color: rgb(20, 121, 20);
          color: #ffffff;
          &:hover {
          background-color: rgb(17, 102, 17);    
          }
        }
    }
  }

  .container {
    background: #ffffff;
    display: flex;
    flex-direction: column;

  .header {
      font-size: 18px;
      font-weight: bold;
      color: rgb(16, 82, 168);
    }
   .header, .body {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      span {
        border-right: 1px solid black;
        border-bottom: 1px solid black;
        padding: 1em;
      }

      span.salary {
        border-right: none;
        display: flex;
        justify-content: space-between;
      }
    }

    .body:last-child {
      span {
        border-bottom: none;
      }
    }
  }


</style>
