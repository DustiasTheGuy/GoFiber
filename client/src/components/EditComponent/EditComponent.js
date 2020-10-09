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