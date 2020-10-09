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


<script src="./AddComponent.js"></script>
<style scoped lang="scss" src="./AddComponent.scss"></style>