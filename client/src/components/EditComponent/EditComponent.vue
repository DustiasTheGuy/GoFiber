<template>
  <div>


    <button v-if="editEmployee" class="btn-blue" @click="selectEmployee(null, false)">All Employees</button>

    <div class="container" v-if="!editEmployee">
      <div class="header">
        <span class="employee">Employee</span>
        <span class="department">Department</span>
        <span class="age">Age</span>
        <span class="salary">Salary</span>
      </div>
      <div @click="selectEmployee(emp, true)" class="body" v-for="emp in data" :key="emp._id" title="Edit Employee">
        <span class="employee"><b>Name:</b> {{emp.name}}</span>
        <span class="department"><b>Department:</b> {{emp.department}}</span>
        <span class="age"><b>Age:</b> {{emp.age}}</span>
        <span class="salary"><b>Salary:</b> {{emp.salary}}:-</span>
      </div>
    </div>
    <div class="container" v-if="editEmployee">
      <div class="edit-form">
        <h1>{{editEmployeeData.name}}</h1>
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
          <button class="btn-red" @click="deleteEmployee()">Delete Employee</button>
          <button class="btn-green" @click="submitUpdate()">Save Changes</button>
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
    data() {
        return {
            data: [],
            editEmployee: false,
            editEmployeeData: undefined,
            message: undefined,
            msgColor: "green",
            showSnackbar: false
        }
    },
    mounted() {
      this.$root.$on("newEmployee", (state) => {
        console.log("State Update:", state)
        this.loadData();
      })
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
                if (response.data.success) {
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
                if (response.data.success) {
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
<style scoped lang="scss" src="./EditComponent.scss"></style>