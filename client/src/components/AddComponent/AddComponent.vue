<template>
    <div>

        <form>
            <h1>Add Employee</h1>

            <md-field>
                <label>Employee Name</label>
                <md-input type="text" placeholder="Mark Smith" v-model="name"></md-input>
            </md-field>
            <md-field>
                <label>Employee Age</label>
                <md-input type="text" placeholder="23" v-model="age"></md-input>
            </md-field>
            <md-field>
                <label>Employee Department</label>
                <md-input type="text" placeholder="Marketing" v-model="department"></md-input>
            </md-field>
            <md-field>
                <label>Employee Salary</label>
                <md-input type="text" placeholder="3900" v-model="salary"></md-input>
            </md-field>

            <div class="button-container" v-if="!showSnackbar">
                <button type="button" @click="submit">Add Employee</button>
            </div>
        </form>

        <md-snackbar :md-active.sync="showSnackbar">
            <span class="msg" v-bind:style="{ color: msgColor }">{{message}}</span>
        </md-snackbar>
    </div>
</template>


<script>
import axios from "axios";

export default {
    name: 'AddComponent',
    props: {},
    data() {
        return {
            name: undefined,
            age: undefined,
            department: undefined,
            salary: undefined,
            message: undefined,
            msgColor: "green",
            showSnackbar: false
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
                    this.showSnackbar = true
                    this.message = response.data.message

                    if (response.data.success) {
                        this.msgColor = "#006e10"
                        this.name = undefined
                        this.age = undefined
                        this.department = undefined
                        this.salary = undefined
                    } else {
                        this.msgColor = "#ff2929"
                    }
                })
                .catch(err => {
                    console.log(err)
                })

        }
    }
}
</script>


<style scoped lang="scss" src="./AddComponent.scss"></style>